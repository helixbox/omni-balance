package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
)

type Transfer interface {
	Do(ctx context.Context, opts TransferOptions) (tx *types.LegacyTx, err error)
}

type Options struct {
	SourceTokenName string
	TargetTokenName string
	SourceChain     string
	TargetChain     string
	Config          configs.Config
	Sender          common.Address
	Recipient       common.Address
	Amount          decimal.Decimal
}

type TransferOptions struct {
	Relayer         common.Address
	TransferId      common.Hash
	TotalFee        decimal.Decimal
	WithdrawNonce   uint64
	DepositedMargin decimal.Decimal
	Client          *ethclient.Client
	_bridge         string
	_typename       string
	// only for v3
	timestamp int64
}

type WithdrawFeeArgs struct {
	Amount        decimal.Decimal
	Sender        common.Address
	Relayer       common.Address
	TransferId    common.Hash
	WithdrawNonce string
}

type GetFeeArgs struct {
	BaseFee          decimal.Decimal
	ProtocolFee      decimal.Decimal
	LiquidityFeeRate decimal.Decimal
	TransferAmount   decimal.Decimal
	Sender           common.Address
	Recipient        common.Address
	Relayer          common.Address
}

type HistoryRecordsResult struct {
	Data struct {
		HistoryRecords struct {
			Total    int             `json:"total"`
			Records  []HistoryRecord `json:"records"`
			Typename string          `json:"__typename"`
		} `json:"historyRecords"`
	} `json:"data"`
}
type HistoryRecord struct {
	RequestTxHash   string `json:"requestTxHash"`
	ResponseTxHash  string `json:"responseTxHash"`
	FromChain       string `json:"fromChain"`
	ToChain         string `json:"toChain"`
	StartTime       int    `json:"startTime"`
	SendToken       string `json:"sendToken"`
	SendAmount      string `json:"sendAmount"`
	ConfirmedBlocks string `json:"confirmedBlocks"`
	Result          int    `json:"result"`
	Id              string `json:"id"`
	Typename        string `json:"__typename"`
}

func GetTransferOptions(_ context.Context, amount decimal.Decimal, decimals int32, fromChain, toChain string,
	token common.Address) (TransferOptions, error) {

	result, err := GetSortedLnBridgeRelayInfos(amount, decimals, fromChain, toChain, token)
	if err != nil {
		return TransferOptions{}, errors.Wrap(err, "get sorted ln bridge relay infos")
	}
	info := result.Data.SortedLnBridgeRelayInfos
	if len(info.Records) == 0 {
		return TransferOptions{}, errors.Errorf("no relay info")
	}
	if decimal.RequireFromString(info.TransferLimit).LessThan(amount) {
		return TransferOptions{}, errors.Errorf("transfer amount is too large, must less than %s", info.TransferLimit)
	}
	record := info.Records[0]
	amount = decimal.NewFromBigInt(chains.EthToWei(amount, decimals), 0)
	return TransferOptions{
		_typename:       record.Typename,
		_bridge:         record.Bridge,
		Relayer:         record.Relayer,
		TransferId:      record.LastTransferId,
		TotalFee:        amount.Mul(record.LiquidityFeeRate.Div(decimal.New(100000, 0))).Add(record.ProtocolFee.Add(record.BaseFee)),
		WithdrawNonce:   uint64(record.WithdrawNonce.IntPart()),
		DepositedMargin: record.Margin,
	}, nil
}
