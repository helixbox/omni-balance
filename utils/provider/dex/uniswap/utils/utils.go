package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
)

type GetTokenQuoteParams struct {
	ChainName string
	TokenIn   string
	TokenOut  string
	// wei
	Amount decimal.Decimal
	Sender common.Address
}

type TokenQuote struct {
	TokenInChainId     int                `json:"tokenInChainId"`
	TokenIn            string             `json:"tokenIn"`
	TokenOutChainId    int                `json:"tokenOutChainId"`
	TokenOut           string             `json:"tokenOut"`
	Amount             string             `json:"amount"`
	SendPortionEnabled bool               `json:"sendPortionEnabled"`
	Type               string             `json:"type"`
	Intent             string             `json:"intent"`
	Configs            []TokenQuoteConfig `json:"configs"`
	UseUniswapX        bool               `json:"useUniswapX"`
	Swapper            string             `json:"swapper"`
}

type TokenQuoteConfig struct {
	UseSyntheticQuotes             bool     `json:"useSyntheticQuotes,omitempty"`
	Swapper                        string   `json:"swapper,omitempty"`
	RoutingType                    string   `json:"routingType"`
	EnableUniversalRouter          bool     `json:"enableUniversalRouter,omitempty"`
	Protocols                      []string `json:"protocols,omitempty"`
	Recipient                      string   `json:"recipient,omitempty"`
	EnableFeeOnTransferFeeFetching bool     `json:"enableFeeOnTransferFeeFetching,omitempty"`
}

