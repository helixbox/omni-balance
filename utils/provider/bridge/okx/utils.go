package okx

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/url"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"strconv"
	"strings"
	"time"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type QuoteParams struct {
	Amount           decimal.Decimal
	FormChainId      int
	ToChainId        int
	ToTokenAddress   common.Address
	FromTokenAddress common.Address
}

func (o *OKX) Quote(ctx context.Context, args QuoteParams) (QuoteData, error) {
	var params = url.Values{}
	params.Set("amount", args.Amount.String())
	params.Set("fromChainId", strconv.Itoa(args.FormChainId))
	params.Set("toChainId", strconv.Itoa(args.ToChainId))
	params.Set("toTokenAddress", StandardizeZeroAddress(args.ToTokenAddress).Hex())
	params.Set("fromTokenAddress", StandardizeZeroAddress(args.FromTokenAddress).Hex())
	// see https://www.okx.com/zh-hans/web3/build/docs/waas/dex-get-route-information
	params.Set("slippage", "0.01")
	var result OksResp
	if err := o.request(ctx, "GET", "/api/v5/dex/cross-chain/quote", params, nil, &result); err != nil {
		return QuoteData{}, errors.Wrap(err, "quote")
	}
	return result.Quote()
}

func (o *OKX) request(ctx context.Context, method string, path string, params url.Values, body map[string]interface{}, dest *OksResp) error {
	signature, timestamp := createSignature(method, path, params, body, o.Key.SecretKey)
	headers := map[string]string{
		"OK-ACCESS-KEY":        o.Key.ApiKey,
		"OK-ACCESS-SIGN":       signature,
		"OK-ACCESS-TIMESTAMP":  timestamp,
		"OK-ACCESS-PASSPHRASE": o.Key.Passphrase,
		"OK-ACCESS-PROJECT":    o.Key.Project,
		"content-type":         "application/json",
	}
	var headersList []string
	for k, v := range headers {
		headersList = append(headersList, k, v)
	}
	var bodyReader io.Reader
	if len(body) != 0 {
		data, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(data)
	}
	u, _ := url.Parse("https://www.okx.com")
	u.RawPath = path
	u.Path = path
	u.RawQuery = params.Encode()
	if err := utils.Request(ctx, method, u.String(), bodyReader, dest, headersList...); err != nil {
		return err
	}
	return dest.Error()
}

