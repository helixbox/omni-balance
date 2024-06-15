package uniswap

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
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
		log          = utils.GetLogFromCtx(ctx).WithField("provide_name", u.Name())
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
	log = log.WithFields(logrus.Fields{
		"tokenIn":         utils.ToMap(tokenIn),
		"tokenOut":        utils.ToMap(tokenOut),
		"amountOut":       amountOut,
		"isTokenInNative": isTokenInNative,
	})
	log.Debugf("start get exact output quote")
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

	log.WithField("quote", utils.ToMap(quote.Quote)).Debug("get quote success")

	ethClient, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return provider.SwapResult{}, errors.New("dial rpc")
	}
	defer ethClient.Close()

	if !isTokenInNative && actionNumber <= 0 && args.LastHistory.Status != provider.TxStatusSuccess.String() {
		// uint256 MAX see https://docs.uniswap.org/contracts/permit2/overview
		amount := decimal.RequireFromString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
		log.Debugf("approve tokenIn amount: %s", amount)
		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
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
			return provider.SwapResult{}, err
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
		return provider.SwapResult{}, err
	}

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
	swapResult := &provider.SwapResult{
		ProviderType: u.Type(),
		ProviderName: u.Name(),
		Order:        quote.Quote,
		CurrentChain: chain.Name,
		TokenInName:  tokenIn.Name,
	}
	log.WithField("txData", utils.ToMap(txData)).Debug("build tx success")
	var tx = args.LastHistory.Tx
	if actionNumber <= 1 && args.LastHistory.Status != provider.TxStatusSuccess.String() {
		recordFn(provider.SwapHistory{Actions: SwapTXSendingAction, Status: string(provider.TxStatusPending), CurrentChain: chain.Name})
		log.Debug("sending tx")
		txHash, err := wallet.SendTransaction(ctx, txData, ethClient)
		if err != nil {
			swapResult.Status = provider.TxStatusFailed
			swapResult.Error = err.Error()
			args.RecordFn(provider.SwapHistory{Actions: SwapTXSendingAction, Status: string(provider.TxStatusFailed), CurrentChain: chain.Name, Amount: args.Amount})
			return *swapResult, errors.Wrap(err, "send tx")
		}
		tx = txHash.Hex()
		log.WithField("txHash", tx).Debug("send tx success")
		recordFn(provider.SwapHistory{Actions: SwapTXSendingAction, Status: string(provider.TxStatusPending), CurrentChain: chain.Name, Tx: tx})
	}
	if tx == "" {
		return provider.SwapResult{}, errors.New("tx is empty")
	}
	swapResult.Tx = tx
	swapResult.OrderId = tx
	if actionNumber <= 2 && args.LastHistory.Status != provider.TxStatusSuccess.String() {
		log.Debug("waiting tx")
		if err := wallet.WaitTransaction(ctx, common.HexToHash(tx), ethClient); err != nil {
			recordFn(provider.SwapHistory{Actions: SwapTXReceivedAction, Status: string(provider.TxStatusFailed), CurrentChain: args.SourceChain, Tx: tx})
			swapResult.Status = provider.TxStatusFailed
			swapResult.Error = err.Error()
			return *swapResult, err
		}
		swapResult.Status = provider.TxStatusSuccess
		recordFn(provider.SwapHistory{Actions: SwapTXReceivedAction, Status: string(provider.TxStatusSuccess), CurrentChain: chain.Name, Tx: tx})
	}
	log.Debug("tx success")
	return *swapResult, nil
}

func (u *Uniswap) Help() []string {
	return []string{"See https://app.uniswap.org/swap"}
}

func (u *Uniswap) Name() string {
	return "uniswap"
}

func (u *Uniswap) Type() configs.LiquidityProviderType {
	return configs.DEX
}

func (u *Uniswap) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	if tokenInChainName != tokenOutChainName {
		return false, nil
	}
	tokenInCosts, err := u.GetTokenIns(ctx, tokenInChainName, tokenName, constant.ZeroAddress, amount)
	return len(tokenInCosts) != 0, err
}
