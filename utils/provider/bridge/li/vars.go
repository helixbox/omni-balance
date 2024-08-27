package li

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	ApproveTransactionAction = "ApproveTransaction"
	SourceChainSendingAction = "sourceChainSending"
	WaitForTxAction          = "WaitForTx"

	subStatus2Text = map[string]string{
		"NOT_PROCESSABLE_REFUND_NEEDED": "The transfer cannot be completed; a refund is required.",
		"OUT_OF_GAS":                    "The transaction ran out of gas.",
		"SLIPPAGE_EXCEEDED":             "The slippage tolerance was exceeded.",
		"INSUFFICIENT_ALLOWANCE":        "The user has not approved the transfer of tokens.",
		"INSUFFICIENT_BALANCE":          "The user does not have enough balance.",
		"UNKNOWN_ERROR":                 "An unknown error occurred.",
		"EXPIRED":                       "The transaction expired.",
		"REFUNDED":                      "The transfer was not successful, and the sent token has been refunded.",
		"PARTIAL":                       "The transfer was partially successful. This can happen for specific bridges like across , hop, stargate or amarok which may provide alternative tokens in case of low liquidity.",
	}
)

type Resp struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Status struct {
	Resp
	TransactionId string `json:"transactionId"`
	Sending       struct {
		TxHash string `json:"txHash"`
		TxLink string `json:"txLink"`
		Amount string `json:"amount"`
		Token  struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"token"`
		ChainId  int    `json:"chainId"`
		GasPrice string `json:"gasPrice"`
		GasUsed  string `json:"gasUsed"`
		GasToken struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"gasToken"`
		GasAmount    string `json:"gasAmount"`
		GasAmountUSD string `json:"gasAmountUSD"`
		AmountUSD    string `json:"amountUSD"`
		Value        string `json:"value"`
		Timestamp    int    `json:"timestamp"`
	} `json:"sending"`
	Receiving struct {
		TxHash string `json:"txHash"`
		TxLink string `json:"txLink"`
		Amount string `json:"amount"`
		Token  struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"token"`
		ChainId  int    `json:"chainId"`
		GasPrice string `json:"gasPrice"`
		GasUsed  string `json:"gasUsed"`
		GasToken struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"gasToken"`
		GasAmount    string `json:"gasAmount"`
		GasAmountUSD string `json:"gasAmountUSD"`
		AmountUSD    string `json:"amountUSD"`
		Value        string `json:"value"`
		Timestamp    int    `json:"timestamp"`
	} `json:"receiving"`
	LifiExplorerLink string `json:"lifiExplorerLink"`
	FromAddress      string `json:"fromAddress"`
	ToAddress        string `json:"toAddress"`
	Tool             string `json:"tool"`
	Status           string `json:"status"`
	Substatus        string `json:"substatus"`
	SubstatusMessage string `json:"substatusMessage"`
	Metadata         struct {
		Integrator string `json:"integrator"`
	} `json:"metadata"`
}