func (o *OKX) GetBestTokenInChain(ctx context.Context, args provider.SwapParams) (tokenInName, tokenInChainName string, tokenInAmount decimal.Decimal, err error) {
	if args.TargetToken == "" || args.TargetChain == "" {
		return "", "", tokenInAmount, errors.New("target token or target chain is empty")
	}
	var (
		tokenOut = o.conf.GetTokenInfoOnChain(args.TargetToken, args.TargetChain)
	)
	getQuote := func(chainName, tokenName string) {
		sourceToken := o.conf.GetTokenInfoOnChain(tokenName, chainName)
		chain := o.conf.GetChainConfig(chainName)
		tokenIn := o.conf.GetTokenInfoOnChain(sourceToken.Name, chainName)
		if tokenIn.ContractAddress == "" {
			log.Debugf("token in contract address is empty, chain: %s, token: %s", chainName, tokenName)
			return
		}
		client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
		if err != nil {
			log.Debugf("new client error: '%s' with %+v", err, chain.RpcEndpoints)
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
			return
		}

		quote, err := o.Quote(ctx, QuoteParams{
			Amount:           decimal.NewFromBigInt(chains.EthToWei(args.Amount, tokenOut.Decimals), 0),
			FormChainId:      constant.GetChainId(args.TargetChain),
			ToChainId:        chain.Id,
			ToTokenAddress:   common.HexToAddress(tokenIn.ContractAddress),
			FromTokenAddress: common.HexToAddress(tokenOut.ContractAddress),
		})
		if err != nil {
			log.Warnf("get %s on %s to %s on %s quote error: %s", tokenName, chainName, tokenOut.Name, args.TargetChain, err)
			return
		}
		if len(quote.RouterList) == 0 {
			log.Warnf("%s on %s to %s on %s no router list", tokenName, chainName, tokenOut.Name, args.TargetChain)
			return
		}
		tokenAmount := quote.RouterList[0].ToTokenAmount
		tokenOutAmount := chains.WeiToEth(tokenAmount.BigInt(), tokenIn.Decimals)
		// 0.2% slippage + 0.2% fee
		needTokenInAmount := tokenOutAmount.Add(tokenOutAmount.Mul(decimal.RequireFromString("0.004")))

		if needTokenInAmount.GreaterThan(balance) {
			log.Debugf("%s need %s on %s balance is greater than balance, need: %s, balance: %s",
				args.Sender.GetAddress(true).Hex(), tokenName, chainName, needTokenInAmount.String(), balance.String())
			return
		}
		if tokenInAmount.Equal(decimal.Zero) {
			tokenInAmount = needTokenInAmount
		}
		if tokenInAmount.GreaterThan(needTokenInAmount) {
			return
		}
		tokenInAmount = needTokenInAmount
		tokenInName = sourceToken.Name
		tokenInChainName = chainName
	}
	if args.SourceChain != "" && args.SourceToken != "" {
		getQuote(args.SourceChain, args.SourceToken)
		if tokenInChainName == "" || tokenInName == "" || tokenInAmount.IsZero() {
			return "", "", tokenInAmount, error_types.ErrUnsupportedTokenAndChain
		}
		log.Debugf("#%d get best route for %s %s is use %s %s token from %s chain by specify source chain and token", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
		return
	}

	if args.SourceToken != "" && len(args.SourceChainNames) > 0 && args.SourceChain == "" {
		for _, v := range args.SourceChainNames {
			getQuote(v, args.SourceToken)
		}
		if tokenInChainName == "" || tokenInName == "" || tokenInAmount.IsZero() {
			return "", "", tokenInAmount, error_types.ErrUnsupportedTokenAndChain
		}
		log.Debugf("#%d best route for %s %s is use %s %s token from %s chain by specify source chains and token", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
		return
	}

	for _, sourceToken := range o.conf.SourceTokens {
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
		return "", "", tokenInAmount, error_types.ErrUnsupportedTokenAndChain
	}
	log.Debugf("#%d best route for %s %s is use %s %s token from %s chain", args.OrderId, args.TargetChain, args.TargetToken, tokenInAmount, tokenInName, tokenInChainName)
	return
}

func (o *OKX) approveTransaction(ctx context.Context, chainId int, tokenAddress common.Address, amountWei decimal.Decimal) (ApproveTransaction, error) {
	params := &url.Values{}
	params.Set("chainId", strconv.Itoa(chainId))
	params.Set("tokenContractAddress", tokenAddress.Hex())
	params.Set("approveAmount", amountWei.String())
	var result OksResp
	if err := o.request(ctx, "GET", "/api/v5/dex/aggregator/approve-transaction", *params, nil, &result); err != nil {
		return ApproveTransaction{}, err
	}
	return result.ApproveTransaction()
}

func (o *OKX) buildTx(ctx context.Context, args provider.SwapParams) (BuildTxData, error) {
	var (
		tokenIn       = o.conf.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
		tokenOut      = o.conf.GetTokenInfoOnChain(args.TargetToken, args.TargetChain)
		tokenInChain  = o.conf.GetChainConfig(args.SourceChain)
		tokenOutChain = o.conf.GetChainConfig(args.TargetChain)
		amountWei     = chains.EthToWei(args.Amount, tokenIn.Decimals)
	)
	if tokenIn.Name == "" ||
		tokenOut.Name == "" ||
		tokenInChain.Name == "" ||
		tokenOutChain.Name == "" ||
		tokenIn.ContractAddress == "" ||
		tokenOut.ContractAddress == "" {
		return BuildTxData{}, errors.New("token info is empty")
	}
	params := &url.Values{}
	params.Set("fromChainId", strconv.Itoa(tokenInChain.Id))
	params.Set("toChainId", strconv.Itoa(tokenOutChain.Id))
	params.Set("fromTokenAddress", StandardizeZeroAddress(common.HexToAddress(tokenIn.ContractAddress)).Hex())
	params.Set("toTokenAddress", StandardizeZeroAddress(common.HexToAddress(tokenOut.ContractAddress)).Hex())
	params.Set("amount", amountWei.String())
	params.Set("slippage", "0.01")
	params.Set("sort", "1")
	params.Set("userWalletAddress", args.Sender.GetAddress(true).Hex())
	if args.Receiver != "" && !strings.EqualFold(args.Receiver, args.Sender.GetAddress().Hex()) {
		params.Set("receiveAddress", args.Receiver)
	}
	var tx OksResp
	if err := o.request(ctx, "GET", "/api/v5/dex/cross-chain/build-tx", *params, nil, &tx); err != nil {
		return BuildTxData{}, err
	}
	return tx.BuildTx()
}

func (o *OKX) TxStatus(ctx context.Context, hash common.Hash, sourceChainId int) (Status, error) {
	params := &url.Values{}
	params.Set("hash", hash.Hex())
	params.Set("chainId", strconv.Itoa(sourceChainId))
	var txStatus OksResp
	if err := o.request(ctx, "GET", "/api/v5/dex/cross-chain/status", *params, nil, &txStatus); err != nil {
		return Status{}, err
	}
	return txStatus.TxStatus()
}

func (o *OKX) WaitForTx(ctx context.Context, hash common.Hash, sourceChainId int) error {
	var (
		t     = time.NewTicker(time.Second * 2)
		count = 0
	)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return errors.New("context done")
		case <-t.C:
			s, err := o.TxStatus(ctx, hash, sourceChainId)
			if err != nil {
				return err
			}
			log.Infof("tx status: %s, detail status: %s", s.Status, okxWaitForTxStatus[s.DetailStatus])
			switch s.Status {
			case "PENDING":
				count = 0
				continue
			case "SUCCESS":
				return nil
			case "FAILURE", "REFUND":
				return errors.New(okxWaitForTxStatus[s.DetailStatus])
			default:
				if count >= 600 {
					return errors.Errorf("unknown status: %s, detail status: %s", s.Status, s.DetailStatus)
				}
				count++
			}
		}
	}
}

func sign(message, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func createSignature(method, requestPath string, params url.Values, body map[string]interface{}, secretKey string) (string, string) {
	timestamp := time.Now().UTC().Format(time.RFC3339)
	message := preHash(timestamp, method, requestPath, params, body)
	signature := sign(message, secretKey)
	return signature, timestamp
}

func preHash(timestamp, method, requestPath string, params url.Values, body map[string]interface{}) string {
	var queryString string
	if params != nil {
		queryString = "?" + params.Encode()
	}
	if body != nil {
		data, _ := json.Marshal(body)
		queryString = string(data)
	}
	return timestamp + method + requestPath + queryString
}
