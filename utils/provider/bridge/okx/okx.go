package okx

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

type OKX struct {
	conf configs.Config
	Key  Config
}

func New(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	o := new(OKX)
	if len(noInit) > 0 && noInit[0] {
		return o, nil
	}
	if err := conf.GetProvidersConfig(o.Name(), o.Type(), &o.Key); err != nil {
		return nil, err
	}
	o.conf = conf

	return o, nil
}

func (o *OKX) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	var (
		err        error
		costAmount decimal.Decimal
	)

	args.SourceToken, args.SourceChain, costAmount, err = o.GetBestTokenInChain(ctx, args)
	if err != nil {
		return nil, err
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

func (o *OKX) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	tokenOut := o.conf.GetTokenInfoOnChain(tokenName, tokenOutChainName)
	tokenIn := o.conf.GetTokenInfoOnChain(tokenName, tokenInChainName)
	amountWei := decimal.NewFromBigInt(chains.EthToWei(amount, tokenIn.Decimals), 0)
	quote, err := o.Quote(ctx, QuoteParams{
		Amount:           amountWei,
		FormChainId:      constant.GetChainId(tokenInChainName),
		ToChainId:        constant.GetChainId(tokenOutChainName),
		ToTokenAddress:   common.HexToAddress(tokenOut.ContractAddress),
		FromTokenAddress: common.HexToAddress(tokenIn.ContractAddress),
	})
	if err != nil {
		return false, err
	}
	if len(quote.RouterList) == 0 {
		return false, error_types.ErrUnsupportedTokenAndChain
	}
	return true, nil
}

