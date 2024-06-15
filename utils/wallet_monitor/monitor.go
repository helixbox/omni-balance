package wallet_monitor

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/wallets"
	"sync"
)

type Monitor struct {
	config configs.Config
}

type ReBalanceTokenChain struct {
	BuyTokenInfo
	ChainName string `json:"chain_name"`
}

type BuyTokenInfo struct {
	TokenBalance decimal.Decimal `json:"balance"`
	Amount       decimal.Decimal `json:"amount"`
}

type Result struct {
	// need rebalance wallet
	Wallet string `json:"wallet"`
	// need rebalance tokens
	Tokens []ReBalanceToken `json:"tokens"`
}

type ReBalanceToken struct {
	Name   string                `json:"name"`
	Chains []ReBalanceTokenChain `json:"chains"`
}

type IgnoreToken struct {
	Name    string `json:"name"`
	Chain   string `json:"chain"`
	Address string `json:"wallet"`
}

func NewMonitor(conf configs.Config) *Monitor {
	return &Monitor{config: conf}
}

func (m *Monitor) GetBalance(ctx context.Context, wallet wallets.Wallets, tokenName, chainName string) (balance decimal.Decimal, err error) {
	var (
		chain = m.config.GetChainConfig(chainName)
		token = chain.GetToken(tokenName)
	)

	if len(chain.RpcEndpoints) == 0 {
		return decimal.Zero, errors.New("rpc endpoints is empty")
	}
	client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "dial rpc")
	}
	defer client.Close()
	return wallet.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, client)
}

func (m *Monitor) Check(ctx context.Context, ignoreToken ...IgnoreToken) (result []Result, err error) {
	var (
		// balances {"wallet": {"token_name": "chain_name": BuyTokenInfo}}
		balances     = map[string]map[string][]ReBalanceTokenChain{}
		w            sync.WaitGroup
		mutex        sync.Mutex
		ignoreTokens = make(map[string]map[string]map[string]struct{})
	)
	for _, token := range ignoreToken {
		if _, ok := ignoreTokens[token.Address]; !ok {
			ignoreTokens[token.Address] = make(map[string]map[string]struct{})
		}
		if _, ok := ignoreTokens[token.Address][token.Name]; !ok {
			ignoreTokens[token.Address][token.Name] = make(map[string]struct{})
		}
		ignoreTokens[token.Address][token.Name][token.Chain] = struct{}{}
	}

	appendBalances := func(wallet, tokenName string, rbtc ReBalanceTokenChain) {
		mutex.Lock()
		defer mutex.Unlock()
		if _, ok := balances[wallet]; !ok {
			balances[wallet] = make(map[string][]ReBalanceTokenChain)
		}
		balances[wallet][tokenName] = append(balances[wallet][tokenName], rbtc)
	}
	for _, wallet := range m.config.Wallets {
		w.Add(1)
		go func(wallet configs.Wallet) {
			defer utils.Recover()
			defer w.Done()
			for _, token := range wallet.Tokens {
				for _, chainName := range token.Chains {
					if _, ok := ignoreTokens[wallet.Address][token.Name][chainName]; ok {
						continue
					}
					balance, err := m.GetBalance(ctx, m.config.GetWallet(wallet.Address), token.Name, chainName)
					if err != nil {
						logrus.Errorf("get %s on %s balance error: %s", token.Name, chainName, err.Error())
						continue
					}

					threshold := m.config.GetTokenThreshold(wallet.Address, token.Name, chainName)
					logrus.Debugf("wallet %s token %s on %s balance is %s, the threshold is %s",
						wallet.Address, token.Name, chainName, balance.String(), threshold.String())
					if balance.GreaterThan(threshold) {
						continue
					}

					amount := m.config.GetTokenPurchaseAmount(wallet.Address, token.Name, chainName)
					if amount.LessThanOrEqual(decimal.Zero) {
						logrus.Errorf("token %s config amount is invalid", token.Name)
						continue
					}
					logrus.Infof("wallet %s token %s on %s need rebalance %s",
						wallet.Address, token.Name, chainName, amount.String())
					appendBalances(wallet.Address, token.Name, ReBalanceTokenChain{
						BuyTokenInfo: BuyTokenInfo{
							TokenBalance: balance,
							Amount:       amount,
						},
						ChainName: chainName,
					})
				}
			}
		}(wallet)
	}
	w.Wait()
	for walletAddress, item := range balances {
		r := &Result{Wallet: walletAddress}
		for tokenName, chains := range item {
			r.Tokens = append(r.Tokens, ReBalanceToken{
				Name:   tokenName,
				Chains: chains,
			})
		}
		result = append(result, *r)
	}
	return
}
