package token_price

import (
	"context"
	"github.com/shopspring/decimal"
	"omni-balance/utils/configs"
	"strings"
)

var (
	tokenPriceProviders = map[string]TokenPrice{}
)

func init() {
	RegisterTokenPriceProvider(new(Binance).Name(), new(Binance))
	RegisterTokenPriceProvider(new(OKX).Name(), new(OKX))
	RegisterTokenPriceProvider(new(Gate).Name(), new(Gate))
	RegisterTokenPriceProvider(new(Bybit).Name(), new(Bybit))
	RegisterTokenPriceProvider(new(Mexc).Name(), new(Mexc))
	RegisterTokenPriceProvider(new(Bitget).Name(), new(Bitget))
}

type TokenPrice interface {
	GetTokenPriceInUSDT(ctx context.Context, tokens ...configs.SourceToken) ([]TokenPriceData, error)
	Name() string
}

type TokenPriceData struct {
	TokeName     string          `json:"toke_name"`
	Price        decimal.Decimal `json:"price"`
	TokenAddress string          `json:"token_address"`
}

func RegisterTokenPriceProvider(name string, provider TokenPrice) {
	tokenPriceProviders[name] = provider
}

func ListTokenPriceProviders() []TokenPrice {
	var providers []TokenPrice
	for name := range tokenPriceProviders {
		providers = append(providers, tokenPriceProviders[name])
	}
	return providers
}

func SourceTokens2Map(tokens []configs.SourceToken) (result map[string]configs.SourceToken) {
	result = make(map[string]configs.SourceToken)
	for index, token := range tokens {
		if strings.EqualFold(token.Name, "USDT") {
			continue
		}
		result[token.Name] = tokens[index]
	}
	return result
}

func AppendUsdtPrice(data []TokenPriceData) []TokenPriceData {
	var (
		result  []TokenPriceData
		hasUsdt bool
	)
	for _, token := range data {
		if strings.EqualFold(token.TokeName, "USDT") {
			hasUsdt = true
			result = append(result, TokenPriceData{
				TokeName:     "USDT",
				Price:        decimal.NewFromInt(1),
				TokenAddress: token.TokenAddress,
			})
			continue
		}
		result = append(result, token)
	}
	if !hasUsdt {
		result = append(result, TokenPriceData{
			TokeName: "USDT",
			Price:    decimal.NewFromInt(1),
		})
	}
	return result
}
