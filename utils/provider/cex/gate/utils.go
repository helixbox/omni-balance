package gate

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/gateio/gateapi-go/v6"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"net/http"
	"omni-balance/utils"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"
	"strings"
	"time"
)

var (
	gateChain2StandardName = map[string]string{
		"ETH":     constant.Ethereum,
		"ARBEVM":  constant.Arbitrum,
		"BSC":     constant.Bsc,
		"KAVAEVM": constant.Kava,
		"MATIC":   constant.Polygon,
		"OPETH":   constant.Optimism,
		"CELO":    constant.Celo,
		"OPBNB":   constant.OpBNB,
		"BNB":     constant.Bnb,
	}
	gateTokenName2StandardName = map[string]string{
		"ARBEVM": "ARB",
	}
)

type Ticker struct {
	Low                      string  `json:"low"`
	Volume                   string  `json:"volume"`
	Last                     string  `json:"last"`
	Open                     string  `json:"open"`
	Deal                     string  `json:"deal"`
	Close                    string  `json:"close"`
	Change                   string  `json:"change"`
	High                     string  `json:"high"`
	Result                   string  `json:"result"`
	Avg                      float64 `json:"avg"`
	RateChangePercentage     string  `json:"rate_change_percentage"`
	RateChangePercentageUtc0 int     `json:"rate_change_percentage_utc0"`
	RateChangePercentageUtc8 int     `json:"rate_change_percentage_utc8"`
}

type TickerResult struct {
	TokenName string
	Price     decimal.Decimal
}

func TokenName2GateTokenName(tokenName string) string {
	for k, v := range gateTokenName2StandardName {
		if strings.EqualFold(tokenName, v) {
			return k
		}
	}
	return tokenName
}

func ChainName2GateChainName(chainName string) string {
	for k, v := range gateChain2StandardName {
		if strings.EqualFold(chainName, v) {
			return k
		}
	}
	return chainName
}

func (g *Gate) ticker(pairs string) (TickerResult, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://data.gateapi.io/api/1/ticker/%s", pairs), nil)
	if err != nil {
		return TickerResult{}, err
	}
	req.Header.Set("content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return TickerResult{}, err
	}
	defer resp.Body.Close()
	var result = new(Ticker)
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return TickerResult{}, err
	}
	if !strings.EqualFold(result.Result, "true") {
		return TickerResult{}, errors.Errorf("pairs %s not found", pairs)
	}
	pairsList := strings.Split(pairs, "_")
	return TickerResult{
		TokenName: pairsList[1],
		Price:     decimal.RequireFromString(result.Last),
	}, nil
}

func (g *Gate) Tickers(pairs ...string) (result []TickerResult, err error) {
	for _, v := range pairs {
		ticker, err := g.ticker(v)
		if err != nil {
			return nil, errors.Wrap(err, "ticker")
		}
		result = append(result, ticker)
	}
	return result, nil
}

func (g *Gate) GetTokenOutChain(ctx context.Context, token string) ([]string, error) {
	chains, _, err := g.client.WalletApi.ListCurrencyChains(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "list currency chains")
	}
	var supportedChains []string
	for _, v := range chains {
		if _, ok := gateChain2StandardName[v.Chain]; !ok {
			logrus.Debugf("chain %s not supported in %s", v.Chain, token)
			continue
		}
		if v.IsWithdrawDisabled == 1 {
			continue
		}
		supportedChains = append(supportedChains, gateChain2StandardName[v.Chain])
	}
	if len(supportedChains) == 0 {
		return nil, error_types.ErrUnsupportedTokenAndChain
	}
	return supportedChains, nil
}

func (g *Gate) buyToken(ctx context.Context, tokenIn provider.TokenInCost, targetToken string, amount decimal.Decimal,
	f func(gateapi.Order) bool) (order gateapi.Order, err error) {

	currencyPair := utils.GetCurrencyPair(targetToken, "_", tokenIn.TokenName)
	order = gateapi.Order{
		Text:         "t-ob",
		CurrencyPair: currencyPair,
		Type:         "limit",
		Side:         "buy",
		Price:        tokenIn.CostAmount.Div(amount).String(),
		Amount:       amount.String(),
	}
	logrus.WithFields(logrus.Fields{
		"order": order,
	}).Debugf("buy token")

	// place order
	order, _, err = g.client.SpotApi.CreateOrder(ctx, order)
	if err != nil {
		return order, errors.Wrap(err, "create order")
	}
	if order.Status == "cancelled" {
		return order, errors.Errorf("order status is %s", order.Status)
	}
	// wait for order filled
	for order.Status == "open" {
		order, _, err = g.client.SpotApi.GetOrder(ctx, order.Id, fmt.Sprintf("%s_%s", tokenIn.TokenName, targetToken), nil)
		if err != nil {
			return order, errors.Wrap(err, "get order")
		}
		if order.Status == "closed" {
			break
		}
		if f(order) {
			break
		}
		time.Sleep(time.Second)
	}
	if order.Status == "cancelled" {
		return order, errors.Errorf("order status is %s", order.Status)
	}
	return order, nil
}

func (g *Gate) IsVerifiedAddress(ctx context.Context, address, tokenName, chainName string) (bool, error) {
	verifiedAddress, _, err := g.client.WalletApi.ListSavedAddress(ctx, tokenName,
		&gateapi.ListSavedAddressOpts{Chain: optional.NewString(ChainName2GateChainName(chainName))})
	if err != nil {
		return false, errors.Wrap(err, "list saved address")
	}
	if len(verifiedAddress) == 0 {
		return false, errors.Errorf("%s not in gate.io verified address for %s on %s", address, tokenName, chainName)
	}
	var (
		isSenderInVerifiedAddress bool
	)
	for _, v := range verifiedAddress {
		if v.Verified == "0" {
			continue
		}
		if !strings.EqualFold(v.Address, address) ||
			!strings.EqualFold(v.Chain, ChainName2GateChainName(chainName)) {
			continue
		}
		isSenderInVerifiedAddress = true
		break
	}
	return isSenderInVerifiedAddress, nil
}
