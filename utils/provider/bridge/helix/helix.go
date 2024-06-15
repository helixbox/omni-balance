package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"
)

type Bridge struct {
	config configs.Config
}

func New(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Bridge{}, nil
	}
	return &Bridge{config: conf}, nil
}

func (b *Bridge) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	var (
		lastHistory  = args.LastHistory
		actionNumber = Action2Int(lastHistory.Actions)
		txHash       = common.HexToHash(lastHistory.Tx)
		log          = args.GetLogs(b.Name())
		wallet       = args.Sender
		recordFn     = func(s provider.SwapHistory, errs ...error) {
			s.ProviderType = string(b.Type())
			s.ProviderName = b.Name()
			s.Amount = args.Amount
			if args.RecordFn == nil {
				return
			}
			args.RecordFn(s, errs...)
		}
	)
	if lastHistory.Actions == targetChainReceivedAction && lastHistory.Status == string(provider.TxStatusSuccess) {
		return provider.SwapResult{
			ProviderType: b.Type(),
			ProviderName: b.Name(),
			OrderId:      lastHistory.Tx,
			Status:       provider.TxStatusSuccess,
			CurrentChain: args.TargetChain,
			Tx:           lastHistory.Tx,
		}, nil
	}
	args.SourceToken = args.TargetToken

	if args.SourceChain == "" {
		sourceChainId, err := b.GetSourceChain(ctx, args.TargetChain, args.TargetToken,
			wallet.GetAddress(true).Hex(), args.Amount)
		if err != nil {
			return result, errors.Wrap(err, "get source chain")
		}
		args.SourceChain = constant.GetChainName(sourceChainId)
	}
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.SourceChain)
	log = args.GetLogs(b.Name()).WithFields(
		logrus.Fields{
			"sourceChain": args.SourceChain,
			"tokenIn":     args.SourceToken,
		},
	)
	chain := b.config.GetChainConfig(args.SourceChain)
	ethClient, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return result, errors.Wrap(err, "dial rpc")
	}
	defer ethClient.Close()

	var tokenInfo = b.config.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
	if tokenInfo.Decimals == 0 || tokenInfo.Name == "" {
		return result, errors.Errorf("token %s not supported", args.SourceToken)
	}

	if actionNumber <= 1 && provider.TxStatus(lastHistory.Status).CanRetry() {
		transferOptions, err := GetTransferOptions(
			ctx,
			args.Amount,
			tokenInfo.Decimals,
			args.SourceChain,
			args.TargetChain,
			common.HexToAddress(tokenInfo.ContractAddress),
		)
		if err != nil {
			return result, errors.Wrap(err, "get transfer options")
		}
		if _, ok := BRIDGES[BridgeType(transferOptions._bridge)]; !ok {
			return result, errors.Errorf("bridge %s not supported", transferOptions._bridge)
		}
		log.Debugf("transfer options: %+v", transferOptions)
		bridge := BRIDGES[BridgeType(transferOptions._bridge)](Options{
			SourceTokenName: args.SourceToken,
			TargetTokenName: args.TargetToken,
			SourceChain:     args.SourceChain,
			TargetChain:     args.TargetChain,
			Config:          b.config,
			Sender:          wallet.GetAddress(true),
			Recipient:       wallet.GetAddress(true),
			Amount:          args.Amount,
		})
		tx, err := bridge.Do(ctx, transferOptions)
		if err != nil {
			return result, errors.Wrap(err, "do transfer")
		}
		tx.Gas = 406775

		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			Receiver:        common.HexToAddress(args.Receiver),
			TokenIn:         args.SourceToken,
			TokenOut:        args.TargetToken,
			TokenInChain:    args.SourceChain,
			TokenOutChain:   args.TargetChain,
			ProviderName:    b.Name(),
			TokenInAmount:   args.Amount,
			TokenOutAmount:  args.Amount,
			TransactionType: provider.TransferTransactionAction,
		})

		recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain})
		txHash, err = wallet.SendTransaction(ctx, tx, ethClient)
		if err != nil {
			recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain})
			return result, errors.Wrap(err, "send signed transaction")
		}
		recordFn(provider.SwapHistory{Actions: sourceChainSendingAction, Status: string(provider.TxStatusSuccess),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
	}

	sr := &provider.SwapResult{
		TokenInName:  args.SourceToken,
		ProviderType: configs.Bridge,
		ProviderName: b.Name(),
		Status:       provider.TxStatusPending,
		CurrentChain: args.SourceChain,
		Tx:           txHash.Hex(),
		OrderId:      txHash.Hex(),
	}

	if actionNumber <= 2 {
		recordFn(provider.SwapHistory{Actions: sourceChainSentAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
		log.Debugf("wait for tx %s", txHash.Hex())
		if err := wallet.WaitTransaction(ctx, txHash, ethClient); err != nil {
			recordFn(provider.SwapHistory{Actions: sourceChainSentAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain, Tx: txHash.Hex()})
			return *sr, errors.Wrap(err, "wait for tx")
		}
		recordFn(provider.SwapHistory{Actions: sourceChainSentAction, Status: string(provider.TxStatusSuccess),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
	}

	var record HistoryRecord
	if actionNumber <= 3 {
		recordFn(provider.SwapHistory{Actions: targetChainSendingAction, Status: string(provider.TxStatusPending),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
		record, err = b.WaitForBridge(ctx, wallet.GetAddress(true), txHash)
		if err != nil {
			recordFn(provider.SwapHistory{Actions: targetChainSendingAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain, Tx: txHash.Hex()})
			sr.Error = err.Error()
			return *sr, errors.Wrap(err, "wait for bridge")
		}
		log.Debugf("bridge result: %+v", record)
		if record.Result != 3 {
			recordFn(provider.SwapHistory{Actions: targetChainSendingAction, Status: string(provider.TxStatusFailed),
				CurrentChain: args.SourceChain, Tx: txHash.Hex()})
			return *sr, errors.Errorf("bridge failed, result: %d", record.Result)
		}
		sr.Order = record
		recordFn(provider.SwapHistory{Actions: targetChainReceivedAction, Status: string(provider.TxStatusSuccess),
			CurrentChain: args.SourceChain, Tx: txHash.Hex()})
		sr.Status = provider.TxStatusSuccess
	}
	return *sr, nil
}

func (b *Bridge) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if args.TargetChain == "" {
		return nil, errors.Errorf("chain is empty")
	}
	if args.TargetToken == "" {
		return nil, errors.Errorf("target token is empty")
	}

	if args.SourceChain == "" {
		sourceChainId, err := b.GetSourceChain(ctx, args.TargetChain, args.TargetToken,
			args.Sender.GetAddress(true).Hex(), args.Amount)
		if err != nil {
			return nil, err
		}
		args.SourceChain = constant.GetChainName(sourceChainId)
	}
	if args.SourceToken == "" {
		args.SourceToken = b.config.GetTokenInfoOnChain(args.TargetToken, args.SourceChain).Name
	}

	if err := args.Sender.CheckFullAccess(ctx); err != nil {
		return nil, errors.Errorf("wallet %s private key error: %s;",
			args.Sender.GetAddress(true), err)
	}
	return provider.TokenInCosts{
		provider.TokenInCost{
			TokenName:  args.TargetToken,
			CostAmount: args.Amount,
		},
	}, nil
}

func (b *Bridge) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string, _ decimal.Decimal) (bool, error) {
	supportedChains, err := GetTokenSupportedChains(tokenName)
	if err != nil {
		return false, err
	}
	for _, v := range supportedChains {
		if v.FromChain == tokenInChainName && utils.InArrayFold(tokenOutChainName, v.ToChains) {
			return true, nil
		}
	}
	return false, nil
}

func (b *Bridge) Help() []string {
	return []string{"See https://helixbridge.app/"}
}

func (b *Bridge) Name() string {
	return "helixbridge"
}

func (b *Bridge) Type() configs.LiquidityProviderType {
	return configs.Bridge
}
