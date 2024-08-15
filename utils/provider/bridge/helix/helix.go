package helix

import (
	"context"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
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
			wallet.GetAddress(true).Hex(), args.Amount, args.SourceChainNames...)
		if err != nil {
			return result, errors.Wrap(err, "get source chain")
		}
		args.SourceChain = constant.GetChainName(sourceChainId)
	}
	sr := new(provider.SwapResult).
		SetTokenInName(args.SourceToken).
		SetTokenInChainName(args.SourceChain).
		SetProviderName(b.Name()).
		SetProviderType(b.Type()).
		SetCurrentChain(args.SourceChain).
		SetTx(lastHistory.Tx)
	sh := &provider.SwapHistory{
		ProviderName: b.Name(),
		ProviderType: string(b.Type()),
		Amount:       args.Amount,
		CurrentChain: args.SourceChain,
		Tx:           lastHistory.Tx,
	}
	isActionSuccess := lastHistory.Status == string(provider.TxStatusSuccess)

	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.SourceChain)
	chain := b.config.GetChainConfig(args.SourceChain)
	ethClient, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return sr.SetError(err).Out(), errors.Wrap(err, "dial rpc")
	}
	defer ethClient.Close()

	var tokenInfo = b.config.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
	if tokenInfo.Decimals == 0 || tokenInfo.Name == "" {
		return sr.SetError(err).Out(), errors.Errorf("token %s not supported", args.SourceToken)
	}

	if actionNumber <= 1 && !isActionSuccess {
		transferOptions, err := GetTransferOptions(
			ctx,
			args.Amount,
			tokenInfo.Decimals,
			args.SourceChain,
			args.TargetChain,
			common.HexToAddress(tokenInfo.ContractAddress),
		)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "get transfer options")
		}
		if _, ok := BRIDGES[BridgeType(transferOptions._bridge)]; !ok {
			err = errors.Errorf("bridge %s not supported", transferOptions._bridge)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
		}
		log.Debugf("transfer options: %+v", transferOptions)
		bridge := BRIDGES[BridgeType(transferOptions._bridge)](Options{
			SourceTokenName: args.SourceToken,
			TargetTokenName: args.TargetToken,
			SourceChain:     args.SourceChain,
			TargetChain:     args.TargetChain,
			Config:          b.config,
			Sender:          wallet.GetAddress(true),
			Recipient:       common.HexToAddress(args.Receiver),
			Amount:          args.Amount,
		})
		tx, err := bridge.Do(ctx, transferOptions)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "do transfer")
		}
		tx.Gas = 406775
		sr = sr.SetReciever(args.Receiver)

		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			OrderId:         args.OrderId,
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
		recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		txHash, err = wallet.SendTransaction(ctx, tx, ethClient)
		if err != nil {
			recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send signed transaction")
		}
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
		recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 2 {
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusPending).Out())
		log.Debugf("wait for tx %s", txHash.Hex())
		if err := wallet.WaitTransaction(ctx, txHash, ethClient); err != nil {
			recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for tx")
		}
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusSuccess).Out())
	}

	var record HistoryRecord
	if actionNumber <= 3 {
		recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusPending).Out())

		realHash, err := wallet.GetRealHash(ctx, txHash, ethClient)
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return result.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "get real hash error")
		}

		record, err = b.WaitForBridge(ctx, wallet.GetAddress(true), realHash)
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "wait for bridge")
		}
		log.Debugf("bridge result: %+v", record)
		if record.Result != 3 {
			err = errors.Errorf("bridge failed, result: %d", record.Result)
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		sr = sr.SetOrder(record).SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetChain)
		recordFn(sh.SetActions(targetChainReceivedAction).SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetToken).Out(), err)
	}
	return sr.Out(), nil
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
			args.Sender.GetAddress(true).Hex(), args.Amount, args.SourceChainNames...)
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

func (b *Bridge) Type() configs.ProviderType {
	return configs.Bridge
}
