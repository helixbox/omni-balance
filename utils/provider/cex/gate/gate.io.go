package gate

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"strings"
	"time"
)

var (
	PurchasingAction  = "purchasing"
	PurchasedAction   = "purchased"
	WithdrawingAction = "withdrawing"
	WithdrawnAction   = "withdrawn"

	successStatus = "success"
	pendingStatus = "pending"
	failStatus    = "fail"
)

func StatusToInt(status string) int {
	switch status {
	case successStatus:
		return 3
	case pendingStatus:
		return 1
	case failStatus:
		return 2
	default:
		return 0
	}
}

func Action2Int(action string) int {
	switch action {
	case PurchasingAction:
		return 1
	case PurchasedAction:
		return 2
	case WithdrawingAction:
		return 3
	case WithdrawnAction:
		return 4
	default:
		return 0
	}
}

type Gate struct {
	authConfig gateConfig
	config     configs.Config
	client     *gateapi.APIClient
}

type gateConfig struct {
	Key    string `json:"key" yaml:"key" help:"API key"`
	Secret string `json:"secret" yaml:"secret" help:"API secret"`
}

func New(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Gate{}, nil
	}
	d := &Gate{
		config: conf,
	}

	err := conf.GetProvidersConfig(d.Name(), configs.CEX, &d.authConfig)
	if err != nil {
		return nil, errors.Wrap(err, "get gate config")
	}
	c := gateapi.NewConfiguration()
	c.Key = d.authConfig.Key
	c.Secret = d.authConfig.Secret
	d.client = gateapi.NewAPIClient(c)
	_, _, err = d.client.AccountApi.GetAccountDetail(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "init gate api")
	}
	return d, nil
}

