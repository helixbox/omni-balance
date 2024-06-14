package token_price

import (
	"context"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"strings"
)

type Bitget struct {
}

type bitgetMarket struct {
	Code        string `json:"code"`
	Msg         string `json:"msg"`
	RequestTime int64  `json:"requestTime"`
	Data        []struct {
		Open         string          `json:"open"`
		Symbol       string          `json:"symbol"`
		High24H      string          `json:"high24h"`
		Low24H       string          `json:"low24h"`
		LastPr       decimal.Decimal `json:"lastPr"`
		QuoteVolume  string          `json:"quoteVolume"`
		BaseVolume   string          `json:"baseVolume"`
		UsdtVolume   string          `json:"usdtVolume"`
		Ts           string          `json:"ts"`
		BidPr        string          `json:"bidPr"`
		AskPr        string          `json:"askPr"`
		BidSz        string          `json:"bidSz"`
		AskSz        string          `json:"askSz"`
		OpenUtc      string          `json:"openUtc"`
		ChangeUtc24H string          `json:"changeUtc24h"`
		Change24H    string          `json:"change24h"`
	} `json:"data"`
}

func (b Bitget) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	var (
		tokenNames = SourceTokens2Map(tokens)
	)
	var bitgetMarket bitgetMarket
	err := utils.Request(ctx, "GET", "https://api.bitget.com/api/v2/spot/market/tickers", nil, &bitgetMarket)
	if err != nil {
		return nil, err
	}
	var result []TokenPriceData
	for _, token := range bitgetMarket.Data {
		name := strings.TrimRight(token.Symbol, "USDT")
		if _, ok := tokenNames[name]; !ok {
			continue
		}
		if token.LastPr.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			TokeName: name,
			Price:    token.LastPr,
		})
	}
	return AppendUsdtPrice(result), nil
}

func (b Bitget) Name() string {
	return "bitget"
}
