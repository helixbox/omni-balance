package market

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"
	"sync"
)

var (
	orderError = make(map[uint]int)
	m          sync.Mutex
)

func createUpdateLog(ctx context.Context, order models.Order, result provider.SwapResult, conf configs.Config,
	client simulated.Client) error {

	wallet := conf.GetWallet(order.Wallet)
	walletBalance := getWalletTokenBalance(ctx, wallet, order.TokenOutName, order.TargetChainName, conf, client)

	updateOrder := &models.Order{
		TokenInName:      result.TokenInName,
		SourceChainName:  result.TokenInChainName,
		CurrentChainName: result.CurrentChain,
		CurrentBalance:   walletBalance,
		ProviderOrderId:  result.OrderId,
		Tx:               result.Tx,
		Order:            result.MarshalOrder(),
		Error:            result.Error,
		Status:           result.Status,
	}
	log := utils.GetLogFromCtx(ctx).WithFields(logrus.Fields{
		"order_id": order.ID,
		"result":   utils.ToMap(result),
	})
	if result.Status == provider.TxStatusSuccess &&
		wallet.IsDifferentAddress() &&
		result.Receiver != order.Wallet &&
		result.Receiver != "" {
		updateOrder.Status = provider.TxStatus(models.OrderStatusWaitTransferFromOperator)
	}
	log.Debugf("order status is %v", updateOrder.Status)
	return db.DB().Model(&models.Order{}).Where("id = ?", order.ID).Limit(1).Updates(updateOrder).Error
}

func getWalletTokenBalance(ctx context.Context, wallet wallets.Wallets, tokenName, chainName string,
	conf configs.Config, client simulated.Client) decimal.Decimal {

	chainConfig := conf.GetChainConfig(chainName)
	if len(chainConfig.RpcEndpoints) == 0 {
		return decimal.Zero
	}
	token := conf.GetTokenInfoOnChain(tokenName, chainName)

	balance, err := wallet.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, client)
	if err != nil {
		return decimal.Zero
	}
	return balance
}

func getBestProvider(ctx context.Context, order models.Order, conf configs.Config) (provider.Provider, error) {
	log := order.GetLogs()
	if order.ProviderType != "" && order.ProviderName != "" {
		log.Debugf("provider type is %s, provider name is %s", order.ProviderType, order.ProviderName)
		fn, err := provider.GetProvider(order.ProviderType, order.ProviderName)
		if err != nil {
			return nil, errors.Wrap(err, "get provider error")
		}
		return fn(conf)
	}
	type canUseProvider struct {
		provider     provider.Provider
		tokenInCosts provider.TokenInCosts
	}
	var canUseProviders []canUseProvider
	providers := provider.ListProvidersByConfig(conf)
	for _, providerFns := range providers {
		for _, providerFn := range providerFns {
			p, err := provider.InitializeBridge(providerFn, conf)
			if err != nil {
				log.Debugf("init provider error: %s", err.Error())
				continue
			}
			log = log.WithFields(logrus.Fields{
				"provider_type": p.Type(),
				"provider_name": p.Name(),
			})
			tokenInCosts, ok := providerSupportsOrder(ctx, p, order, conf, log)
			if !ok || len(tokenInCosts) == 0 {
				continue
			}
			log.Debugf("provider %s can use %s on %s. The tokenInCosts is %+v",
				p.Name(), order.TokenOutName, order.TargetChainName, tokenInCosts)
			canUseProviders = append(canUseProviders, canUseProvider{
				provider:     p,
				tokenInCosts: tokenInCosts,
			})
		}
	}

	if len(canUseProviders) <= 0 {
		return nil, error_types.ErrNoProvider
	}
	if len(canUseProviders) == 1 {
		log.Debugf("can use %s provider, the tokenIn is %+v", canUseProviders[0].provider.Name(), canUseProviders[0].tokenInCosts)
		return canUseProviders[0].provider, nil
	}
	var (
		minPrice    decimal.Decimal
		providerObj provider.Provider
	)
	for _, canUseProvider := range canUseProviders {
		var (
			tokenNames      []string
			tokenInCostsMap = make(map[string]decimal.Decimal)
		)
		for _, tokenIn := range canUseProvider.tokenInCosts {
			if tokenIn.TokenName == order.TokenInName {
				return canUseProvider.provider, nil
			}
			tokenInCostsMap[tokenIn.TokenName] = tokenIn.CostAmount
			tokenNames = append(tokenNames, tokenIn.TokenName)

		}
		tokenName2Price, err := models.FindTokenPrice(db.DB(), tokenNames)
		if err != nil {
			log.Warnf("find token price error: %s", err.Error())
			continue
		}

		for name, v := range tokenName2Price {
			log.Debugf("token %s price %s on %s", name, v.String(), canUseProvider.provider.Name())
			price := v.Mul(tokenInCostsMap[name])
			if price.IsZero() {
				continue
			}
			if price.LessThan(minPrice) {
				minPrice = price
				providerObj = canUseProvider.provider
				continue
			}
			if minPrice.IsZero() {
				minPrice = price
				providerObj = canUseProvider.provider
			}
		}
	}
	if providerObj == nil {
		return nil, errors.New("no provider can use")
	}
	log.Debugf("min price %s, provider %s", minPrice, providerObj.Name())
	return providerObj, nil
}

func providerSupportsOrder(ctx context.Context, p provider.Provider, order models.Order,
	conf configs.Config, log *logrus.Entry) (provider.TokenInCosts, bool) {
	tokenInCosts, err := p.GetCost(ctx, provider.SwapParams{
		SourceToken: order.TokenInName,
		Sender:      conf.GetWallet(order.Wallet),
		TargetToken: order.TokenOutName,
		Receiver:    order.Wallet,
		TargetChain: order.TargetChainName,
		Amount:      order.Amount,
	})
	if err != nil {
		log.Debugf("check token %s on %s use %s error: %s", order.TokenOutName, order.TargetChainName, p.Name(), err.Error())
		return nil, false
	}
	return tokenInCosts, true
}

func addOrderError(orderId uint) int {
	m.Lock()
	defer m.Unlock()
	orderError[orderId] = orderError[orderId] + 1
	return orderError[orderId]
}

func removeOrderError(orderId uint) {
	m.Lock()
	defer m.Unlock()
	delete(orderError, orderId)
}
