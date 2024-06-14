package darwinia

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"
	"strings"
	"time"
)

var (
	// action
	sourceChainSendingAction  = "sourceChainSending"
	sourceChainWaitAction     = "sourceChainWait"
	targetChainSendingAction  = "targetChainSending"
	targetChainReceivedAction = "targetChainReceived"
)

func Action2Int(action string) int {
	switch action {
	case sourceChainSendingAction:
		return 1
	case sourceChainWaitAction:
		return 2
	case targetChainSendingAction:
		return 3
	case targetChainReceivedAction:
		return 4
	default:
		return 0
	}
}

type Bridge struct {
	config configs.Config
}

var graphOlQuery = "query GetHistory($sender: String, $page: Int, $row: Int) {\n  historyRecords(sender: $sender, page: $page, row: $row) {\n    total\n    records {\n      requestTxHash\n      responseTxHash\n      fromChain\n      toChain\n      startTime\n      sendToken\n      sendAmount\n      result\n      id\n      __typename\n    }\n    __typename\n  }\n}"

type SwapParams struct {
	Sender    wallets.Wallets
	TokenName string
	Amount    decimal.Decimal
	Nonce     int64
	Client    simulated.Client
}

type GraphQLRequest struct {
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
	Query         string                 `json:"query"`
}

type HistoryRecordResponse struct {
	Data struct {
		HistoryRecords struct {
			Total    int             `json:"total"`
			Records  []HistoryRecord `json:"records"`
			Typename string          `json:"__typename"`
		} `json:"historyRecords"`
	} `json:"data"`
}

type HistoryRecord struct {
	RequestTxHash  string `json:"requestTxHash"`
	ResponseTxHash string `json:"responseTxHash"`
	FromChain      string `json:"fromChain"`
	ToChain        string `json:"toChain"`
	StartTime      int    `json:"startTime"`
	SendToken      string `json:"sendToken"`
	SendAmount     string `json:"sendAmount"`
	Result         int    `json:"result"`
	Id             string `json:"id"`
	Typename       string `json:"__typename"`
}

func New(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Bridge{}, nil
	}
	d := &Bridge{
		config: conf,
	}
	return d, nil
}

func (b *Bridge) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	amount decimal.Decimal) (bool, error) {

	sourceChains := GetSourceChains(int64(constant.GetChainId(tokenInChainName)), tokenName)
	if len(sourceChains) == 0 {
		return false, nil
	}
	return utils.InArray(int64(constant.GetChainId(tokenInChainName)), sourceChains), nil
}

func (b *Bridge) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if args.TargetToken == "" || args.TargetChain == "" {
		return nil, errors.Errorf("target token or target chain is empty")
	}
	if args.SourceChain == "" {
		sourceChains := b.FindValidSourceChains(ctx, constant.GetChainId(args.TargetChain), args.TargetToken,
			args.Sender.GetAddress(true).Hex(), args.Amount)
		if len(sourceChains) == 0 {
			return nil, error_types.ErrUnsupportedTokenAndChain
		}
		args.SourceChain = constant.GetChainName(utils.Choose(sourceChains))
	}

	if args.SourceToken == "" {
		args.SourceToken = b.config.GetTokenInfoOnChain(args.TargetToken, args.SourceChain).Name
	}

	return provider.TokenInCosts{
		{
			TokenName:  args.SourceToken,
			CostAmount: args.Amount,
		},
	}, nil
}

