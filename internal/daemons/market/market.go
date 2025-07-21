package market

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"omni-balance/internal/daemons"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/bot"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/notice"
	"omni-balance/utils/provider"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.uber.org/zap/zapcore"
)

var (
	hasRunQueue  = atomic.Bool{}
	processTasks sync.Map
)

// Run initializes and manages the market order processing system.
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - conf: Application configuration containing provider and chain settings
//
// Returns:
//   - error: Initialization errors including:
//   - Task listing failures
//   - Queue initialization problems
//
// The function:
// 1. Starts background queue processing (once only)
// 2. Recovers incomplete orders from persistent storage
// 3. Pushes unfinished tasks back into the processing queue
func Run(ctx context.Context, conf configs.Config) error {
	if !hasRunQueue.Load() {
		go runFromQueue(ctx, conf)
	}
	hasRunQueue.Store(true)
	tasks, err := models.ListNotSuccessTasks(ctx, db.DB(), func(order models.Order) bool {
		if order.HasLocked() {
			return false
		}
		return !utils.InArrayFold(order.Status.String(), []string{
			models.OrderStatusWaitTransferFromOperator.String(),
			models.OrderStatusWaitCrossChain.String(),
		})
	})
	if err != nil {
		return errors.Wrap(err, "list tasks error")
	}
	for index := range tasks {
		if _, ok := processTasks.Load(tasks[index].Id); ok {
			continue
		}
		PushTask(Task{
			Id:          tasks[index].Id,
			ProcessType: bot.ProcessType(tasks[index].ProviderType),
		})
	}
	return nil
}

// runFromQueue continuously processes tasks from the market queue.
// Parameters:
//   - ctx: Context for cancellation
//   - conf: Application configuration
//
// This background process:
// 1. Listens for incoming tasks from the global queue
// 2. Handles panics with automatic restart
// 3. Processes orders concurrently based on process type
// 4. Manages task locking to prevent duplicate processing
func runFromQueue(ctx context.Context, conf configs.Config) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("market job error: %v, restart after 5 second", err)
			time.Sleep(time.Second * 5)
			runFromQueue(ctx, conf)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return

		case task := <-taskQueue:
			if _, ok := processTasks.Load(task.Id); ok {
				return
			}
			orders, err := models.ListOrdersByTaskId(ctx, db.DB(), task.Id)
			if err != nil {
				log.Errorf("get orders by task id error: %s", err.Error())
				continue
			}
			if len(orders) == 0 {
				continue
			}
			fn := func(orders []models.Order, task Task) func() {
				processTasks.Store(task.Id, struct{}{})
				return func() {
					defer func() { processTasks.Delete(task.Id) }()

					var taskWait sync.WaitGroup
					for index := range orders {
						if orders[index].HasLocked() {
							continue
						}
						if task.ProcessType == bot.Parallel {
							taskWait.Add(1)
							go func(order models.Order) {
								taskWait.Done()
								do(ctx, order, conf)
							}(orders[index])
							continue
						}
						do(ctx, orders[index], conf)
					}
					taskWait.Wait()
				}
			}
			utils.Go(fn(orders, task))
		}
	}
}

