package token_price

import (
	"context"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"strings"
)

type OKX struct {
}

type marketResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		InstId  string          `json:"instId"`
		IdxPx   decimal.Decimal `json:"idxPx"`
		High24H decimal.Decimal `json:"high24h"`
		SodUtc0 decimal.Decimal `json:"sodUtc0"`
		Open24H decimal.Decimal `json:"open24h"`
		Low24H  decimal.Decimal `json:"low24h"`
		SodUtc8 decimal.Decimal `json:"sodUtc8"`
		Ts      string          `json:"ts"`
	} `json:"data"`
}

func (O OKX) GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error) {
	var tokenNames = SourceTokens2Map(tokens)
	var data marketResp
	if err := utils.Request(ctx, "GET", "https://www.okx.com/api/v5/market/index-tickers?quoteCcy=USDT", nil, &data); err != nil {
		return nil, errors.Wrap(err, "decode response")
	}

	var result []TokenPriceData
	for _, token := range data.Data {
		tokenName := strings.Split(token.InstId, "-")[0]
		if _, ok := tokenNames[tokenName]; !ok {
			continue
		}
		if token.IdxPx.LessThan(decimal.Zero) {
			continue
		}
		result = append(result, TokenPriceData{
			Price:    token.IdxPx,
			TokeName: tokenName,
		})
	}
	return AppendUsdtPrice(result), nil
}

func (O OKX) Name() string {
	return "okx"
}