type Quote struct {
	Routing string `json:"routing"`
	Quote   struct {
		MethodParameters struct {
			Calldata string `json:"calldata"`
			Value    string `json:"value"`
			To       string `json:"to"`
		} `json:"methodParameters"`
		BlockNumber                        string    `json:"blockNumber"`
		Amount                             string    `json:"amount"`
		AmountDecimals                     string    `json:"amountDecimals"`
		Quote                              string    `json:"quote"`
		QuoteDecimals                      string    `json:"quoteDecimals"`
		QuoteGasAdjusted                   string    `json:"quoteGasAdjusted"`
		QuoteGasAdjustedDecimals           string    `json:"quoteGasAdjustedDecimals"`
		QuoteGasAndPortionAdjusted         string    `json:"quoteGasAndPortionAdjusted"`
		QuoteGasAndPortionAdjustedDecimals string    `json:"quoteGasAndPortionAdjustedDecimals"`
		GasUseEstimateQuote                string    `json:"gasUseEstimateQuote"`
		GasUseEstimateQuoteDecimals        string    `json:"gasUseEstimateQuoteDecimals"`
		GasUseEstimate                     string    `json:"gasUseEstimate"`
		GasUseEstimateUSD                  string    `json:"gasUseEstimateUSD"`
		SimulationStatus                   string    `json:"simulationStatus"`
		SimulationError                    bool      `json:"simulationError"`
		GasPriceWei                        string    `json:"gasPriceWei"`
		Route                              [][]Route `json:"route"`
		RouteString                        string    `json:"routeString"`
		QuoteId                            string    `json:"quoteId"`
		HitsCachedRoutes                   bool      `json:"hitsCachedRoutes"`
		PortionBips                        int       `json:"portionBips"`
		PortionRecipient                   string    `json:"portionRecipient"`
		PortionAmount                      string    `json:"portionAmount"`
		PortionAmountDecimals              string    `json:"portionAmountDecimals"`
		RequestId                          string    `json:"requestId"`
		PermitData                         struct {
			Domain struct {
				Name              string `json:"name"`
				ChainId           int    `json:"chainId"`
				VerifyingContract string `json:"verifyingContract"`
			} `json:"domain"`
			Types struct {
				PermitSingle []struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"PermitSingle"`
				PermitDetails []struct {
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"PermitDetails"`
			} `json:"types"`
			Values struct {
				Details struct {
					Token      string `json:"token"`
					Amount     string `json:"amount"`
					Expiration string `json:"expiration"`
					Nonce      string `json:"nonce"`
				} `json:"details"`
				Spender     string `json:"spender"`
				SigDeadline string `json:"sigDeadline"`
			} `json:"values"`
		} `json:"permitData"`
		TradeType string  `json:"tradeType"`
		Slippage  float64 `json:"slippage"`
	} `json:"quote"`
	RequestId string `json:"requestId"`
	AllQuotes []struct {
		Routing string `json:"routing"`
		Quote   struct {
			MethodParameters struct {
				Calldata string `json:"calldata"`
				Value    string `json:"value"`
				To       string `json:"to"`
			} `json:"methodParameters"`
			BlockNumber                        string `json:"blockNumber"`
			Amount                             string `json:"amount"`
			AmountDecimals                     string `json:"amountDecimals"`
			Quote                              string `json:"quote"`
			QuoteDecimals                      string `json:"quoteDecimals"`
			QuoteGasAdjusted                   string `json:"quoteGasAdjusted"`
			QuoteGasAdjustedDecimals           string `json:"quoteGasAdjustedDecimals"`
			QuoteGasAndPortionAdjusted         string `json:"quoteGasAndPortionAdjusted"`
			QuoteGasAndPortionAdjustedDecimals string `json:"quoteGasAndPortionAdjustedDecimals"`
			GasUseEstimateQuote                string `json:"gasUseEstimateQuote"`
			GasUseEstimateQuoteDecimals        string `json:"gasUseEstimateQuoteDecimals"`
			GasUseEstimate                     string `json:"gasUseEstimate"`
			GasUseEstimateUSD                  string `json:"gasUseEstimateUSD"`
			SimulationStatus                   string `json:"simulationStatus"`
			SimulationError                    bool   `json:"simulationError"`
			GasPriceWei                        string `json:"gasPriceWei"`
			Route                              [][]struct {
				Type    string `json:"type"`
				Address string `json:"address"`
				TokenIn struct {
					ChainId    int    `json:"chainId"`
					Decimals   string `json:"decimals"`
					Address    string `json:"address"`
					Symbol     string `json:"symbol"`
					BuyFeeBps  string `json:"buyFeeBps,omitempty"`
					SellFeeBps string `json:"sellFeeBps,omitempty"`
				} `json:"tokenIn"`
				TokenOut struct {
					ChainId    int    `json:"chainId"`
					Decimals   string `json:"decimals"`
					Address    string `json:"address"`
					Symbol     string `json:"symbol"`
					BuyFeeBps  string `json:"buyFeeBps,omitempty"`
					SellFeeBps string `json:"sellFeeBps,omitempty"`
				} `json:"tokenOut"`
				Reserve0 struct {
					Token struct {
						ChainId    int    `json:"chainId"`
						Decimals   string `json:"decimals"`
						Address    string `json:"address"`
						Symbol     string `json:"symbol"`
						BuyFeeBps  string `json:"buyFeeBps"`
						SellFeeBps string `json:"sellFeeBps"`
					} `json:"token"`
					Quotient string `json:"quotient"`
				} `json:"reserve0"`
				Reserve1 struct {
					Token struct {
						ChainId  int    `json:"chainId"`
						Decimals string `json:"decimals"`
						Address  string `json:"address"`
						Symbol   string `json:"symbol"`
					} `json:"token"`
					Quotient string `json:"quotient"`
				} `json:"reserve1"`
				AmountIn  string `json:"amountIn,omitempty"`
				AmountOut string `json:"amountOut,omitempty"`
			} `json:"route"`
			RouteString           string     `json:"routeString"`
			QuoteId               string     `json:"quoteId"`
			HitsCachedRoutes      bool       `json:"hitsCachedRoutes"`
			PortionBips           int        `json:"portionBips"`
			PortionRecipient      string     `json:"portionRecipient"`
			PortionAmount         string     `json:"portionAmount"`
			PortionAmountDecimals string     `json:"portionAmountDecimals"`
			RequestId             string     `json:"requestId"`
			PermitData            PermitData `json:"permitData"`
			TradeType             string     `json:"tradeType"`
			Slippage              float64    `json:"slippage"`
		} `json:"quote"`
	} `json:"allQuotes"`
}