type Quote struct {
	Resp
	Action struct {
		FromToken struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			LogoURI  string `json:"logoURI"`
			PriceUSD string `json:"priceUSD"`
		} `json:"fromToken"`
		FromAmount string `json:"fromAmount"`
		ToToken    struct {
			Address  string `json:"address"`
			ChainId  int    `json:"chainId"`
			Symbol   string `json:"symbol"`
			Decimals int    `json:"decimals"`
			Name     string `json:"name"`
			CoinKey  string `json:"coinKey"`
			PriceUSD string `json:"priceUSD"`
		} `json:"toToken"`
		FromChainId int     `json:"fromChainId"`
		ToChainId   int     `json:"toChainId"`
		Slippage    float64 `json:"slippage"`
		FromAddress string  `json:"fromAddress"`
		ToAddress   string  `json:"toAddress"`
	} `json:"action"`
	Estimate struct {
		ApprovalAddress string          `json:"approvalAddress"`
		ToAmountMin     decimal.Decimal `json:"toAmountMin"`
		ToAmount        decimal.Decimal `json:"toAmount"`
		FromAmount      decimal.Decimal `json:"fromAmount"`
		FeeCosts        []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Token       struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"token"`
			Amount     string `json:"amount"`
			AmountUSD  string `json:"amountUSD"`
			Percentage string `json:"percentage"`
			Included   bool   `json:"included"`
		} `json:"feeCosts"`
		GasCosts []struct {
			Type      string `json:"type"`
			Price     string `json:"price"`
			Estimate  string `json:"estimate"`
			Limit     string `json:"limit"`
			Amount    string `json:"amount"`
			AmountUSD string `json:"amountUSD"`
			Token     struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"token"`
		} `json:"gasCosts"`
		ExecutionDuration float64 `json:"executionDuration"`
		FromAmountUSD     string  `json:"fromAmountUSD"`
		ToAmountUSD       string  `json:"toAmountUSD"`
	} `json:"estimate"`
	IncludedSteps []struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Action struct {
			FromChainId int    `json:"fromChainId"`
			FromAmount  string `json:"fromAmount"`
			FromToken   struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI"`
				PriceUSD string `json:"priceUSD"`
			} `json:"fromToken"`
			ToChainId int `json:"toChainId"`
			ToToken   struct {
				Address  string `json:"address"`
				ChainId  int    `json:"chainId"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				Name     string `json:"name"`
				CoinKey  string `json:"coinKey"`
				LogoURI  string `json:"logoURI,omitempty"`
				PriceUSD string `json:"priceUSD"`
			} `json:"toToken"`
			Slippage                  float64 `json:"slippage"`
			FromAddress               string  `json:"fromAddress"`
			ToAddress                 string  `json:"toAddress"`
			DestinationGasConsumption string  `json:"destinationGasConsumption,omitempty"`
			DestinationCallData       string  `json:"destinationCallData,omitempty"`
		} `json:"action"`
		Estimate struct {
			Tool              string  `json:"tool"`
			FromAmount        string  `json:"fromAmount"`
			ToAmount          string  `json:"toAmount"`
			ToAmountMin       string  `json:"toAmountMin"`
			ApprovalAddress   string  `json:"approvalAddress"`
			ExecutionDuration float64 `json:"executionDuration"`
			GasCosts          []struct {
				Type      string `json:"type"`
				Price     string `json:"price"`
				Estimate  string `json:"estimate"`
				Limit     string `json:"limit"`
				Amount    string `json:"amount"`
				AmountUSD string `json:"amountUSD"`
				Token     struct {
					Address  string `json:"address"`
					ChainId  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"token"`
			} `json:"gasCosts"`
			FeeCosts []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				Token       struct {
					Address  string `json:"address"`
					ChainId  int    `json:"chainId"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					Name     string `json:"name"`
					CoinKey  string `json:"coinKey"`
					LogoURI  string `json:"logoURI"`
					PriceUSD string `json:"priceUSD"`
				} `json:"token"`
				Amount     string `json:"amount"`
				AmountUSD  string `json:"amountUSD"`
				Percentage string `json:"percentage"`
				Included   bool   `json:"included"`
			} `json:"feeCosts,omitempty"`
		} `json:"estimate"`
		Tool        string `json:"tool"`
		ToolDetails struct {
			Key     string `json:"key"`
			Name    string `json:"name"`
			LogoURI string `json:"logoURI"`
		} `json:"toolDetails"`
	} `json:"includedSteps"`
	TransactionRequest struct {
		Data    string         `json:"data"`
		To      common.Address `json:"to"`
		Value   string         `json:"value"`
		From    common.Address `json:"from"`
		ChainId int            `json:"chainId"`
	} `json:"transactionRequest"`
}

type QuoteParams struct {
	FromChainId   int             `json:"fromChainId"`
	ToChainId     int             `json:"toChainId"`
	FromToken     common.Address  `json:"fromToken"`
	ToToken       common.Address  `json:"toToken"`
	FromAmountWei decimal.Decimal `json:"fromAmount"`
	FromAddress   common.Address  `json:"fromAddress"`
	ToAddress     common.Address  `json:"toAddress"`
}

func (r Resp) Error() error {
	if r.Code != 0 {
		return errors.New(r.Message)
	}
	return nil
}

func Action2Int(action string) int {
	switch action {
	case ApproveTransactionAction:
		return 1
	case SourceChainSendingAction:
		return 2
	case WaitForTxAction:
		return 3
	default:
		return 0
	}
}
