package token_price

import (
	"context"
	"github.com/shopspring/decimal"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider/cex/gate"
	"strings"
)

type Gate struct {
}

func NewGatePrice() TokenPrice {
	return &Gate{}
}

func (gp *Gate) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	g, err := gate.New(configs.Config{}, true)
	if err != nil {
		return nil, err
	}
	var (
		pairs      []string
		tokenNames = SourceTokens2Map(tokens)
	)
	for _, v := range tokenNames {
		pairs = append(pairs, strings.ToUpper(v.Name+"_USDT"))
	}

	tickerList, err := g.(*gate.Gate).Tickers(pairs...)
	if err != nil {
		return nil, err
	}
	var result []TokenPriceData
	for _, v := range tickerList {
		if v.Price.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			TokeName: v.TokenName,
			Price:    v.Price,
		})
	}
	return AppendUsdtPrice(result), nil
}

func (g *Gate) Name() string {
	return "gate.io"
}
