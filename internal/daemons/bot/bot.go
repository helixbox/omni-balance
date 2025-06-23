package bot

import (
	"context"
	"fmt"
	"sync"

	"omni-balance/internal/daemons/market"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils/bot"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Run starts the bot daemon process that manages token balance monitoring and task execution across chains.
// Parameters:
//   - ctx: Context for request cancellation and timeouts
//   - conf: Application configuration containing chain, wallet, and token settings
//
// Returns:
//   - error: Any error that occurs during setup or execution, including client creation errors, order creation failures, etc.
//
// The function:
// 1. Initializes blockchain clients for all configured chains
// 2. Identifies existing buy tokens to ignore
// 3. Processes all wallet-token-chain combinations concurrently
// 4. Creates and dispatches tasks to the market system
// 5. Waits for all goroutines to complete before exiting
func Run(ctx context.Context, conf configs.Config) error {
	existBuyTokens, err := getExistingBuyTokens()
	if err != nil {
		return errors.Wrap(err, "find buy tokens error")
	}
	clients := make(map[string]simulated.Client)
	for _, chain := range conf.Chains {
		client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
		if err != nil {
			return errors.Wrap(err, "create client error")
		}
		clients[chain.Name] = client
	}
	defer func() {
		for _, client := range clients {
			client.(*chains.Client).Close()
		}
	}()

	var (
		ignoreTokens = createIgnoreTokens(existBuyTokens)
		w            sync.WaitGroup
	)

	for _, wallet := range conf.Wallets {
		for _, token := range wallet.Tokens {
			for _, chainName := range token.Chains {
				if ignoreTokens.Contains(token.Name, chainName, wallet.Address) {
					continue
				}
				botTypes := conf.ListBotNames(wallet.Address, chainName, token.Name)
				botTypes = append(botTypes, "balance_on_chain")
				log.Debugf("wallet %s token %s on chain %s has %+v bots to execute", wallet.Address, token.Name, chainName, botTypes)
				for _, botType := range botTypes {
					w.Add(1)
					go func(f func() ([]bot.Task, bot.ProcessType, error), botType string) {
						defer w.Done()
						tasks, processType, err := f()
						if err != nil {
							log.Errorf("bot error: %s", err)
							return
						}
						if len(tasks) == 0 {
							return
						}
						orders, taskId, err := createOrder(ctx, tasks, processType)
						if err != nil {
							log.Errorf("create order error: %s", err)
							return
						}
						if len(orders) == 0 {
							return
						}
						log.Infof("create %d tasks, based %s on %s using %s bot", len(tasks), tasks[0].TokenOutName,
							tasks[0].TokenOutChainName, botType)
						market.PushTask(market.Task{
							Id:          taskId,
							ProcessType: processType,
						})
					}(process(ctx, conf, wallet.Address, token.Name, chainName, botType, clients[chainName]), botType)
				}
			}
		}
	}
	w.Wait()
	return nil
}

// process creates a task generator function for a specific wallet-token-chain-bot combination.
// Parameters:
//   - ctx: Context for request cancellation and timeouts
//   - conf: Application configuration
//   - walletAddress: Target wallet address to monitor
//   - tokenName: Token symbol to check balances for
//   - chainName: Blockchain network name
//   - botType: Type of bot logic to use (e.g., "balance_on_chain")
//   - client: Blockchain client connection
//
// Returns:
//   - func() ([]bot.Task, bot.ProcessType, error): Closure that when executed will:
//     1. Check current token balance and market conditions
//     2. Generate appropriate tasks based on the bot's logic
//     3. Adjust task parameters for balance mode wallets
//     4. Return tasks ready for execution
func process(ctx context.Context, conf configs.Config, walletAddress, tokenName, chainName, botType string,
	client simulated.Client,
) func() ([]bot.Task, bot.ProcessType, error) {
	m := bot.GetBot(botType)
	if m == nil {
		panic(fmt.Sprintf("%s botType not found", botType))
	}
	return func() ([]bot.Task, bot.ProcessType, error) {
		// get token price
		tokenPrices, err := models.FindTokenPrice(db.DB(), []string{tokenName})
		if err != nil {
			return nil, bot.Parallel, err
		}
		if len(tokenPrices) == 0 {
			tokenPrices = make(map[string]decimal.Decimal)
		}

		tasks, processType, err := m.Check(ctx, bot.Params{
			Conf: conf,
			Info: bot.Config{
				Wallet:     conf.GetWallet(walletAddress),
				TokenName:  tokenName,
				Chain:      chainName,
				TokenPrice: tokenPrices[tokenName],
			},
			Client: client,
		})
		if err != nil {
			return nil, bot.Parallel, err
		}
		var result []bot.Task
		for index, task := range tasks {
			if task.TokenInName != "" {
				result = append(result, tasks[index])
				continue
			}
			walletConf := conf.GetWalletConfig(task.Wallet)
			if !walletConf.Mode.IsBalance() {
				continue
			}
			t := new(bot.Task)
			if err := copier.Copy(t, &task); err != nil {
				return nil, bot.Parallel, err
			}
			t.TokenInName = task.TokenOutName
			log.Debugf("%s mode is balance, change tokenInName  to %s", task.Wallet, t.TokenInName)
			result = append(result, *t)
		}
		return result, processType, nil
	}
}
