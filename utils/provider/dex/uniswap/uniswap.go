package uniswap

import (
	"context"
	"math/big"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	uniswapConfigs "omni-balance/utils/provider/dex/uniswap/configs"
	"strings"
	"time"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	SwapApproveSendingAction = "SwapApproveSending"
	SwapTXSendingAction      = "sourceChainSending"
	SwapTXReceivedAction     = "targetChainReceived"
)

func action2Int(action string) int {
	switch action {
	case SwapApproveSendingAction:
		return 0
	case SwapTXSendingAction:
		return 1
	case SwapTXReceivedAction:
		return 2
	default:
		return -1
	}
}

type Uniswap struct {
	conf configs.Config
}

func NewUniswap(conf configs.Config, _ ...bool) (provider.Provider, error) {
	return &Uniswap{conf: conf}, nil
}

func (u *Uniswap) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if !args.Sender.IsSupportEip712() {
		return nil, error_types.ErrUnsupportedWalletType
	}
	if strings.EqualFold(uniswapConfigs.GetContractAddress(args.TargetChain).UniversalRouter.Hex(), constant.ZeroAddress.Hex()) {
		return nil, error_types.ErrUnsupportedTokenAndChain
	}
	return u.GetTokenIns(ctx, args.TargetChain, args.TargetToken, args.Sender.GetAddress(true), args.Amount)
}

