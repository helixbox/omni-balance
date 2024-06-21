package rebalance

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"omni-balance/internal/daemons"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/notice"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"
	"sync"
	"time"
)

func init() {
	daemons.RegisterIntervalTask(daemons.Task{
		Name:            "rebalance",
		Description:     "Responsible for injecting specified assets into monitored wallets in the most efficient way when the balance of specified tokens is insufficient.",
		TaskFunc:        Run,
		DefaultInterval: time.Minute * 10,
	})
}

var (
	orderError = make(map[uint]int)
	m          sync.Mutex
)

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

func Run(ctx context.Context, conf configs.Config) error {
	orders, err := listOrders(ctx)
	if err != nil {
		return errors.Wrap(err, "find orders error")
	}
	if len(orders) == 0 {
		return nil
	}
	var w sync.WaitGroup
	for index := range orders {
		if utils.InArray(orders[index].Status.String(), []string{models.OrderStatusWaitCrossChain.String()}) {
			continue
		}
		if orders[index].HasLocked() {
			logrus.Debugf("order #%d has locked, skip", orders[index].ID)
			continue
		}
		w.Add(1)
		go func(order *models.Order) {
			defer utils.Recover()
			defer w.Done()
			log := order.GetLogs()
			subCtx, cancel := context.WithCancel(utils.SetLogToCtx(ctx, log))
			defer cancel()
			utils.Go(func() {
				defer cancel()
				var t = time.NewTicker(time.Second * 5)
				defer t.Stop()

				for {
					select {
					case <-subCtx.Done():
						return
					case <-t.C:
						var count int64
						_ = db.DB().Model(&models.Order{}).Where("id = ?", order.ID).Count(&count)
						if count == 0 {
							log.Infof("order #%d not found, exit this order rebalance", order.ID)
							return
						}
					}
				}
			})

			err := reBalance(subCtx, order, conf)
			if errors.Is(err, error_types.ErrNoProvider) {
				if addOrderError(order.ID) > 3 {
					log.Errorf("order #%d not found provider, exit this order rebalance. send notice", order.ID)
					_ = notice.Send(
						provider.WithNotify(ctx, provider.WithNotifyParams{
							OrderId:       order.ID,
							Receiver:      common.HexToAddress(order.Wallet),
							TokenOut:      order.TokenOutName,
							TokenOutChain: order.TargetChainName,
						}),
						"Find provider error",
						fmt.Sprintf("order #%d not found provider, please check the provider configuration.", order.ID),
						logrus.ErrorLevel,
					)
					return
				}
			}
			if err != nil {
				log.Errorf("reBalance order #%d error: %s", order.ID, err)
				return
			}
			removeOrderError(order.ID)
			order = models.GetOrder(ctx, db.DB(), order.ID)

			err = notice.Send(
				provider.WithNotify(ctx, provider.WithNotifyParams{
					OrderId:        order.ID,
					Receiver:       common.HexToAddress(order.Wallet),
					CurrentBalance: order.CurrentBalance,
				}),
				fmt.Sprintf("rebalance %s on %s success", order.TokenOutName, order.TargetChainName),
				fmt.Sprintf("rebalance %s %s from %s to %s use %s %s",
					order.TokenOutName, order.Amount, order.SourceChainName, order.TargetChainName,
					order.ProviderName, order.ProviderType),
				logrus.InfoLevel,
			)
			if err != nil {
				log.Debugf("notice error: %s", err)
			}
			log.Infof("reBalance order #%d success", order.ID)
		}(orders[index])
	}
	w.Wait()
	return nil
}

func transfer(ctx context.Context, order *models.Order, args provider.SwapParams,
	conf configs.Config, client simulated.Client) (bool, error) {
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, order.TargetChainName)
	if order.Status != models.OrderStatusWaitTransferFromOperator {
		return false, errors.Errorf("order #%d status is %s, not wait transfer from operator", order.ID, order.Status)
	}
	result, err := provider.Transfer(ctx, conf, args, client)
	if errors.Is(err, error_types.ErrNativeTokenInsufficient) ||
		errors.Is(err, error_types.ErrWalletLocked) ||
		errors.Is(err, context.Canceled) {
		return true, errors.Wrap(err, "transfer error")
	}
	if err == nil {
		return true, createUpdateLog(ctx, order, result, conf, client)
	}
	if !errors.Is(errors.Unwrap(err), error_types.ErrInsufficientBalance) &&
		!errors.Is(errors.Unwrap(err), error_types.ErrInsufficientLiquidity) && err != nil {
		return true, errors.Wrap(err, "transfer not is insufficient balance")
	}
	return false, nil
}

