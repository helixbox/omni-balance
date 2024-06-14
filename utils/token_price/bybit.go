package token_price

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"strings"
)

type Bybit struct {
}

type bybitMarketResp struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string `json:"category"`
		List     []struct {
			Symbol       string          `json:"symbol"`
			Bid1Price    string          `json:"bid1Price"`
			Bid1Size     string          `json:"bid1Size"`
			Ask1Price    string          `json:"ask1Price"`
			Ask1Size     string          `json:"ask1Size"`
			LastPrice    decimal.Decimal `json:"lastPrice"`
			PrevPrice24H string          `json:"prevPrice24h"`
			Price24HPcnt string          `json:"price24hPcnt"`
			HighPrice24H string          `json:"highPrice24h"`
			LowPrice24H  string          `json:"lowPrice24h"`
			Turnover24H  string          `json:"turnover24h"`
			Volume24H    string          `json:"volume24h"`
		} `json:"list"`
	} `json:"result"`
}

func (b Bybit) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	var (
		tokenNames = SourceTokens2Map(tokens)
	)
	url := "https://api.bybit.com/v5/market/tickers?category=spot&baseCoin=USDT"
	var data bybitMarketResp
	if err := utils.Request(ctx, "GET", url, nil, &data); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}
	var result []TokenPriceData
	for _, token := range data.Result.List {
		name := strings.TrimRight(token.Symbol, "USDT")
		if _, ok := tokenNames[name]; !ok {
			continue
		}
		if token.LastPrice.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			TokeName:     name,
			Price:        token.LastPrice,
			TokenAddress: "",
		})
	}
	return AppendUsdtPrice(result), nil
}

func (b Bybit) Name() string {
	return "bybit"
}