func (g *Gate) Swap(ctx context.Context, args provider.SwapParams) (provider.SwapResult, error) {
	var (
		recordFn = func(s provider.SwapHistory, errs ...error) {
			s.ProviderType = string(g.Type())
			s.ProviderName = g.Name()
			s.Amount = args.Amount
			if args.RecordFn == nil {
				return
			}
			args.RecordFn(s, errs...)
		}
		log = logrus.WithFields(logrus.Fields{
			"provider":    g.Name(),
			"amount":      args.Amount,
			"tokenOut":    args.TargetToken,
			"targetChain": args.TargetChain,
		})
		wallet = args.Sender
	)

	isSenderInVerifiedAddress, err := g.IsVerifiedAddress(ctx, wallet.GetAddress(true).Hex(), args.TargetToken, args.TargetChain)
	if err != nil {
		return provider.SwapResult{}, errors.Wrap(err, "check sender address")
	}
	if !isSenderInVerifiedAddress {
		return provider.SwapResult{}, errors.Errorf("%s not in gate.io verified address for %s",
			args.Sender, args.TargetToken)
	}

	actionNumber := Action2Int(args.LastHistory.Actions)

	tokenInCosts, err := g.GetCost(ctx, args)
	if err != nil {
		return provider.SwapResult{}, err
	}
	log.Debugf("tokenInCosts: %+v", tokenInCosts)
	tokenIn, err := tokenInCosts.GetBest()
	if err != nil {
		return provider.SwapResult{}, errors.Wrapf(err, "%s get best token in cost", g.Name())
	}
	log.Debugf("get best token in cost: %s %s", tokenIn.TokenName, tokenIn.CostAmount)

	accounts, _, err := g.client.SpotApi.ListSpotAccounts(context.Background(), &gateapi.ListSpotAccountsOpts{})
	if err != nil {
		return provider.SwapResult{}, errors.Wrap(err, "list spot accounts")
	}
	if len(accounts) == 0 {
		return provider.SwapResult{}, errors.Errorf("no %s in spot account", tokenIn.TokenName)
	}
	var (
		buyAmount       = args.Amount
		tokenInAccount  gateapi.SpotAccount
		tokenOutAccount gateapi.SpotAccount
	)
	if strings.EqualFold(tokenIn.TokenName, args.TargetToken) {
		return provider.SwapResult{}, errors.Errorf("token in and token out are the same")
	}
	for _, account := range accounts {
		log.Debugf("account: %+v, tokenInName: %s, tokenOutName: %s", account, tokenIn.TokenName, args.TargetToken)
		if strings.EqualFold(account.Currency, tokenIn.TokenName) {
			tokenInAccount = account
		}
		if strings.EqualFold(account.Currency, args.TargetToken) {
			tokenOutAccount = account
			continue
		}
	}
	log.Debugf("token in account: %+v, token out account: %+v", tokenInAccount, tokenOutAccount)

	if actionNumber <= 1 && decimal.RequireFromString(tokenInAccount.Available).LessThan(tokenIn.CostAmount) {
		return provider.SwapResult{}, errors.Errorf("not enough %s in spot account, available: %s, need: %s",
			tokenIn.TokenName, tokenInAccount.Available, tokenIn.CostAmount)
	}

	var result = &provider.SwapResult{
		ProviderType: g.Type(),
		ProviderName: g.Name(),
		TokenInName:  tokenIn.TokenName,
		OrderId:      args.LastHistory.Tx,
	}

	if buyAmount.GreaterThan(decimal.Zero) && actionNumber <= 1 && StatusToInt(args.LastHistory.Status) <= 2 { // buy token
		recordFn(provider.SwapHistory{Actions: PurchasingAction, Status: pendingStatus})
		order, err := g.buyToken(ctx, tokenIn, args.TargetToken, buyAmount.Abs(), func(order gateapi.Order) bool {
			logrus.Debugf("wait for order %s filled, the status is %s", order.Id, order.Status)
			return false
		})
		if err != nil {
			recordFn(provider.SwapHistory{Actions: PurchasingAction, Status: failStatus}, err)
			return provider.SwapResult{}, err
		}
		result.Order = order
		result.OrderId = order.Id
		recordFn(provider.SwapHistory{Actions: PurchasedAction, Status: successStatus, Tx: result.OrderId})
	}

	// 提现
	var withdrawOrderId = args.LastHistory.Tx

	if actionNumber < 4 && StatusToInt(args.LastHistory.Status) <= 2 {
		balance, _ := decimal.NewFromString(tokenOutAccount.Available)
		if balance.LessThanOrEqual(decimal.Zero) {
			err = errors.Errorf("not enough %s in spot account", args.TargetToken)
			recordFn(provider.SwapHistory{
				Actions: WithdrawingAction, Status: failStatus, CurrentChain: args.SourceChain, Tx: result.OrderId}, err)
			return provider.SwapResult{}, err
		}
		if balance.LessThanOrEqual(args.Amount) {
			args.Amount = balance
		}
		recordFn(provider.SwapHistory{Actions: WithdrawingAction, Status: pendingStatus, Tx: result.OrderId})
		withdrawOrder, _, err := g.client.WithdrawalApi.Withdraw(ctx, gateapi.LedgerRecord{
			Amount:   args.Amount.String(),
			Currency: args.TargetToken,
			Address:  wallet.GetAddress(true).Hex(),
			Chain:    ChainName2GateChainName(args.TargetChain),
		})
		if err != nil {
			recordFn(provider.SwapHistory{Actions: WithdrawingAction, Status: failStatus, CurrentChain: args.SourceChain, Tx: result.OrderId}, err)
			return provider.SwapResult{}, errors.Wrap(err, "withdraw")
		}
		withdrawOrderId = withdrawOrder.Id
		result.Order = withdrawOrder
		recordFn(provider.SwapHistory{Actions: WithdrawnAction, Status: pendingStatus, CurrentChain: args.SourceChain, Tx: withdrawOrder.Id})
	}

	for {
		logrus.Debugf("wait for %s withdraw order done", withdrawOrderId)
		withdrawalRecords, _, err := g.client.WalletApi.ListWithdrawals(ctx, &gateapi.ListWithdrawalsOpts{
			Currency: optional.NewString(args.TargetToken),
		})
		if err != nil {
			return *result, errors.Wrap(err, "list withdraw status")
		}
		if len(withdrawalRecords) == 0 {
			return *result, errors.Errorf("withdraw order not found")
		}
		var withdrawalRecord gateapi.WithdrawalRecord
		for index, v := range withdrawalRecords {
			if withdrawOrderId == "" &&
				strings.EqualFold(v.Chain, ChainName2GateChainName(args.TargetChain)) &&
				strings.EqualFold(v.Currency, args.TargetToken) {
				withdrawalRecord = withdrawalRecords[index]
				break
			}
			if v.Id == withdrawOrderId {
				withdrawalRecord = withdrawalRecords[index]
				break
			}
		}
		if withdrawalRecord.Status == "" {
			return *result, errors.Errorf("withdraw order not found")
		}
		switch strings.ToUpper(withdrawalRecord.Status) {
		case "DONE":
			recordFn(provider.SwapHistory{Actions: WithdrawnAction, Status: successStatus, CurrentChain: args.TargetChain, Tx: withdrawalRecord.Txid})
			result.Tx = withdrawalRecord.Txid
			result.Status = provider.TxStatusSuccess
			result.CurrentChain = args.TargetChain
			result.OrderId = withdrawalRecord.Id
			result.Order = withdrawalRecord
			log.Infof("withdraw done, txid: %s", withdrawalRecord.Txid)
			return *result, nil
		case "CANCEL", "FAIL":
			recordFn(provider.SwapHistory{
				Actions:      WithdrawnAction,
				Status:       failStatus,
				CurrentChain: args.SourceChain,
				Tx:           withdrawalRecord.Id})
			return *result, errors.Errorf("withdraw failed, status: %s", withdrawalRecord.Status)
		case "REQUEST", "BCODE", "EXTPEND", "VERIFY", "PROCES", "PEND", "LOCKED":
			time.Sleep(time.Second)
			continue
		default:
			recordFn(provider.SwapHistory{
				Actions:      WithdrawnAction,
				Status:       failStatus,
				CurrentChain: args.SourceChain,
				Tx:           withdrawalRecord.Txid})
			return *result, errors.Errorf("unknown status: %+v", withdrawalRecord)
		}
	}
}

