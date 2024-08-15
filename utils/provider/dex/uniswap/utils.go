package uniswap

import (
	"context"
	"math/big"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"omni-balance/utils/provider/dex/uniswap/abi/permit2"
	uniswapConfigs "omni-balance/utils/provider/dex/uniswap/configs"
	uniswapUtils "omni-balance/utils/provider/dex/uniswap/utils"
	"omni-balance/utils/wallets"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type BuildTxParams struct {
	TokenIn, TokenOut configs.Token
	Chain             configs.Chain
	Sender            wallets.Wallets
	Quote             uniswapUtils.Quote
	Deadline          *big.Int
	Client            simulated.Client
	Amount            decimal.Decimal
}

func (u *Uniswap) GetTokenExactOutputQuote(ctx context.Context, chainName string, tokenIn, tokenOut configs.Token,
	sender common.Address, amount decimal.Decimal) (uniswapUtils.Quote, error) {
	var tokenInAddress, tokenOutAddress = tokenIn.ContractAddress, tokenOut.ContractAddress
	if u.conf.IsNativeToken(chainName, tokenIn.Name) {
		tokenInAddress = tokenIn.Name
	}
	if u.conf.IsNativeToken(chainName, tokenOut.Name) {
		tokenOutAddress = tokenOut.Name
	}

	return uniswapUtils.GetTokenExactOutputQuote(ctx, uniswapUtils.GetTokenQuoteParams{
		ChainName: chainName,
		TokenIn:   tokenInAddress,
		TokenOut:  tokenOutAddress,
		Amount:    amount,
		Sender:    sender,
	})
}

func (u *Uniswap) BuildTx(ctx context.Context, args BuildTxParams) ([]byte, error) {
	var (
		isTokenInNative  = u.conf.IsNativeToken(args.Chain.Name, args.TokenIn.Name)
		isTokenOutNative = u.conf.IsNativeToken(args.Chain.Name, args.TokenOut.Name)
		chainId          = int64(args.Chain.Id)
		wallet           = args.Sender
		contract         = uniswapConfigs.GetContractAddress(args.Chain.Name)
	)
	universalRouterExecute := new(uniswapUtils.Execute)
	if isTokenInNative {
		universalRouterExecute = universalRouterExecute.WrapEth(uniswapConfigs.WETH_RECIPIENT_ADDRESS,
			decimal.RequireFromString(args.Quote.Quote.Route[0][0].AmountIn).BigInt(), true)
	}

	if args.Quote.Quote.PermitData.Domain.ChainId != 0 && !isTokenInNative {
		p, err := permit2.NewPermit2(contract.Permit2, args.Client)
		if err != nil {
			return nil, errors.Wrap(err, "permit2")
		}
		allowance, err := p.Allowance(
			nil, wallet.GetAddress(true), common.HexToAddress(args.TokenIn.ContractAddress),
			contract.UniversalRouter)
		if err != nil {
			return nil, errors.Wrap(err, "permit2 allowance")
		}
		if allowance.Amount.Cmp(decimal.RequireFromString(args.Quote.Quote.Quote).BigInt()) == -1 {
			universalRouterExecute = universalRouterExecute.Permit2Permit(uniswapUtils.PermitSingle{
				Details: uniswapUtils.PermitDetails{
					Token:      common.HexToAddress(args.TokenIn.ContractAddress),
					Amount:     decimal.RequireFromString(args.Quote.Quote.Quote).Mul(decimal.NewFromInt(2)).BigInt(),
					Expiration: big.NewInt(args.Deadline.Int64() + 1*24*60*60),
					Nonce:      allowance.Nonce,
				},
				Spender:           contract.UniversalRouter,
				SigDeadline:       big.NewInt(args.Deadline.Int64() + 1*24*60*60),
				VerifyingContract: contract.Permit2,
				ChainId:           math.NewHexOrDecimal256(chainId),
			}, wallet.SignRawMessage)
		}
	}

	var routerPool = make(map[string][]uniswapUtils.Route)
	for index, routers := range args.Quote.Quote.Route {
		for routerIndex, v := range routers {
			routerPool[v.Type] = append(routerPool[v.Type], args.Quote.Quote.Route[index][routerIndex])
		}
	}
	for poolType, routers := range routerPool {
		var paths uniswapUtils.Paths
		for _, router := range routers {
			p := &uniswapUtils.Path{
				TokenIn:  common.HexToAddress(router.TokenIn.Address),
				TokenOut: common.HexToAddress(router.TokenOut.Address),
			}
			if poolType == uniswapConfigs.V3Pool {
				p.Fee = decimal.RequireFromString(router.Fee).IntPart()
			}
			paths = append(paths, *p)
		}
		switch poolType {
		case uniswapConfigs.V3Pool:
			universalRouterExecute = universalRouterExecute.V3SwapExactOut(
				uniswapConfigs.WETH_RECIPIENT_ADDRESS,
				decimal.RequireFromString(routers[len(routers)-1].AmountOut).BigInt(),
				decimal.RequireFromString(routers[0].AmountIn).BigInt(),
				paths,
				!isTokenInNative, false,
			)
		case uniswapConfigs.V2Pool:
			universalRouterExecute = universalRouterExecute.V2SwapExactOut(
				uniswapConfigs.WETH_RECIPIENT_ADDRESS,
				decimal.RequireFromString(routers[len(routers)-1].AmountOut).BigInt(),
				decimal.RequireFromString(routers[0].AmountIn).BigInt(),
				paths,
				!isTokenInNative, false,
			)
		}
	}
	if !isTokenOutNative {
		universalRouterExecute = universalRouterExecute.Sweep(
			common.HexToAddress(args.TokenOut.ContractAddress),
			uniswapConfigs.UNWRAP_WETH_RECIPIENT_ADDRESS,
			decimal.RequireFromString(args.Quote.Quote.Amount).BigInt(),
			true,
		)
	}

	if isTokenInNative {
		universalRouterExecute = universalRouterExecute.UnwrapWEth(
			uniswapConfigs.UNWRAP_WETH_RECIPIENT_ADDRESS,
			big.NewInt(0),
			true,
		)
	}
	if isTokenOutNative {
		universalRouterExecute = universalRouterExecute.UnwrapWEth(
			uniswapConfigs.UNWRAP_WETH_RECIPIENT_ADDRESS,
			chains.EthToWei(args.Amount, args.TokenIn.Decimals),
			true,
		)
	}
	return universalRouterExecute.Build(args.Deadline)
}

func (u *Uniswap) GetTokenIns(ctx context.Context, chainName, tokenOutName string, wallet common.Address,
	tokenOutAmount decimal.Decimal) (tokenInCosts provider.TokenInCosts, err error) {

	tokenOut := u.conf.GetTokenInfoOnChain(tokenOutName, chainName)
	chain := u.conf.GetChainConfig(chainName)
	client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		log.Warnf("get %s chain client error: %s", chainName, err)
		return
	}
	defer client.Close()
	for _, v := range u.conf.GetSourceTokenNamesByChain(chainName) {
		tokenIn := u.conf.GetTokenInfoOnChain(v, chainName)
		result, err := u.GetTokenExactOutputQuote(
			ctx,
			chainName,
			tokenIn,
			tokenOut,
			wallet,
			decimal.NewFromBigInt(chains.EthToWei(tokenOutAmount, tokenOut.Decimals), 0))
		if errors.Is(err, error_types.ErrInsufficientLiquidity) {
			continue
		}
		if err != nil {
			log.Warnf("get quote error: %s", err)
			continue
		}

		tokenInBalance, err := chains.GetTokenBalance(ctx, client, tokenIn.ContractAddress, wallet.Hex(), tokenIn.Decimals)
		if err != nil {
			log.Warnf("get balance error: %s", err)
			continue
		}
		if tokenInBalance.Cmp(decimal.RequireFromString(result.Quote.QuoteDecimals)) < 0 {
			continue
		}
		tokenInCosts = append(tokenInCosts, provider.TokenInCost{
			TokenName:  v,
			CostAmount: decimal.RequireFromString(result.Quote.QuoteDecimals),
		})
	}
	return tokenInCosts, nil
}
