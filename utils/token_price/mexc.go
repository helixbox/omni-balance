package token_price

import (
	"context"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"strings"
)

type Mexc struct {
}

type mexcticker struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

func (m Mexc) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	var tokenNames = SourceTokens2Map(tokens)

	var tickers []mexcticker
	err := utils.Request(ctx, "GET", "https://api.mexc.com/api/v3/ticker/price", nil, &tickers)
	if err != nil {
		return nil, err
	}
	var result []TokenPriceData
	for _, ticker := range tickers {
		name := strings.TrimRight(ticker.Symbol, "USDT")
		if _, ok := tokenNames[name]; !ok {
			continue
		}
		if ticker.Price.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			TokeName: name,
			Price:    ticker.Price,
		})
	}
	return AppendUsdtPrice(result), nil
}

func (m Mexc) Name() string {
	return "mexc"
}