type Route struct {
	Type    string `json:"type"`
	Address string `json:"address"`
	TokenIn struct {
		ChainId    int    `json:"chainId"`
		Decimals   string `json:"decimals"`
		Address    string `json:"address"`
		Symbol     string `json:"symbol"`
		BuyFeeBps  string `json:"buyFeeBps,omitempty"`
		SellFeeBps string `json:"sellFeeBps,omitempty"`
	} `json:"tokenIn"`
	TokenOut struct {
		ChainId    int    `json:"chainId"`
		Decimals   string `json:"decimals"`
		Address    string `json:"address"`
		Symbol     string `json:"symbol"`
		BuyFeeBps  string `json:"buyFeeBps,omitempty"`
		SellFeeBps string `json:"sellFeeBps,omitempty"`
	} `json:"tokenOut"`
	Reserve0 struct {
		Token struct {
			ChainId    int    `json:"chainId"`
			Decimals   string `json:"decimals"`
			Address    string `json:"address"`
			Symbol     string `json:"symbol"`
			BuyFeeBps  string `json:"buyFeeBps"`
			SellFeeBps string `json:"sellFeeBps"`
		} `json:"token"`
		Quotient string `json:"quotient"`
	} `json:"reserve0"`
	Reserve1 struct {
		Token struct {
			ChainId  int    `json:"chainId"`
			Decimals string `json:"decimals"`
			Address  string `json:"address"`
			Symbol   string `json:"symbol"`
		} `json:"token"`
		Quotient string `json:"quotient"`
	} `json:"reserve1"`
	Fee          string `json:"fee"`
	SqrtRatioX96 string `json:"sqrtRatioX96,omitempty"`
	AmountIn     string `json:"amountIn,omitempty"`
	AmountOut    string `json:"amountOut,omitempty"`
}

type PermitData struct {
	Domain struct {
		Name              string `json:"name"`
		ChainId           int    `json:"chainId"`
		VerifyingContract string `json:"verifyingContract"`
	} `json:"domain"`
	Types struct {
		PermitSingle []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"PermitSingle"`
		PermitDetails []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"PermitDetails"`
	} `json:"types"`
	Values struct {
		Details struct {
			Token      string `json:"token"`
			Amount     string `json:"amount"`
			Expiration string `json:"expiration"`
			Nonce      string `json:"nonce"`
		} `json:"details"`
		Spender     string `json:"spender"`
		SigDeadline string `json:"sigDeadline"`
	} `json:"values"`
}

func GetTokenExactOutputQuote(ctx context.Context, args GetTokenQuoteParams) (Quote, error) {
	chainId := constant.GetChainId(args.ChainName)
	body, err := json.Marshal(TokenQuote{
		TokenInChainId:     chainId,
		TokenIn:            args.TokenIn,
		TokenOutChainId:    chainId,
		TokenOut:           args.TokenOut,
		Amount:             args.Amount.String(),
		SendPortionEnabled: true,
		Type:               "EXACT_OUTPUT",
		Intent:             "quote",
		Configs: []TokenQuoteConfig{
			{
				RoutingType: "DUTCH_LIMIT",
				Swapper:     args.Sender.Hex(),
			},
			{
				EnableFeeOnTransferFeeFetching: true,
				EnableUniversalRouter:          true,
				Protocols:                      []string{"V2", "V3", "MIXED"},
				Recipient:                      args.Sender.Hex(),
				RoutingType:                    "CLASSIC",
			},
		},
		Swapper: args.Sender.Hex(),
	})
	if err != nil {
		return Quote{}, err
	}
	req, err := http.NewRequest("POST", "https://interface.gateway.uniswap.org/v2/quote", bytes.NewReader(body))
	if err != nil {
		return Quote{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	// must set origin
	req.Header.Set("Origin", "https://bafybeidjiiinp4v64dyircmic5w3lti4sj7e6jd37siispbgtytxx37gai.ipfs.dweb.link")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Quote{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return Quote{}, error_types.ErrInsufficientLiquidity
	}
	if resp.StatusCode != 200 {
		data, _ := io.ReadAll(resp.Body)
		return Quote{}, fmt.Errorf("status code: %d, error: %s", resp.StatusCode, data)
	}
	var result Quote
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Quote{}, err
	}
	return result, nil
}

func AbiPack(types []string, values ...interface{}) ([]byte, error) {
	if len(types) != len(values) || len(types) == 0 {
		return nil, fmt.Errorf("types and values must be same length")
	}
	var args abi.Arguments
	for _, v := range types {
		args = append(args, abi.Argument{
			Type: chains.MostNewAbiType(v, v, nil),
		})
	}
	if args == nil {
		return nil, fmt.Errorf("values must not be empty")
	}
	return args.Pack(values...)
}
