package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"omni-balance/utils/chains"
	"omni-balance/utils/provider/bridge/helix/abi_v3"
	"strings"
	"time"
)

type V3 struct {
	opts Options
}

func NewV3(opts Options) Transfer {
	return &V3{opts: opts}
}

func (v *V3) getContract(sourceChain, targetChain string) (sourceAddress common.Address, targetAddress common.Address) {
	sourceAddress = common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10")
	targetAddress = common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10")
	if strings.EqualFold(sourceChain, "blast") {
		sourceAddress = common.HexToAddress("0xB180D7DcB5CC161C862aD60442FA37527546cAFC")
	}
	if strings.EqualFold(targetChain, "blast") {
		targetAddress = common.HexToAddress("0xB180D7DcB5CC161C862aD60442FA37527546cAFC")
	}

	if strings.EqualFold(sourceChain, "astar-zkevm") {
		sourceAddress = common.HexToAddress("0xD476650e03a45E70202b0bcAfa04E1513920f83a")
	}
	if strings.EqualFold(targetChain, "astar-zkevm") {
		targetAddress = common.HexToAddress("0xD476650e03a45E70202b0bcAfa04E1513920f83a")
	}
	return
}

func (v *V3) Do(_ context.Context, opts TransferOptions) (tx *types.LegacyTx, err error) {
	var (
		sourceContractAddress, _ = v.getContract(v.opts.SourceChain, v.opts.TargetChain)
		targetChain              = v.opts.Config.GetChainConfig(v.opts.TargetChain)
		targetToken              = v.opts.Config.GetTokenInfoOnChain(v.opts.TargetTokenName, v.opts.TargetChain)
		sourceToken              = v.opts.Config.GetTokenInfoOnChain(v.opts.SourceTokenName, v.opts.SourceChain)
		isNativeToken            = v.opts.Config.IsNativeToken(v.opts.SourceChain, v.opts.SourceTokenName)
	)
	if opts.timestamp <= 0 {
		opts.timestamp = time.Now().Unix()
	}
	amount := decimal.NewFromBigInt(chains.EthToWei(v.opts.Amount, sourceToken.Decimals), 0)
	data, err := v.lockAndRemoteRelease(abi_v3.LnBridgeSourceV3TransferParams{
		RemoteChainId: big.NewInt(int64(targetChain.Id)),
		Provider:      opts.Relayer,
		SourceToken:   common.HexToAddress(sourceToken.ContractAddress),
		TargetToken:   common.HexToAddress(targetToken.ContractAddress),
		TotalFee:      opts.TotalFee.BigInt(),
		Amount:        amount.BigInt(),
		Receiver:      v.opts.Recipient,
		Timestamp:     big.NewInt(opts.timestamp),
	})
	if err != nil {
		return nil, errors.Wrap(err, "lockAndRemoteRelease")
	}

	a := &types.LegacyTx{
		To:    &sourceContractAddress,
		Value: amount.Add(opts.TotalFee).BigInt(),
		Data:  data,
	}
	if !isNativeToken {
		a.Value = nil
	}
	return a, nil
}

func (v *V3) lockAndRemoteRelease(args abi_v3.LnBridgeSourceV3TransferParams) ([]byte, error) {
	abiObj, err := abi_v3.AbiV3MetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "GetAbi")
	}
	return abiObj.Pack("lockAndRemoteRelease", args)
}