func reBalance(ctx context.Context, order *models.Order, conf configs.Config) error {
	log := utils.GetLogFromCtx(ctx)
	if order.Lock(db.DB()) {
		return errors.Errorf("order #%d locked, unlock time is %s", order.ID, time.Unix(order.LockTime+60*60*1, 0))
	}
	defer order.UnLock(db.DB())
	var (
		orderProcess = models.GetLastOrderProcess(ctx, db.DB(), order.ID)
		args         = daemons.CreateSwapParams(*order, orderProcess, log, conf.GetWallet(order.Wallet))
		wallet       = conf.GetWallet(order.Wallet)
		token        = conf.GetTokenInfoOnChain(order.TokenOutName, order.TargetChainName)
		chain        = conf.GetChainConfig(order.TargetChainName)
		client, err  = chains.NewTryClient(ctx, chain.RpcEndpoints)
	)

	if err != nil {
		return errors.Wrap(err, "new evm client error")
	}
	defer client.Close()
	if wallet.IsDifferentAddress() || order.Status == models.OrderStatusWaitTransferFromOperator {
		ok, err := transfer(ctx, order, args, conf, client)
		if err != nil && ok {
			return errors.Wrap(err, "transfer error")
		}
		if ok {
			return nil
		}
		log.Infof("cannot use transfer, try other providers.")
	}

	balance, err := wallet.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, client)
	if err != nil {
		return errors.Wrap(err, "check balance error")
	}

	for _, v := range conf.GetWalletConfig(order.Wallet).Tokens {
		if !utils.InArray(order.TargetChainName, v.Chains) {
			continue
		}
		if order.TokenOutName != v.Name {
			continue
		}
		if !balance.GreaterThan(balance) {
			break
		}
		log.Infof("%s balance on %s is enough, skip", v.Name, order.TargetChainName)
		if err := order.Success(db.DB(), "", nil, balance); err != nil {
			return errors.Wrap(err, "update order success error")
		}
		return nil
	}

	providerObj, err := getReBalanceProvider(ctx, *order, conf)
	if err != nil {
		return errors.Wrap(err, "get reBalance provider error")
	}

	if err := order.SaveProvider(db.DB(), providerObj.Type(), providerObj.Name()); err != nil {
		return errors.Wrap(err, "save provider error")
	}

	log.Infof("start reBalance #%d %s on %s use %s provider", order.ID, order.TokenOutName,
		order.TargetChainName, providerObj.Name())
	result, providerErr := providerObj.Swap(ctx, args)
	if result.Status == "" {
		return errors.New("the result status is empty")
	}
	if result.CurrentChain != args.TargetChain {
		result.Status = models.OrderStatusWaitCrossChain
	}
	if err := createUpdateLog(ctx, order, result, conf, client); err != nil {
		return errors.Wrap(err, "create update log error")
	}
	if providerErr != nil {
		return errors.Wrap(providerErr, "provider error")
	}
	if args.Receiver != result.Receiver && result.Receiver != "" {
		order = models.GetOrder(ctx, db.DB(), order.ID)
		if order == nil {
			return errors.New("order not found")
		}
		_, err = transfer(ctx, order, daemons.CreateSwapParams(*order, orderProcess, log, conf.GetWallet(order.Wallet)), conf, client)
		if err != nil {
			return errors.Wrap(err, "transfer error")
		}
	}
	return nil
}

func listOrders(_ context.Context) ([]*models.Order, error) {
	var orders []*models.Order
	err := db.DB().Where("status != ?", provider.TxStatusSuccess).Find(&orders).Error
	if err != nil {
		return nil, errors.Wrap(err, "find buy tokens error")
	}
	return orders, nil
}

func createUpdateLog(ctx context.Context, order *models.Order, result provider.SwapResult, conf configs.Config,
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

func getReBalanceProvider(ctx context.Context, order models.Order, conf configs.Config) (provider.Provider, error) {
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

func providerSupportsOrder(ctx context.Context, p provider.Provider, order models.Order, conf configs.Config, log *logrus.Entry) (provider.TokenInCosts, bool) {
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
