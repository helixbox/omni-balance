package helix_liquidity_claim

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/shopspring/decimal"
	"omni-balance/utils/constant"
)

var (
	messagerAddress = map[string]map[string]common.Address{
		constant.Arbitrum: {
			"msgline":   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			"layerzero": common.HexToAddress("0x509354A4ebf98aCC7a65d2264694A65a2938cac9"),
		},
		constant.Ethereum: {
			"msgline": common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
		},
		constant.Polygon: {
			"msgline":   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			"layerzero": common.HexToAddress("0x463D1730a8527CA58d48EF70C7460B9920346567"),
		},
		constant.DarwiniaDvm: {
			"msgline": common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
		},
		constant.CrabDvm: {
			"msgline": common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
		},
		constant.Mantle: {
			"layerzero": common.HexToAddress("0x61B6B8c7C00aA7F060a2BEDeE6b11927CC9c3eF1"),
		},
		constant.Zksync: {
			"layerzero": common.HexToAddress("0x96892F3EaD26515592Da38432cFABad991BBd69d"),
		},
		constant.Scroll: {
			"layerzero": common.HexToAddress("0x463d1730a8527ca58d48ef70c7460b9920346567"),
		},
		constant.Bsc: {
			"layerzero": common.HexToAddress("0x89AF830781A2C1d3580Db930bea11094F55AfEae"),
		},
		constant.Linea: {
			"layerzero": common.HexToAddress("0x61B6B8c7C00aA7F060a2BEDeE6b11927CC9c3eF1"),
		},
		constant.Optimism: {
			"layerzero": common.HexToAddress("0x61B6B8c7C00aA7F060a2BEDeE6b11927CC9c3eF1"),
		},
		constant.Gnosis: {
			"layerzero": common.HexToAddress("0x3F7DF5866591e7E48D18C8EbeAE61Bc343a63283"),
		},
		constant.Blast: {
			"msgline": common.HexToAddress("0x98982b1685a63596834a05C1288dA7fbF27d684E"),
		},
		constant.AstarZkevm: {
			"layerzero": common.HexToAddress("0x61B6B8c7C00aA7F060a2BEDeE6b11927CC9c3eF1"),
		},
		constant.Moonbeam: {
			"layerzero": common.HexToAddress("0x61B6B8c7C00aA7F060a2BEDeE6b11927CC9c3eF1"),
		},
	}
	lnv3BridgeContractAddress = map[string]common.Address{
		constant.Arbitrum:    common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Ethereum:    common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Polygon:     common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.DarwiniaDvm: common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.CrabDvm:     common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Mantle:      common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Zksync:      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		constant.Scroll:      common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Bsc:         common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Linea:       common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Optimism:    common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Gnosis:      common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		constant.Blast:       common.HexToAddress("0xB180D7DcB5CC161C862aD60442FA37527546cAFC"),
		constant.AstarZkevm:  common.HexToAddress("0xD476650e03a45E70202b0bcAfa04E1513920f83a"),
		constant.Moonbeam:    common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
	}
)

type (
	Records      []Record
	RelayRecords []RelayRecord
)

type MessagerArgs struct {
	ContractAddress        common.Address
	FromChainId, ToChainId int
	RemoteMessager         common.Address
	Payload                []byte
	Refunder               common.Address
	Client                 simulated.Client
}

type EncodePayloadArgs struct {
	FromChainId                       int
	LocalAppAddress, RemoteAppAddress common.Address
	Message                           []byte
}

type Messager interface {
	EncodePayload(ctx context.Context, args EncodePayloadArgs) (payload []byte, err error)
	Params(ctx context.Context, args MessagerArgs) (params MessagerParams, err error)
}

type TxData struct {
	ChainName string
	TxData    []*types.LegacyTx
}

type MessagerParams struct {
	Fee       decimal.Decimal
	ExtParams []byte
}

type ClaimRecords struct {
	Data struct {
		HistoryRecords struct {
			Records []Record `json:"records"`
		} `json:"historyRecords"`
	} `json:"data"`
}

type Record struct {
	Id                  string          `json:"id"`
	Sender              string          `json:"sender"`
	Relayer             string          `json:"relayer"`
	LastRequestWithdraw string          `json:"lastRequestWithdraw"`
	SendAmount          decimal.Decimal `json:"sendAmount"`
	RecvTokenAddress    string          `json:"recvTokenAddress"`
	SendTokenAddress    string          `json:"sendTokenAddress"`
	ToChain             string          `json:"toChain"`
	FromChain           string          `json:"fromChain"`
}

type NeedWithdrawRecords struct {
	ClaimInfo
	TokenName string `json:"token"`
	FromChain string `json:"fromChain"`
	ToChain   string `json:"toChain"`
}

type QueryLnBridgeRelayInfos struct {
	Data struct {
		QueryLnBridgeRelayInfos struct {
			Records RelayRecords `json:"records"`
		} `json:"queryLnBridgeRelayInfos"`
	} `json:"data"`
}

type RelayRecord struct {
	MessageChannel string `json:"messageChannel"`
	FromChain      string `json:"fromChain"`
}

type ClaimInfo struct {
	TransferIds []string
	TotalAmount decimal.Decimal
	Channel     string
}
