package li

import (
	"context"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"strings"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func init() {
	provider.Register(configs.Bridge, New)
}

type Li struct {
	conf configs.Config
}

func New(conf configs.Config, _ ...bool) (provider.Provider, error) {
	return Li{
		conf: conf,
	}, nil
}

func (l Li) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	var (
		err        error
		costAmount decimal.Decimal
	)
	if args.SourceChain == "" || args.SourceToken == "" {
		args.SourceToken, args.SourceChain, costAmount, _, err = l.GetBestQuote(ctx, args)
		if err != nil {
			return nil, err
		}
	}
	if args.SourceChain == "" || args.SourceToken == "" || costAmount.IsZero() {
		return nil, error_types.ErrUnsupportedTokenAndChain
	}
	return provider.TokenInCosts{
		provider.TokenInCost{
			TokenName:  args.SourceToken,
			CostAmount: costAmount,
		},
	}, nil
}

func (l Li) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	var (
		tokenIn      = l.conf.GetTokenInfoOnChain(tokenName, tokenInChainName)
		tokenOut     = l.conf.GetTokenInfoOnChain(tokenName, tokenOutChainName)
		tokenInChain = l.conf.GetChainConfig(tokenInChainName)
		tokenOuChain = l.conf.GetChainConfig(tokenOutChainName)
		fromAddress  = common.HexToAddress("0x91A5B1E18d8838989CDCd32ff26bc3B606e9B857")
	)
	if len(l.conf.Wallets) != 0 {
		fromAddress = l.conf.GetWallet(l.conf.Wallets[0].Address).GetAddress(true)
	}

	_, err := l.Quote(ctx, QuoteParams{
		FromAddress:   fromAddress,
		FromChainId:   tokenInChain.Id,
		ToChainId:     tokenOuChain.Id,
		FromToken:     common.HexToAddress(tokenIn.ContractAddress),
		ToToken:       common.HexToAddress(tokenOut.ContractAddress),
		FromAmountWei: decimal.NewFromBigInt(chains.EthToWei(amount, tokenIn.Decimals), 0),
	})
	return err == nil, err
}