func (b *Bridge) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	var (
		history  = args.LastHistory
		log      = args.GetLogs("darwinia")
		recordFn = func(s provider.SwapHistory, errs ...error) {
			s.ProviderType = string(b.Type())
			s.ProviderName = b.Name()
			s.Amount = args.Amount
			if args.RecordFn == nil {
				return
			}
			args.RecordFn(s, errs...)
		}
	)
	if history.Actions == targetChainReceivedAction && history.Status == string(provider.TxStatusSuccess) {
		log.Debugf("target chain received, order id: %s", history.Tx)
		return provider.SwapResult{
			ProviderType: b.Type(),
			ProviderName: b.Name(),
			OrderId:      history.Tx,
			Status:       provider.TxStatusSuccess,
			CurrentChain: args.TargetChain,
			Tx:           history.Tx,
		}, nil
	}
	if args.SourceChain == args.TargetChain && history.Status == string(provider.TxStatusSuccess) {
		log.Debugf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
		return provider.SwapResult{}, errors.Errorf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
	}
	if history.CurrentChain != "" && history.CurrentChain != args.TargetChain {
		args.SourceChain = history.CurrentChain
	}
	actionNumber := Action2Int(history.Actions)
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.SourceChain)
	if actionNumber == 0 && args.SourceChain == "" {
		validSourceChain := b.GetValidSourceChain(ctx, constant.GetChainId(args.TargetChain),
			args.TargetToken, args.Sender.GetAddress(true).Hex(), args.Amount)
		if validSourceChain == 0 {
			return result, errors.Errorf("can not find source chain for %s %s", args.TargetToken, args.TargetChain)
		}
		args.SourceChain = constant.GetChainName(validSourceChain)
	}
	var sourceChainConf = b.config.GetChainConfig(args.SourceChain)
	if args.SourceToken == "" {
		args.SourceToken = b.config.GetTokenInfoOnChain(args.TargetToken, args.SourceChain).Name
	}
	sourceToken := b.config.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)

	targetChainConf := b.config.GetChainConfig(args.TargetChain)

	fn := BRIDGE_DIRECTION[int64(sourceChainConf.Id)][int64(targetChainConf.Id)]
	if fn == nil {
		return result, errors.Errorf("not support swap %s to %s", args.SourceChain, args.TargetChain)
	}
	wallet := args.Sender
	ethClient, err := chains.NewTryClient(ctx, sourceChainConf.RpcEndpoints)
	if err != nil {
		return result, errors.Wrap(err, "dial rpc")
	}
	defer ethClient.Close()
	log.Debugf("start transfer %s from %s to %s", args.SourceToken, args.SourceChain, args.TargetChain)
	tx, err := fn(ctx, SwapParams{
		Sender:    wallet,
		TokenName: args.SourceToken,
		Amount:    decimal.NewFromBigInt(chains.EthToWei(args.Amount, sourceToken.Decimals), 0),
		Nonce:     time.Now().UnixMilli(),
		Client:    ethClient,
	})
	if err != nil {
		return result, errors.Wrap(err, "swap")
	}

	if tx.Gas != 0 {
		tx.Gas = tx.Gas * 2
	}

	sr := &provider.SwapResult{
		TokenInName:  args.SourceToken,
		ProviderType: configs.Bridge,
		ProviderName: b.Name(),
		Status:       provider.TxStatusPending,
		CurrentChain: args.SourceChain,
	}

	if actionNumber <= 1 && history.Status != string(provider.TxStatusSuccess) {
		recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain})
		txHash, err := args.Sender.SendTransaction(ctx, tx, ethClient)
		if err != nil {
			recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain}, err)
			return result, errors.Wrap(err, "send signed transaction")
		}
		recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusSuccess),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
		sr.Tx = txHash.Hex()
		sr.OrderId = sr.Tx
	}

	if actionNumber <= 2 && history.Status != string(provider.TxStatusSuccess) {
		if sr.Tx == "" {
			sr.Tx = history.Tx
		}
		recordFn(provider.SwapHistory{Actions: sourceChainWaitAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain, Tx: sr.Tx})
		err = wallet.WaitTransaction(ctx, common.HexToHash(sr.Tx), ethClient)
		if err != nil && strings.Contains(err.Error(), "tx failed") {
			recordFn(provider.SwapHistory{Actions: sourceChainWaitAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain, Tx: sr.Tx}, err)
			return result, err
		}
		if err != nil {
			recordFn(provider.SwapHistory{Actions: sourceChainWaitAction, Status: string(provider.TxStatusPending),
				CurrentChain: args.SourceChain, Amount: args.Amount, Tx: sr.Tx}, err)
			return result, errors.Wrap(err, "wait for tx")
		}
		recordFn(provider.SwapHistory{Actions: sourceChainWaitAction, Status: string(provider.TxStatusSuccess),
			CurrentChain: args.SourceChain, Tx: sr.Tx})
	}

	if actionNumber <= 4 && history.Status != string(provider.TxStatusSuccess) {
		if sr.Tx == "" {
			sr.Tx = history.Tx
		}
		recordFn(provider.SwapHistory{Actions: targetChainSendingAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain, Tx: sr.Tx})
		log.Debugf("waiting for bridge success")
		record, err := b.WaitForBridgeSuccess(ctx, sr.Tx, args.Sender.GetAddress(true).Hex())
		if err != nil {
			recordFn(provider.SwapHistory{Actions: targetChainSendingAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain, Tx: sr.Tx}, err)
			return *sr, errors.Wrap(err, "wait for bridge success")
		}
		sr.Order = record

		if record.Result == 3 {
			recordFn(provider.SwapHistory{Actions: targetChainReceivedAction, Status: string(provider.TxStatusSuccess),
				CurrentChain: args.TargetChain, Tx: sr.Tx})
			sr.Status = provider.TxStatusSuccess
			sr.CurrentChain = args.TargetChain
		}
		if record.Result != 3 {
			log.Debugf("bridge failed: %+v", record)
			recordFn(provider.SwapHistory{Actions: targetChainReceivedAction,
				Status: string(provider.TxStatusFailed), CurrentChain: args.SourceChain, Tx: sr.Tx},
				errors.Errorf("bridge failed: %d", record.Result))
			sr.Status = provider.TxStatusFailed
		}
	}

	return *sr, nil
}

func (b *Bridge) Check(ctx context.Context, args provider.CheckParams) (bool, error) {
	if args.Chain == "" {
		return false, errors.Errorf("target chain is empty")
	}
	if err := b.config.GetWallet(args.Wallet).CheckFullAccess(ctx); err != nil {
		return false, errors.Wrap(err, "check full access")
	}

	targetChain := b.config.GetChainConfig(args.Chain)
	sourceChains := b.FindValidSourceChains(ctx, targetChain.Id, args.Token, args.Wallet, args.Amount)
	if len(sourceChains) == 0 {
		return false, nil
	}
	return true, nil
}

func (b *Bridge) Help() []string {
	return []string{"See https://bridge.darwinia.network"}
}

func (b *Bridge) Name() string {
	return "darwinia-bridge"
}

func (b *Bridge) Type() configs.LiquidityProviderType {
	return configs.Bridge
}