func (o *OKX) Swap(ctx context.Context, args provider.SwapParams) (provider.SwapResult, error) {
	var (
		err               error
		targetChain       = o.conf.GetChainConfig(args.TargetChain)
		tokenOut          = o.conf.GetTokenInfoOnChain(args.TargetToken, targetChain.Name)
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
	)

	args.SourceToken, args.SourceChain, costAmount, err = o.GetBestTokenInChain(ctx, args)
	if err != nil {
		return provider.SwapResult{}, err
	}
	if args.SourceChain == "" || args.SourceToken == "" {
		log.Fatalf("#%d %s source chain and token is required", args.OrderId, args.TargetToken)
	}

	tokenIn = o.conf.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
	tokenInAmount = costAmount
	tokenInAmountWei = decimal.NewFromBigInt(chains.EthToWei(tokenInAmount, tokenIn.Decimals), 0)
	sourceChain = o.conf.GetChainConfig(args.SourceChain)
	isTokenInNative := o.conf.IsNativeToken(sourceChain.Name, tokenIn.Name)
	quote, err := o.Quote(ctx, QuoteParams{
		Amount:           tokenInAmountWei,
		FormChainId:      sourceChain.Id,
		ToChainId:        targetChain.Id,
		ToTokenAddress:   common.HexToAddress(tokenOut.ContractAddress),
		FromTokenAddress: common.HexToAddress(tokenIn.ContractAddress),
	})
	if err != nil {
		return provider.SwapResult{}, err
	}
	tokenOutAmount = chains.WeiToEth(quote.RouterList[0].ToTokenAmount.BigInt(), tokenOut.Decimals)
	if tokenOutAmount.LessThanOrEqual(decimal.Zero) {
		return provider.SwapResult{}, errors.New("token out amount is zero")
	}
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, sourceChain.Name)

	client, err := chains.NewTryClient(ctx, sourceChain.RpcEndpoints)
	if err != nil {
		return provider.SwapResult{}, err
	}
	var (
		sr = new(provider.SwapResult).
			SetTokenInName(tokenIn.Name).
			SetTokenInChainName(args.SourceChain).
			SetProviderName(o.Name()).
			SetProviderType(o.Type()).
			SetCurrentChain(args.SourceChain).
			SetTx(args.LastHistory.Tx).
			SetReciever(args.Receiver)
		sh = &provider.SwapHistory{
			ProviderName: o.Name(),
			ProviderType: string(o.Type()),
			Amount:       args.Amount,
			CurrentChain: args.SourceChain,
			Tx:           history.Tx,
		}
	)
	if !isTokenInNative && actionNumber <= 1 && !isActionSuccess {
		log.Debugf("#%d %s is not native token, need approve", args.OrderId, tokenIn.Name)
		args.RecordFn(sh.SetStatus(provider.TxStatusPending).SetActions(ApproveTransactionAction).Out())
		approveTransaction, err := o.approveTransaction(ctx, sourceChain.Id,
			common.HexToAddress(tokenIn.ContractAddress), tokenInAmountWei)
		if err != nil {
			err = errors.Wrap(err, "get approve transaction from okx error")
			args.RecordFn(sh.SetStatus(provider.TxStatusFailed).SetActions(ApproveTransactionAction).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		log.Debugf("get approve transaction from okx, the spender is %s", approveTransaction.DexContractAddress)
		err = chains.TokenApprove(ctx, chains.TokenApproveParams{
			ChainId:         int64(sourceChain.Id),
			TokenAddress:    common.HexToAddress(tokenIn.ContractAddress),
			Owner:           args.Sender.GetAddress(true),
			SendTransaction: args.Sender.SendTransaction,
			WaitTransaction: args.Sender.WaitTransaction,
			Spender:         common.HexToAddress(approveTransaction.DexContractAddress),
			// for save next gas, multiply 2
			AmountWei:   tokenInAmountWei.Mul(decimal.RequireFromString("1.02")),
			IsNotWaitTx: o.conf.GetWalletConfig(string(args.Sender.GetAddress().Hex())).MultiSignType != "",
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
			ProviderName:    o.Name(),
			TokenInAmount:   tokenInAmount,
			TokenOutAmount:  tokenOutAmount,
			TransactionType: provider.SwapTransactionAction,
			CurrentBalance:  args.CurrentBalance,
		})
		buildTx, err := o.buildTx(ctx, args)
		if err != nil {
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "build tx error")
		}
		sr = sr.SetOrder(buildTx)
		args.Amount = amount
		if buildTx.ToTokenAmount.Div(tokenOutAmountWei).LessThan(decimal.RequireFromString("0.5")) {
			err = errors.Errorf("minmum receive is too low, minmum receive: %s, amount: %s",
				buildTx.ToTokenAmount, tokenOutAmountWei)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		if isTokenInNative && buildTx.Tx.Value.IsZero() {
			err = errors.Errorf("tokenin is native token, but value is zero")
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		if isTokenInNative && buildTx.Tx.Value.GreaterThan(tokenInAmountWei.Mul(decimal.RequireFromString("1.5"))) {
			err = errors.Errorf("tx value is too high, tx value: %s, amount: %s", buildTx.Tx.Value, tokenInAmountWei)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}

		args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		txHash, err := args.Sender.SendTransaction(ctx, &types.DynamicFeeTx{
			To:    &buildTx.Tx.To,
			Value: buildTx.Tx.Value.BigInt(),
			Data:  common.Hex2Bytes(strings.TrimPrefix(buildTx.Tx.Data, "0x")),
		}, client)
		if err != nil {
			args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "send tx error")
		}
		log.Debugf("#%d %s sending tx on chain success", args.OrderId, tokenIn.Name)
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
		args.RecordFn(sh.SetActions(SourceChainSendingAction).SetStatus(provider.TxStatusSuccess).SetTx(txHash.Hex()).Out())
		tx = txHash.Hex()
	}
	args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusPending).Out())
	log.Debugf("#%d %s waiting for tx on chain", args.OrderId, tokenIn.Name)
	if err := args.Sender.WaitTransaction(ctx, common.HexToHash(tx), client); err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "wait tx error")
	}

	realHash, err := args.Sender.GetRealHash(ctx, common.HexToHash(tx), client)
	if err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "get real hash error")
	}
	log.Debugf("#%d %s waiting for tx in okx", args.OrderId, tokenIn.Name)
	if err := o.WaitForTx(ctx, realHash, sourceChain.Id); err != nil {
		args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusFailed).Out(), err)
		return sr.SetError(err).SetStatus(provider.TxStatusFailed).Out(), errors.Wrap(err, "wait okx error")
	}
	log.Debugf("#%d %s waiting for tx success in okx", args.OrderId, tokenIn.Name)
	args.RecordFn(sh.SetActions(WaitForTxAction).SetStatus(provider.TxStatusSuccess).SetCurrentChain(targetChain.Name).Out())
	return sr.SetCurrentChain(args.TargetChain).SetStatus(provider.TxStatusSuccess).Out(), nil
}

func (o *OKX) Help() []string {
	var result []string
	for _, v := range utils.ExtractTagFromStruct(&Config{}, "yaml", "help") {
		result = append(result, v["yaml"]+": "+v["help"])
	}
	result = append(result, "You can get these configs from https://www.okx.com/zh-hans/web3/build/docs/waas/introduction-to-developer-portal-interface")
	return result
}

func (o *OKX) Name() string {
	return "okx"
}

func (o *OKX) Type() configs.ProviderType {
	return configs.Bridge
}
