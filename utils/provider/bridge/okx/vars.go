package okx

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils/constant"
)

var (
	ApproveTransactionAction = "ApproveTransaction"
	SourceChainSendingAction = "sourceChainSending"
	WaitForTxAction          = "WaitForTx"
)

var (
	/*
		WAITING (订单处理中)
		FROM_SUCCESS (源链订单成功)
		BRIDGE_PENDING (跨链桥订单执行中)
		SUCCESS (跨链兑换订单成功)
		REFUND (订单退款)
		FAILURE (订单失败)
	*/
	okxWaitForTxStatus = map[string]string{
		"WAITING":        "order processing",
		"FROM_SUCCESS":   "source chain order success",
		"BRIDGE_PENDING": "cross chain bridge order execution",
		"SUCCESS":        "cross chain exchange order success",
		"REFUND":         "order refund",
		"FAILURE":        "order failure",
	}
)

var okxZeroAddress = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

type OksResp struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (o OksResp) ToInstance(dest interface{}) error {
	data, _ := json.Marshal(o.Data)
	return json.Unmarshal(data, dest)
}

func (o OksResp) TxStatus() (Status, error) {
	var s []Status
	if err := o.ToInstance(&s); err != nil {
		return Status{}, errors.Wrap(err, "unmarshal tx status")
	}
	if len(s) == 0 {
		return Status{}, errors.New("empty tx status")
	}
	return s[0], nil
}

func (o OksResp) ApproveTransaction() (ApproveTransaction, error) {
	var result []ApproveTransaction
	if err := o.ToInstance(&result); err != nil {
		return ApproveTransaction{}, errors.Wrap(err, "unmarshal approve transaction")
	}
	if len(result) == 0 {
		return ApproveTransaction{}, errors.New("empty approve transaction")
	}
	return result[0], nil
}

func (o OksResp) BuildTx() (BuildTxData, error) {
	var result []BuildTxData
	if err := o.ToInstance(&result); err != nil {
		return BuildTxData{}, errors.Wrap(err, "unmarshal build tx")
	}
	if len(result) == 0 {
		return BuildTxData{}, errors.New("empty build tx")
	}
	return result[0], nil
}

func (o OksResp) Quote() (QuoteData, error) {
	var result []QuoteData
	if err := o.ToInstance(&result); err != nil {
		return QuoteData{}, errors.Wrap(err, "unmarshal quote")
	}
	if len(result) == 0 {
		return QuoteData{}, errors.New("empty quote")
	}
	return result[0], nil
}

type Status struct {
	BridgeHash  string `json:"bridgeHash"`
	FromChainId int    `json:"fromChainId"`
	ToChainId   int    `json:"toChainId"`
	//ToAmount               decimal.Decimal `json:"toAmount"`
	ErrorMsg               string `json:"errorMsg"`
	ToTxHash               string `json:"toTxHash"`
	FromTxHash             string `json:"fromTxHash"`
	SourceChainGasfee      string `json:"sourceChainGasfee"`
	DestinationChainGasfee string `json:"destinationChainGasfee"`
	CrossChainFee          struct {
		Symbol  string `json:"symbol"`
		Address string `json:"address"`
		Amount  string `json:"amount"`
	} `json:"crossChainFee"`
	DetailStatus string `json:"detailStatus"`
	Status       string `json:"status"`
}

type ApproveTransaction struct {
	Data               string `json:"data"`
	DexContractAddress string `json:"dexContractAddress"`
	GasLimit           string `json:"gasLimit"`
	GasPrice           string `json:"gasPrice"`
}

type BuildTxData struct {
	FromTokenAmount string `json:"fromTokenAmount"`
	Router          struct {
		BridgeId                  int    `json:"bridgeId"`
		BridgeName                string `json:"bridgeName"`
		CrossChainFee             string `json:"crossChainFee"`
		OtherNativeFee            string `json:"otherNativeFee"`
		CrossChainFeeTokenAddress string `json:"crossChainFeeTokenAddress"`
	} `json:"router"`
	ToTokenAmount decimal.Decimal `json:"toTokenAmount"`
	Tx            struct {
		Data                 string          `json:"data"`
		From                 common.Address  `json:"from"`
		To                   common.Address  `json:"to"`
		Value                decimal.Decimal `json:"value"`
		GasLimit             decimal.Decimal `json:"gasLimit"`
		GasPrice             decimal.Decimal `json:"gasPrice"`
		MaxPriorityFeePerGas string          `json:"maxPriorityFeePerGas"`
	} `json:"tx"`
}