func (l Li) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	var (
		targetChain       = l.conf.GetChainConfig(args.TargetChain)
		tokenOut          = l.conf.GetTokenInfoOnChain(args.TargetToken, targetChain.Name)
		tokenOutAmount    = args.Amount
		tokenOutAmountWei = decimal.NewFromBigInt(chains.EthToWei(tokenOutAmount, tokenOut.Decimals), 0)
		tokenIn           configs.Token
		tokenInAmount     decimal.Decimal
		tokenInAmountWei  decimal.Decimal
		costAmount        decimal.Decimal
		history           = args.LastHistory
		tx                = history.Tx
		actionNumber      = Action2Int(history.Actions)
		isActionSuccess   = history.Status == provider.TxStatusSuccess.String()
		sourceChain       configs.Chain
		quote             Quote
	)

	args.SourceToken, args.SourceChain, costAmount, quote, err = l.GetBestQuote(ctx, args)
	if err != nil {
		return provider.SwapResult{}, err
	}

	if args.SourceChain == "" || args.SourceToken == "" {
		log.Fatalf("#%d %s source chain and token is required", args.OrderId, args.TargetToken)
	}

	tokenIn = l.conf.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
	tokenInAmount = costAmount
	tokenInAmountWei = decimal.NewFromBigInt(chains.EthToWei(tokenInAmount, tokenIn.Decimals), 0)
	sourceChain = l.conf.GetChainConfig(args.SourceChain)
	isTokenInNative := l.conf.IsNativeToken(sourceChain.Name, tokenIn.Name)

	quote, err = l.Quote(ctx, QuoteParams{
		FromChainId:   sourceChain.Id,
		ToChainId:     constant.GetChainId(args.TargetChain),
		FromToken:     common.HexToAddress(tokenIn.ContractAddress),
		ToToken:       common.HexToAddress(tokenOut.ContractAddress),
		FromAmountWei: tokenInAmountWei,
		FromAddress:   args.Sender.GetAddress(true),
		ToAddress:     common.HexToAddress(args.Receiver),
	})
	if err != nil {
		log.Debugf("get quote error: %s", err)
		return
	}

	if quote.Estimate.ToAmountMin.IsZero() {
		return provider.SwapResult{}, errors.New("token out amount is zero")
	}

	client, err := chains.NewTryClient(ctx, sourceChain.RpcEndpoints)
	if err != nil {
		return provider.SwapResult{}, err
	}
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, sourceChain.Name)

	var (
		sr = new(provider.SwapResult).
			SetTokenInName(tokenIn.Name).
			SetTokenInChainName(args.SourceChain).
			SetProviderName(l.Name()).
			SetProviderType(l.Type()).
			SetCurrentChain(args.SourceChain).
			SetTx(args.LastHistory.Tx).
			SetReciever(args.Receiver)
		sh = &provider.SwapHistory{
			ProviderName: l.Name(),
			ProviderType: string(l.Type()),
			Amount:       args.Amount,
			CurrentChain: args.SourceChain,
			Tx:           history.Tx,
		}
	)

	if !isTokenInNative && actionNumber <= 1 && !isActionSuccess {
		log.Debugf("#%d %s is not native token, need approve", args.OrderId, tokenIn.Name)
		args.RecordFn(sh.SetActions(ApproveTransactionAction).SetStatus(provider.TxStatusPending).Out())
		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			OrderId:         args.OrderId,
			TokenIn:         tokenIn.Name,
			TokenOut:        tokenOut.Name,
			TokenInChain:    args.SourceChain,
			TokenOutChain:   args.TargetChain,
			ProviderName:    l.Name(),
			TokenInAmount:   tokenInAmount,
			TokenOutAmount:  tokenOutAmount,
			TransactionType: provider.ApproveTransactionAction,
		})
		err = chains.TokenApprove(ctx,
			chains.TokenApproveParams{
				ChainId:         int64(sourceChain.Id),
				TokenAddress:    common.HexToAddress(tokenIn.ContractAddress),
				Owner:           args.Sender.GetAddress(true),
				SendTransaction: args.Sender.SendTransaction,
				WaitTransaction: args.Sender.WaitTransaction,
				Spender:         common.HexToAddress(quote.Estimate.ApprovalAddress),
				// for save next gas, multiply 2
				AmountWei:   tokenInAmountWei.Mul(decimal.RequireFromString("2")),
				IsNotWaitTx: l.conf.GetWalletConfig(string(args.Sender.GetAddress().Hex())).MultiSignType != "",
				Client:      client,
			})
		if err != nil {
			args.RecordFn(sh.SetActions(ApproveTransactionAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		log.Debugf("#%d %s approve transaction success", args.OrderId, tokenIn.Name)
		args.RecordFn(sh.SetActions(ApproveTransactionAction).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 2 && (!isActionSuccess || actionNumber == 1) {
		amount := args.Amount.Copy()
		args.Amount = tokenInAmount
		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			OrderId:         args.OrderId,
			Receiver:        common.HexToAddress(args.Receiver),
			TokenIn:         tokenIn.Name,
			TokenOut:        tokenOut.Name,
			TokenInChain:    args.SourceChain,
			TokenOutChain:   args.TargetChain,
			ProviderName:    l.Name(),
			TokenInAmount:   tokenInAmount,
			TokenOutAmount:  tokenOutAmount,
			TransactionType: provider.SwapTransactionAction,
		})
		txn := quote.TransactionRequest
		value := decimal.RequireFromString(utils.HexToString(txn.Value))
		if common.HexToAddress(quote.Action.ToToken.Address).Cmp(common.HexToAddress(tokenOut.ContractAddress)) != 0 {
			err = errors.Errorf("token out address is not match, expect: %s, actual: %s",
				tokenOut.ContractAddress, quote.Action.ToToken.Address)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "build tx error")
		}
		if common.HexToAddress(quote.Action.FromToken.Address).Cmp(common.HexToAddress(tokenIn.ContractAddress)) != 0 {
			err = errors.Errorf("token in address is not match, expect: %s, actual: %s",
				tokenIn.ContractAddress, quote.Action.FromToken.Address)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "build tx error")
		}

		args.Amount = amount
		if quote.Estimate.ToAmountMin.Div(tokenOutAmountWei).LessThan(decimal.RequireFromString("0.5")) {
			err = errors.Errorf("minmum receive is too low, minmum receive: %s, amount: %s",
				quote.Estimate.ToAmountMin, tokenOutAmountWei)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		if isTokenInNative && value.IsZero() {
			err = errors.Errorf("tokenin is native token, but value is zero")
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		if isTokenInNative && value.GreaterThan(tokenInAmountWei.Mul(decimal.RequireFromString("1.5"))) {
			err = errors.Errorf("tx value is too high, tx value: %s, amount: %s", value, tokenInAmountWei)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		sr.SetOrder(quote)
		args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		log.Debugf("#%d sending tx on chain", args.OrderId)
		txHash, err := args.Sender.SendTransaction(ctx, &types.LegacyTx{
			To:    &txn.To,
			Value: value.BigInt(),
			Data:  common.Hex2Bytes(strings.TrimPrefix(txn.Data, "0x")),
		}, client)
		if err != nil {
			args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "send tx error")
		}
		log.Debugf("#%d sending tx on chain success: %s", args.OrderId, txHash.Hex())
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
		args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusSuccess).Out())
		tx = txHash.Hex()
	}

	args.RecordFn(sh.SetStatus(provider.TxStatusPending).Out())
	log.Debugf("#%d waiting for tx on chain", args.OrderId)
	if err := args.Sender.WaitTransaction(ctx, common.HexToHash(tx), client); err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrapf(err, "wait tx %s error", tx)
	}
	realHash, err := args.Sender.GetRealHash(ctx, common.HexToHash(tx), client)
	if err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "get real hash error")
	}
	log.Debugf("#%d waiting for tx in li: %s", args.OrderId, realHash)
	if err := l.WaitForTx(ctx, realHash); err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "wait li error")
	}
	log.Debugf("#%d waiting for tx success in li: %s", args.OrderId, realHash)
	args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusSuccess).SetCurrentChain(targetChain.Name).Out())
	return sr.SetStatus(provider.TxStatusSuccess).SetCurrentChain(targetChain.Name).Out(), nil
}

func (l Li) Help() []string {
	return []string{"see https://jumper.exchange"}
}

func (l Li) Name() string {
	return "li"
}

func (l Li) Type() configs.ProviderType {
	return configs.Bridge
}
