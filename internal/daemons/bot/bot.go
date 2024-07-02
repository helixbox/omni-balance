package bot

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"omni-balance/internal/daemons/market"
	"omni-balance/utils/bot"
	"omni-balance/utils/bot/balance_on_chain"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"sync"
)

func Run(ctx context.Context, conf configs.Config) error {
	existBuyTokens, err := getExistingBuyTokens()
	if err != nil {
		return errors.Wrap(err, "find buy tokens error")
	}
	var (
		clients = make(map[string]simulated.Client)
	)
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
					logrus.Debugf("ignore token %s on chain %s", token.Name, chainName)
					continue
				}
				monitorType := token.MonitorTypes[chainName]
				if monitorType == "" {
					monitorType = balance_on_chain.BalanceOnChain{}.Name()
				}
				fn := process(ctx, conf, wallet.Address, token.Name, chainName, monitorType, clients[chainName])
				w.Add(1)
				go func() {
					defer w.Done()
					tasks, processType, err := fn()
					if err != nil {
						logrus.Errorf("bot error: %s", err)
						return
					}
					_, taskId, err := createOrder(tasks, processType)
					if err != nil {
						logrus.Errorf("create order error: %s", err)
						return
					}
					market.PushTask(market.Task{
						Id:          taskId,
						ProcessType: processType,
					})
				}()
			}
		}
	}
	w.Wait()
	return nil
}

func process(ctx context.Context, conf configs.Config, walletAddress, tokenName, chainName, monitorType string,
	client simulated.Client) func() ([]bot.Task, bot.ProcessType, error) {
	m := bot.GetMonitor(monitorType)
	return func() ([]bot.Task, bot.ProcessType, error) {
		return m.Check(ctx, bot.Params{
			Conf: conf,
			Info: bot.Config{
				Wallet:    conf.GetWallet(walletAddress),
				TokenName: tokenName,
				Chain:     chainName,
			},
			Client: client,
		})
	}
}