func (u *Uniswap) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	var (
		wallet       = args.Sender
		contract     = uniswapConfigs.GetContractAddress(args.TargetChain)
		chain        = u.conf.GetChainConfig(args.TargetChain)
		tokenOut     = u.conf.GetTokenInfoOnChain(args.TargetToken, args.TargetChain)
		amountOut    = args.Amount
		amountOutWei = decimal.NewFromBigInt(chains.EthToWei(amountOut, tokenOut.Decimals), 0)
		actionNumber = action2Int(args.LastHistory.Actions)
		recordFn     = func(s provider.SwapHistory, errs ...error) {
			s.ProviderType = string(u.Type())
			s.ProviderName = u.Name()
			s.Amount = args.Amount
			if args.RecordFn == nil {
				return
			}
			args.RecordFn(s, errs...)
		}
	)
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.TargetChain)
	if !utils.InArray(args.LastHistory.Actions,
		[]string{SwapApproveSendingAction, SwapTXReceivedAction, SwapTXReceivedAction, ""}) {
		return provider.SwapResult{}, error_types.ErrUnsupportedActions
	}

	tokenIns, err := u.GetTokenIns(ctx, chain.Name, tokenOut.Name, wallet.GetAddress(true), amountOut)
	if err != nil {
		return provider.SwapResult{}, errors.Wrap(err, "get token ins")
	}
	if len(tokenIns) == 0 {
		return provider.SwapResult{}, error_types.ErrUnsupportedTokenAndChain
	}
	tokenInPrice, err := tokenIns.GetBest()
	if err != nil {
		return provider.SwapResult{}, errors.Wrap(err, "get best token in")
	}
	var (
		tokenIn         = u.conf.GetTokenInfoOnChain(tokenInPrice.TokenName, chain.Name)
		isTokenInNative = u.conf.IsNativeToken(chain.Name, tokenIn.Name)
	)

	quote, err := u.GetTokenExactOutputQuote(
		ctx,
		chain.Name,
		tokenIn,
		tokenOut,
		wallet.GetAddress(true),
		amountOutWei)
	if err != nil {
		return provider.SwapResult{}, errors.Wrap(err, "get quote")
	}
	if len(quote.Quote.Route) == 0 {
		return provider.SwapResult{}, errors.Errorf("no route found")
	}

	ethClient, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return provider.SwapResult{}, errors.New("dial rpc")
	}
	defer ethClient.Close()

	var (
		sr = new(provider.SwapResult).
			SetTokenInName(args.SourceToken).
			SetTokenInChainName(args.SourceChain).
			SetProviderName(u.Name()).
			SetProviderType(u.Type()).
			SetCurrentChain(args.SourceChain).
			SetTx(args.LastHistory.Tx).
			SetReciever(wallet.GetAddress(true).Hex())
		sh = &provider.SwapHistory{
			ProviderName: u.Name(),
			ProviderType: string(u.Type()),
			Amount:       args.Amount,
			CurrentChain: args.SourceChain,
			Tx:           args.LastHistory.Tx,
		}
		isActionSuccess = args.LastHistory.Status == provider.TxStatusSuccess.String()
	)

	if !isTokenInNative && actionNumber <= 0 && isActionSuccess {
		// uint256 MAX see https://docs.uniswap.org/contracts/permit2/overview
		amount := decimal.RequireFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
		log.Debugf("approve tokenIn amount: %s", amount)
		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			OrderId:         args.OrderId,
			TokenIn:         tokenIn.Name,
			TokenOut:        args.TargetToken,
			TokenInChain:    args.TargetChain,
			TokenOutChain:   args.TargetChain,
			ProviderName:    u.Name(),
			TokenInAmount:   args.Amount,
			TokenOutAmount:  args.Amount,
			TransactionType: provider.ApproveTransactionAction,
		})
		if err := chains.TokenApprove(ctx, chains.TokenApproveParams{
			ChainId:         int64(chain.Id),
			TokenAddress:    common.HexToAddress(tokenIn.ContractAddress),
			Owner:           wallet.GetAddress(true),
			SendTransaction: wallet.SendTransaction,
			WaitTransaction: wallet.WaitTransaction,
			Spender:         contract.Permit2,
			AmountWei:       amount,
			Client:          ethClient,
		}); err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
		}
	}

	log.Debug("build tx")
	txRawData, err := u.BuildTx(ctx, BuildTxParams{
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		Chain:    chain,
		Sender:   args.Sender,
		Quote:    quote,
		Deadline: big.NewInt(time.Now().Add(time.Hour * 10).Unix()),
		Client:   ethClient,
		Amount:   args.Amount,
	})
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
	}
	sr = sr.SetOrder(common.Bytes2Hex(txRawData))

	var value = decimal.Zero
	if isTokenInNative {
		value = decimal.RequireFromString(quote.Quote.Route[0][0].AmountIn).Add(decimal.RequireFromString(quote.Quote.GasUseEstimate))
	}

	txData := &types.LegacyTx{
		GasPrice: decimal.RequireFromString(quote.Quote.GasPriceWei).Mul(decimal.NewFromInt(2)).BigInt(),
		Gas:      uint64(decimal.RequireFromString(quote.Quote.GasUseEstimate).Mul(decimal.NewFromInt(2)).IntPart()),
		To:       &contract.UniversalRouter,
		Data:     txRawData,
		Value:    value.BigInt(),
	}

	var tx = args.LastHistory.Tx
	if actionNumber <= 1 && !isActionSuccess {
		recordFn(sh.SetStatus(provider.TxStatusPending).SetActions(SwapTXSendingAction).Out())
		txHash, err := wallet.SendTransaction(ctx, txData, ethClient)
		if err != nil {
			recordFn(sh.SetStatus(provider.TxStatusFailed).SetActions(SwapTXSendingAction).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		tx = txHash.Hex()
		sr = sr.SetTx(tx).SetOrderId(tx)
		sh = sh.SetTx(tx)

		recordFn(sh.SetActions(SwapTXSendingAction).SetStatus(provider.TxStatusSuccess).Out())
	}
	if tx == "" {
		err := errors.New("tx is empty")
		recordFn(sh.SetStatus(provider.TxStatusFailed).SetActions(SwapTXSendingAction).Out(), err)
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
	}
	if actionNumber <= 2 && !isActionSuccess {
		if err := wallet.WaitTransaction(ctx, common.HexToHash(tx), ethClient); err != nil {
			recordFn(sh.SetStatus(provider.TxStatusFailed).SetActions(SwapTXReceivedAction).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
		}
		recordFn(sh.SetStatus(provider.TxStatusSuccess).SetActions(SwapTXReceivedAction).Out())
	}
	return sr.SetStatus(provider.TxStatusSuccess).Out(), nil
}

func (u *Uniswap) Help() []string {
	return []string{"See https://app.uniswap.org/swap"}
}

func (u *Uniswap) Name() string {
	return "uniswap"
}

func (u *Uniswap) Type() configs.ProviderType {
	return configs.DEX
}

func (u *Uniswap) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	if tokenInChainName != tokenOutChainName {
		return false, nil
	}
	tokenInCosts, err := u.GetTokenIns(ctx, tokenInChainName, tokenName, constant.ZeroAddress, amount)
	return len(tokenInCosts) != 0, err
}