type QuoteData struct {
	FromChainId int `json:"fromChainId"`
	FromToken   struct {
		Decimals             int    `json:"decimals"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
	} `json:"fromToken"`
	FromTokenAmount string `json:"fromTokenAmount"`
	ToChainId       int    `json:"toChainId"`
	ToToken         struct {
		Decimals             int    `json:"decimals"`
		TokenContractAddress string `json:"tokenContractAddress"`
		TokenSymbol          string `json:"tokenSymbol"`
	} `json:"toToken"`
	RouterList []struct {
		EstimateTime      string `json:"estimateTime"`
		FromDexRouterList []struct {
			Router        string `json:"router"`
			RouterPercent string `json:"routerPercent"`
			SubRouterList []struct {
				DexProtocol []struct {
					DexName string `json:"dexName"`
					Percent string `json:"percent"`
				} `json:"dexProtocol"`
				FromToken struct {
					Decimals             int    `json:"decimals"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
				} `json:"fromToken"`
				ToToken struct {
					Decimals             int    `json:"decimals"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
				} `json:"toToken"`
			} `json:"subRouterList"`
		} `json:"fromDexRouterList"`
		MinimumReceived decimal.Decimal `json:"minimumReceived"`
		NeedApprove     int             `json:"needApprove"`
		Router          struct {
			BridgeId                  int    `json:"bridgeId"`
			BridgeName                string `json:"bridgeName"`
			CrossChainFee             string `json:"crossChainFee"`
			OtherNativeFee            string `json:"otherNativeFee"`
			CrossChainFeeTokenAddress string `json:"crossChainFeeTokenAddress"`
		} `json:"router"`
		ToDexRouterList []struct {
			Router        string `json:"router"`
			RouterPercent string `json:"routerPercent"`
			SubRouterList []struct {
				DexProtocol []struct {
					DexName string `json:"dexName"`
					Percent string `json:"percent"`
				} `json:"dexProtocol"`
				FromToken struct {
					Decimals             int    `json:"decimals"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
				} `json:"fromToken"`
				ToToken struct {
					Decimals             int    `json:"decimals"`
					TokenContractAddress string `json:"tokenContractAddress"`
					TokenSymbol          string `json:"tokenSymbol"`
				} `json:"toToken"`
			} `json:"subRouterList"`
		} `json:"toDexRouterList"`
		ToTokenAmount string `json:"toTokenAmount"`
	} `json:"routerList"`
}

type SupportedChain struct {
	OksResp
	Data []SupportedChainData `json:"data"`
}

type SupportedChainData struct {
	ChainId                int    `json:"chainId"`
	ChainName              string `json:"chainName"`
	DexTokenApproveAddress string `json:"dexTokenApproveAddress"`
}

type Config struct {
	ApiKey     string `json:"api_key" yaml:"api_key" comment:"api key"`
	SecretKey  string `json:"secret_key" yaml:"secret_key"`
	Passphrase string `json:"passphrase" yaml:"passphrase"`
	Project    string `json:"project" yaml:"project"`
}

func (r OksResp) Error() error {
	if r.Code != "0" {
		return errors.Errorf(r.Msg)
	}
	return nil
}

func (s Status) ToMap() map[string]interface{} {
	var result = make(map[string]interface{})
	data, _ := json.Marshal(s)
	_ = json.Unmarshal(data, &result)
	return result
}

func (b BuildTxData) ToMap() map[string]interface{} {
	var result = make(map[string]interface{})
	data, _ := json.Marshal(b)
	_ = json.Unmarshal(data, &result)
	return result
}

func (s SupportedChain) Standardize() SupportedChain {
	var result = &SupportedChain{
		OksResp: s.OksResp,
	}

	for _, v := range s.Data {
		chainName := constant.GetChainName(v.ChainId)
		if chainName == "" {
			continue
		}
		result.Data = append(result.Data, SupportedChainData{
			ChainId:                v.ChainId,
			ChainName:              chainName,
			DexTokenApproveAddress: v.DexTokenApproveAddress,
		})
	}
	return *result
}

// StandardizeZeroAddress standardizes the zero address.
// If the provided zeroAddress is equal to constant.ZeroAddress, it returns okxZeroAddress; otherwise, it returns the original address.
// Parameter:
// zeroAddress common.Address - The address to be standardized.
// Return:
// common.Address - The standardized okxZeroAddress
func StandardizeZeroAddress(address common.Address) common.Address {
	if address.Cmp(constant.ZeroAddress) == 0 {
		return okxZeroAddress
	}
	return address
}

// StandardizeZeroAddress2Evm standardizes the zero address to the okxZeroAddress.
func StandardizeZeroAddress2Evm(address common.Address) common.Address {
	if address.Cmp(okxZeroAddress) == 0 {
		return constant.ZeroAddress
	}
	return address
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
