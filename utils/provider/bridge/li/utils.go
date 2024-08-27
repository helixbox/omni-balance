package li

import (
	"context"
	"net/url"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"strconv"
	"time"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

// Quote see https://apidocs.li.fi/reference/get_quote
func (l Li) Quote(ctx context.Context, args QuoteParams) (Quote, error) {
	var params = &url.Values{}
	params.Set("fromChain", strconv.Itoa(args.FromChainId))
	params.Set("toChain", strconv.Itoa(args.ToChainId))
	params.Set("fromToken", args.FromToken.Hex())
	params.Set("toToken", args.ToToken.Hex())
	params.Set("fromAmount", args.FromAmountWei.String())
	params.Set("fromAddress", args.FromAddress.Hex())
	if args.ToAddress.Cmp(constant.ZeroAddress) != 0 {
		params.Set("toAddress", args.ToAddress.Hex())
	}
	params.Set("order", "RECOMMENDED")
	params.Set("slippage", "0.01")
	var result Quote
	err := utils.Request(ctx, "GET", "https://li.quest/v1/quote?"+params.Encode(), nil, &result)
	if err != nil {
		return result, err
	}
	return result, result.Error()
}

func (l Li) GetBestQuote(ctx context.Context, args provider.SwapParams) (tokenInName, tokenInChainName string,
	tokenInAmount decimal.Decimal, quote Quote, err error) {
	if args.TargetToken == "" || args.TargetChain == "" {
		return tokenInName, tokenInChainName, tokenInAmount, Quote{}, errors.New("target token or target chain is empty")
	}

	var (
		tokenOut = l.conf.GetTokenInfoOnChain(args.TargetToken, args.TargetChain)
	)

	getQuote := func(chainName, tokenName string) {
		sourceToken := l.conf.GetTokenInfoOnChain(tokenName, chainName)
		chain := l.conf.GetChainConfig(chainName)
		tokenIn := l.conf.GetTokenInfoOnChain(sourceToken.Name, chainName)
		if tokenIn.ContractAddress == "" {
			log.Warnf("tokenIn contract address is empty on %s", chainName)
			return
		}
		client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
		if err != nil {
			log.Warnf("get chain %s client error: %s", chain.Name, err)
			return
		}
		defer client.Close()
		balance, err := chains.GetTokenBalance(ctx, client, tokenIn.ContractAddress,
			args.Sender.GetAddress(true).Hex(), tokenIn.Decimals)
		if err != nil {
			log.Warnf("get %s on %s balance error: %s", tokenName, chainName, err)
			return
		}
		if balance.LessThanOrEqual(decimal.Zero) {
			log.Debugf("%s on %s balance is 0", tokenName, chainName)
			return
		}
		quoteData, err := l.Quote(ctx, QuoteParams{
			FromChainId:   constant.GetChainId(args.TargetChain),
			ToChainId:     chain.Id,
			FromToken:     common.HexToAddress(tokenOut.ContractAddress),
			ToToken:       common.HexToAddress(tokenIn.ContractAddress),
			FromAmountWei: decimal.NewFromBigInt(chains.EthToWei(args.Amount, tokenOut.Decimals), 0),
			FromAddress:   args.Sender.GetAddress(true),
			ToAddress:     common.HexToAddress(args.Receiver),
		})
		if err != nil {
			log.Debugf("get tokenin %s on %s quote error: %s", tokenName, chainName, err)
			return
		}

		needTokenInAmount := chains.WeiToEth(quoteData.Estimate.ToAmount.BigInt(), tokenIn.Decimals)

		log.Debugf("get quote, need %s %s on %s to get %s %s on %s",
			needTokenInAmount, tokenName, chainName, args.Amount, tokenOut.Name, args.TargetChain)
		// 2% slippage
		needTokenInAmount = needTokenInAmount.Add(needTokenInAmount.Mul(decimal.RequireFromString("0.002")))

		// if needTokenInAmount.GreaterThan(balance) {
		// 	log.Debugf("%s need %s on %s balance is greater than balance, need: %s, balance: %s",
		// 		args.Sender.GetAddress(true).Hex(), tokenName, chainName, needTokenInAmount.String(), balance.String())
		// 	return
		// }
		if tokenInAmount.Equal(decimal.Zero) {
			tokenInAmount = needTokenInAmount
		}
		if tokenInAmount.GreaterThan(needTokenInAmount) {
			return
		}

		tokenInAmount = needTokenInAmount
		tokenInName = sourceToken.Name
		tokenInChainName = chainName
		quote = quoteData
		log.Debugf("get best token in chain, token: %s on %s, tokenInAmount: %s", tokenName, chainName, tokenInAmount.String())
	}

	if args.SourceChain != "" && args.SourceToken != "" {
		getQuote(args.SourceChain, args.SourceToken)
		if tokenInChainName == "" || tokenInName == "" || tokenInAmount.IsZero() {
			return "", "", tokenInAmount, Quote{}, error_types.ErrUnsupportedTokenAndChain
		}
		log.Debugf("#%d get best route for %s %s is use %s %s token from %s chain by specify source chain and token", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
		return
	}

	if args.SourceToken != "" && len(args.SourceChainNames) > 0 && args.SourceChain == "" {
		for _, v := range args.SourceChainNames {
			getQuote(v, args.SourceToken)
		}
		if tokenInChainName == "" || tokenInName == "" || tokenInAmount.IsZero() {
			return "", "", tokenInAmount, Quote{}, error_types.ErrUnsupportedTokenAndChain
		}
		log.Debugf("#%d best route for %s %s is use %s %s token from %s chain by specify source chains and token", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
		return
	}

	for _, sourceToken := range l.conf.SourceTokens {
		if args.SourceToken != "" && sourceToken.Name != args.SourceToken {
			continue
		}
		for _, v := range sourceToken.Chains {
			if args.SourceChain != "" && v != args.SourceChain {
				continue
			}
			getQuote(v, sourceToken.Name)
		}
	}
	if tokenInChainName == "" || tokenInName == "" || tokenInAmount.IsZero() {
		return "", "", tokenInAmount, Quote{}, error_types.ErrUnsupportedTokenAndChain
	}
	log.Debugf("#%d best route for %s %s is use %s %s token from %s chain", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
	return
}

func (l Li) WaitForTx(ctx context.Context, hash common.Hash) error {
	var (
		t     = time.NewTicker(time.Second * 10)
		count int64
	)
	defer t.Stop()
	for count < 600 {
		select {
		case <-ctx.Done():
			log.Debugf("wait for tx %s canceled", hash)
			return context.Canceled
		case <-t.C:
			log.Debugf("start check tx %s", hash)
			status, err := l.Status(ctx, hash)
			if errors.Is(err, error_types.ErrNotFound) {
				count++
				log.Debugf("tx %s not found, count: %d", hash, count)
				continue
			}

			log.Infof("tx %s status: %s, substatus: %s", hash, status.Status, status.Substatus)
			if status.Status == "" {
				log.Debugf("tx %s status is empty, try it 1 minute later", hash)
				time.Sleep(time.Minute)
				continue
			}
			switch status.Status {
			case "PENDING":
				count = 0
				continue
			case "DONE":
				if status.Substatus != "COMPLETED" {
					return errors.New(subStatus2Text[status.Substatus])
				}
				return nil
			case "FAILED":
				return errors.New(subStatus2Text[status.Substatus])
			default:
				count++
				continue
			}
		}
	}
	return errors.New("wait for tx timeout")
}

func (l Li) Status(ctx context.Context, hash common.Hash) (Status, error) {
	var params = &url.Values{}
	params.Set("txHash", hash.Hex())
	var status Status
	err := utils.Request(ctx, "GET", "https://li.quest/v1/status?"+params.Encode(), nil, &status)
	if err != nil {
		return Status{}, err
	}
	if err := status.Error(); err != nil {
		return Status{}, err
	}
	return status, nil
}