func (g *Gate) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	log := logrus.WithFields(logrus.Fields{
		"tokenOut":    args.TargetToken,
		"targetChain": args.TargetChain,
		"amount":      args.Amount,
		"wallet":      args.Sender,
	})
	verifiedAddress, _, err := g.client.WalletApi.ListSavedAddress(ctx, args.TargetToken, &gateapi.ListSavedAddressOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "list saved address")
	}
	if len(verifiedAddress) == 0 {
		return nil, errors.Errorf("%s not in gate.io verified address for %s", args.Sender, args.TargetToken)
	}
	chains, err := g.GetTokenOutChain(ctx, args.TargetToken)
	if err != nil {
		return nil, errors.Wrap(err, "get token out chains")
	}
	log.Debugf("token %s out chains: %+v", args.TargetToken, chains)
	if !utils.InArray(args.TargetChain, chains) {
		providers, err := provider.GetTokenCrossChainProviders(ctx, provider.GetTokenCrossChainProvidersParams{
			SourceChains: chains,
			TargetChain:  args.TargetChain,
			TokenName:    args.TargetToken,
			Amount:       args.Amount,
			Conf:         g.config,
		})
		if err != nil {
			return nil, errors.Wrap(err, "get token cross chain providers")
		}
		if len(providers) == 0 {
			return nil, error_types.ErrUnsupportedTokenAndChain
		}
	}

	currency, _, err := g.client.SpotApi.GetCurrency(ctx, TokenName2GateTokenName(args.TargetToken))
	if err != nil {
		return nil, errors.Wrap(err, "get currency")
	}
	log.Debugf("currency: %+v", currency)
	if currency.Delisted ||
		currency.WithdrawDisabled {
		return nil, error_types.ErrUnsupportedTokenAndChain
	}

	// 检查是否有余额
	accounts, _, err := g.client.SpotApi.ListSpotAccounts(ctx, &gateapi.ListSpotAccountsOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "list spot accounts")
	}
	var (
		tokenIns []gateapi.SpotAccount
	)

	for index, v := range accounts {
		log.Debugf("spot %s token balance: %s", v.Currency, v.Available)
		if strings.EqualFold(v.Currency, TokenName2GateTokenName(args.TargetToken)) {
			continue
		}
		for _, tokenIn := range g.config.SourceToken {
			if !strings.EqualFold(v.Currency, tokenIn.Name) {
				continue
			}
			tokenIns = append(tokenIns, accounts[index])
			continue
		}
	}

	currencyPairs, _, err := g.client.SpotApi.ListCurrencyPairs(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "list currency pairs")
	}
	var (
		tokenInNames []string
		pairs        []string
	)
	for _, v := range currencyPairs {
		if !strings.EqualFold(v.Base, args.TargetToken) {
			continue
		}
		for _, tokenIn := range tokenIns {
			if !strings.EqualFold(tokenIn.Currency, v.Quote) {
				continue
			}
			log.Debugf("currency pair: %s_%s", v.Base, v.Quote)
			tokenInNames = append(tokenInNames, tokenIn.Currency)
			pairs = append(pairs, strings.ToLower(fmt.Sprintf("%s_%s", v.Base, v.Quote)))
		}
	}
	if len(tokenInNames) == 0 {
		return nil, errors.Errorf("no token in")
	}

	pairsList, err := g.Tickers(pairs...)
	if err != nil {
		return nil, errors.Wrap(err, "get market list")
	}
	if len(pairsList) == 0 {
		return nil, errors.Errorf("no market list")
	}
	var providerTokenIns []provider.TokenInCost
	for _, v := range pairsList {
		if strings.EqualFold(v.TokenName, args.TargetToken) {
			providerTokenIns = append(providerTokenIns, provider.TokenInCost{
				TokenName:  strings.ToUpper(v.TokenName),
				CostAmount: args.Amount,
			})
			continue
		}
		providerTokenIns = append(providerTokenIns, provider.TokenInCost{
			TokenName:  strings.ToUpper(v.TokenName),
			CostAmount: v.Price.Mul(args.Amount),
		})
	}
	return providerTokenIns, nil
}

func (g *Gate) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	// 检查token提现
	token, _, err := g.client.SpotApi.GetCurrency(ctx, TokenName2GateTokenName(tokenName))
	if err != nil {
		return false, errors.Wrap(err, "get currency")
	}
	return !token.WithdrawDisabled, err
}

func (g *Gate) Help() []string {
	var result []string
	for _, v := range utils.ExtractTagFromStruct(&gateConfig{}, "yaml", "help") {
		result = append(result, v["yaml"]+": "+v["help"])
	}
	result = append(result, "api key must be granted with the following permissions: Spot Trade, Wallet, Withdraw, Account")
	result = append(result, "API Key Type must be selected as API v4 Key, and since gate.io has a limitation, it cannot be used with sub-accounts")
	return result
}

func (g *Gate) Name() string {
	return "gate.io"
}

func (g *Gate) Type() configs.LiquidityProviderType {
	return configs.CEX
}
