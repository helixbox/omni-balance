package token_price

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"strings"
)

type Binance struct {
}

type BinancePriceResp struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

func (b Binance) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	if len(tokens) == 0 {
		return nil, errors.Errorf("tokens is empty")
	}
	var (
		symbol2Name = make(map[string]string)
		tokenNames  = SourceTokens2Map(tokens)
	)

	for _, v := range tokenNames {
		piarsName := strings.ToUpper(fmt.Sprintf("%sUSDT", v.Name))
		symbol2Name[piarsName] = v.Name
	}
	var resp []BinancePriceResp
	if err := utils.Request(ctx, "GET", "https://api.binance.com/api/v3/ticker/price", nil, &resp); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}
	var result []TokenPriceData
	for _, v := range resp {
		if _, ok := symbol2Name[v.Symbol]; !ok {
			continue
		}
		if v.Price.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			TokeName: symbol2Name[v.Symbol],
			Price:    v.Price,
		})
	}
	return AppendUsdtPrice(result), nil
}

func (b Binance) Name() string {
	return "binance"
}