// do handles individual order processing with status monitoring.
// Parameters:
//   - ctx: Parent context for cancellation
//   - order: Order to process
//   - conf: Application configuration
//
// The function:
// 1. Creates a monitored sub-context with cancellation
// 2. Implements periodic order existence checks
// 3. Delegates to processOrder for core logic
// 4. Handles error notifications and success alerts
// 5. Maintains error counters for retry logic
func do(ctx context.Context, order models.Order, conf configs.Config) {
	defer utils.Recover()
	subCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	utils.Go(func() {
		defer cancel()
		t := time.NewTicker(time.Second * 5)
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

	err := processOrder(subCtx, order, conf)
	if errors.Is(err, context.Canceled) || errors.Is(ctx.Err(), context.Canceled) || errors.Is(subCtx.Err(), context.Canceled) {
		return
	}
	if err != nil {
		if addOrderError(order.ID) >= 10 {
			_ = notice.Send(
				provider.WithNotify(ctx, provider.WithNotifyParams{
					TaskId:        order.TaskId,
					OrderId:       order.ID,
					Receiver:      common.HexToAddress(order.Wallet),
					TokenOut:      order.TokenOutName,
					TokenOutChain: order.TargetChainName,
				}),
				"process order error",
				fmt.Sprintf("order #%d error: %s, please check the provider configuration then restart the application.", order.ID, err),
				zapcore.ErrorLevel,
			)
			return
		}
		log.Errorf(" order #%d error: %s", order.ID, err)
		return
	}
	removeOrderError(order.ID)
	order = models.GetOrder(ctx, db.DB(), order.ID)
	if order.SourceChainName == "" || order.Tx == "" {
		return
	}
	err = notice.Send(
		provider.WithNotify(ctx, provider.WithNotifyParams{
			TaskId:         order.TaskId,
			OrderId:        order.ID,
			Receiver:       common.HexToAddress(order.Wallet),
			CurrentBalance: order.CurrentBalance,
		}),
		fmt.Sprintf("rebalance %s on %s success", order.TokenOutName, order.TargetChainName),
		fmt.Sprintf("rebalance %s %s from %s to %s use %s %s",
			order.TokenOutName, order.Amount, order.SourceChainName, order.TargetChainName,
			order.ProviderName, order.ProviderType),
		zapcore.InfoLevel,
	)
	if err != nil {
		log.Debugf("notice error: %s", err)
	}
	log.Infof(" order #%d success", order.ID)
}

// processOrder executes the core order fulfillment logic.
// Parameters:
//   - ctx: Context for cancellation
//   - order: Target order to process
//   - conf: Application configuration
//
// Returns:
//   - error: Execution errors including:
//   - Order locking failures
//   - Balance check errors
//   - Provider selection errors
//   - Swap execution failures
//
// The function:
// 1. Acquires order lock
// 2. Checks wallet balances against thresholds
// 3. Selects optimal liquidity provider
// 4. Executes cross-chain swap
// 5. Updates order status and handles token transfers
func processOrder(ctx context.Context, order models.Order, conf configs.Config) error {
	if order.Lock(db.DB()) {
		return errors.Errorf("order #%d locked, unlock time is %s", order.ID, time.Unix(order.LockTime+60*60*1, 0))
	}
	defer order.UnLock(db.DB())
	var (
		orderProcess = models.GetLastOrderProcess(ctx, db.DB(), order.ID)
		args         = daemons.CreateSwapParams(order, orderProcess, conf.GetWallet(order.Wallet))
		wallet       = conf.GetWallet(order.Wallet)
		token        = conf.GetTokenInfoOnChain(order.TokenOutName, order.TargetChainName)
		chain        = conf.GetChainConfig(order.TargetChainName)
		client, err  = chains.NewTryClient(ctx, chain.RpcEndpoints)
	)

	if err != nil {
		return errors.Wrap(err, "new evm client error")
	}
	defer client.Close()
	if wallet.IsDifferentAddress() && order.Status == models.OrderStatusWaitTransferFromOperator {
		ok, err := transfer(ctx, order, args, conf, client)
		if err != nil && ok {
			return errors.Wrap(err, "transfer error")
		}
		if ok {
			return nil
		}
	}

	balance, err := wallet.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, client)
	if err != nil {
		return errors.Wrap(err, "check balance error")
	}
	walletConfig := conf.GetWalletConfig(order.Wallet)
	for _, v := range walletConfig.Tokens {
		if !utils.InArray(order.TargetChainName, v.Chains) {
			continue
		}
		if order.TokenOutName != v.Name {
			continue
		}
		threshold := conf.GetTokenThreshold(order.Wallet, v.Name, order.TargetChainName)
		if !balance.GreaterThan(threshold) {
			break
		}
		if strings.EqualFold(order.ProcessType, "Bridge") {
			break
		}
		log.Debugf("%s balance on %s is enough, skip", v.Name, order.TargetChainName)
		if err := order.Success(db.DB(), "", nil, balance); err != nil {
			return errors.Wrap(err, "update order success error")
		}
		return nil
	}

	order, err = generateOrderByWalletMode(ctx, order, conf)
	if err != nil {
		return errors.Wrap(err, "generate order by wallet mode error")
	}

	providerObj, err := getBestProvider(ctx, order, conf)
	if err != nil {
		return errors.Wrap(err, "get  provider error")
	}

	if err := order.SaveProvider(db.DB(), providerObj.Type(), providerObj.Name()); err != nil {
		return errors.Wrap(err, "save provider error")
	}

	log.Infof("start rebalance %s on %s use %s provider", order.TokenOutName, order.TargetChainName, providerObj.Name())
	args.SourceChainNames = order.TokenInChainNames
	args.SourceToken = order.TokenInName
	result, providerErr := providerObj.Swap(ctx, args)
	if errors.Is(providerErr, context.Canceled) {
		return nil
	}
	if result.Status == "" {
		return errors.Errorf("the result status is empty: %v", providerErr)
	}
	if result.CurrentChain != args.TargetChain && providerErr == nil {
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
		if order.ID == 0 {
			return errors.New("order not found")
		}
		_, err = transfer(ctx, order, daemons.CreateSwapParams(order, orderProcess, conf.GetWallet(order.Wallet)), conf, client)
		if err != nil {
			return errors.Wrap(err, "transfer error")
		}
	}
	return nil
}

// generateOrderByWalletMode configures order parameters based on wallet balance mode.
// Parameters:
//   - ctx: Context for cancellation
//   - order: Original order object
//   - conf: Application configuration
//
// Returns:
//   - models.Order: Modified order with source chain information
//   - error: Configuration errors including:
//   - Chain client initialization failures
//   - Balance check errors
//
// For 'balance' mode wallets:
// 1. Identifies source chains with sufficient balances
// 2. Configures order for multi-chain sourcing
// 3. Maintains original order if no modifications needed
func generateOrderByWalletMode(ctx context.Context, order models.Order, conf configs.Config) (models.Order, error) {
	if order.TokenInName != "" {
		return order, nil
	}
	if conf.GetWalletConfig(order.Wallet).Mode != "balance" ||
		(order.TokenInName != "" && order.SourceChainName != "") {
		return order, nil
	}
	token := conf.GetWalletTokenInfo(order.Wallet, order.TokenOutName)
	threshold := conf.GetTokenThreshold(order.Wallet, token.Name, order.TargetChainName)
	var sourceChains []string
	for _, v := range token.Chains {
		client, err := chains.NewTryClient(ctx, conf.GetChainConfig(v).RpcEndpoints)
		if err != nil {
			return order, errors.Wrap(err, "new evm client error")
		}
		bots := conf.ListBotNames(order.Wallet, v, token.Name)
		if len(bots) == 0 {
			bots = append(bots, "balance_on_chain")
		}
		total := decimal.Zero
		for _, botType := range bots {
			balance, err := bot.GetBot(botType).Balance(ctx, bot.Params{
				Conf: conf,
				Info: bot.Config{
					Wallet:    conf.GetWallet(order.Wallet),
					TokenName: token.Name,
					Chain:     v,
				},
				Client: client,
			})
			if err != nil {
				return order, errors.Wrap(err, "get bot balance error")
			}
			total = total.Add(balance)
		}
		if total.Sub(order.Amount).LessThan(threshold) {
			continue
		}
		log.Debugf("wallet %s token %s on chain %s balance is %s, amount is %s, balance - amount >= threshold, can rebalance from this chain", order.Wallet, token.Name, v, total, order.Amount)
		sourceChains = append(sourceChains, v)
	}
	if len(sourceChains) != 0 {
		newOrder := new(models.Order)
		_ = copier.Copy(newOrder, order)
		newOrder.TokenInName = token.Name
		newOrder.TokenInChainNames = sourceChains
		order = *newOrder
		log.Debugf("order mode is 'balance', use %s on %+v as token in, token out is %s on %s", order.TokenInName, order.TokenInChainNames, order.TokenOutName, order.TargetChainName)
	}
	return order, nil
}
