package market

import (
	"context"
	"sync"

	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils/configs"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"omni-balance/utils/provider/bridge/darwinia"
	"omni-balance/utils/wallets"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	orderError = make(map[uint]int)
	m          sync.Mutex
)

func createUpdateLog(ctx context.Context, order models.Order, result provider.SwapResult, conf configs.Config,
	client simulated.Client,
) error {
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
	if result.Status == provider.TxStatusSuccess &&
		wallet.IsDifferentAddress() &&
		result.Receiver != order.Wallet &&
		result.Receiver != "" {
		updateOrder.Status = provider.TxStatus(models.OrderStatusWaitTransferFromOperator)
	}
	if result.Error != "" {
		log.Errorf("#%d wallet %s rebalance %s on %s %s token failed, error: %s", order.ID, order.Wallet, order.TokenOutName, order.TargetChainName, order.TokenInName, result.Error)
	}

	return db.DB().Model(&models.Order{}).Where("id = ?", order.ID).Limit(1).Updates(updateOrder).Error
}

func getWalletTokenBalance(ctx context.Context, wallet wallets.Wallets, tokenName, chainName string,
	conf configs.Config, client simulated.Client,
) decimal.Decimal {
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
	if order.ProviderType != "" && order.ProviderName != "" {
		log.Debugf("#%d wallet %s rebalance %s on %s %s token use provider %s, provider name is %s", order.ID, order.Wallet, order.TokenOutName, order.TargetChainName, order.TokenInName, order.ProviderType, order.ProviderName)
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
			p, err := provider.Init(providerFn, conf)
			if err != nil {
				log.Debugf("init provider error: %s", err.Error())
				continue
			}
			tokenInCosts, ok := providerSupportsOrder(ctx, p, order, conf)
			if !ok || len(tokenInCosts) == 0 {
				continue
			}
			log.Infof("provider %s can use %s on %s. The tokenInCosts is %+v",
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
			log.Warnf("find %+v token price error: %s", tokenNames, err.Error())
			continue
		}
		// if token out is RING and provider is darwinia, then use darwinia as the provider
		if order.TokenOutName == "RING" && canUseProvider.provider.Name() == new(darwinia.Bridge).Name() {
			return canUseProvider.provider, nil
		}
		if len(tokenName2Price) == 0 {
			log.Warnf("not found %+v price, set to 0", tokenNames)
			minPrice = decimal.Zero
			providerObj = canUseProvider.provider
			continue
		}

		for name, v := range tokenName2Price {
			log.Debugf("token %s price %s on %s", name, v.String(), canUseProvider.provider.Name())
			if v.LessThanOrEqual(decimal.Zero) {
				log.Warnf("token %s price %s is less than or equal to zero, set to 1",
					name, v.String())
				v = decimal.RequireFromString("1")
			}
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
	conf configs.Config,
) (provider.TokenInCosts, bool) {
	tokenInCosts, err := p.GetCost(ctx, provider.SwapParams{
		SourceToken:      order.TokenInName,
		Sender:           conf.GetWallet(order.Wallet),
		TargetToken:      order.TokenOutName,
		SourceChain:      order.SourceChainName,
		Receiver:         order.Wallet,
		TargetChain:      order.TargetChainName,
		Amount:           order.Amount,
		SourceChainNames: order.TokenInChainNames,
		CurrentBalance:   order.CurrentBalance,
	})
	if err != nil {
		log.Debugf("token %s on %s cannot use %s provider, source chain is '%s', source token is '%s', err : %s", order.TokenOutName, order.TargetChainName, p.Name(), order.SourceChainName, order.TokenInName, err)
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
