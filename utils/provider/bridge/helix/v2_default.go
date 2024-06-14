package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"omni-balance/utils/chains"
	"omni-balance/utils/provider/bridge/helix/abi_v2_default"
	"strings"
)

type V2Default struct {
	opts Options
}

func NewV2Default(opts Options) Transfer {
	return &V2Default{opts: opts}
}

func (v *V2Default) getContract(sourceChain, targetChain string) (
	sourceContractAddress, targetContractAddress common.Address) {

	switch {
	case strings.EqualFold("zksync", sourceChain):
		sourceContractAddress = common.HexToAddress("0x767Bc046c989f5e63683fB530f939DD34b91ceAC")
		targetContractAddress = common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337")
	case strings.EqualFold("zksync-sepolia", sourceChain):
		sourceContractAddress = common.HexToAddress("0xBe23e871318E49C747CB909AC65aCCFAEAac3a37")
		targetContractAddress = common.HexToAddress("0x8429D7Dfd91D6F970ba89fFC005e67D15f1E4739")
	case strings.EqualFold("zksync-sepolia", targetChain):
		sourceContractAddress = common.HexToAddress("0x8429D7Dfd91D6F970ba89fFC005e67D15f1E4739")
		targetContractAddress = common.HexToAddress("0xBe23e871318E49C747CB909AC65aCCFAEAac3a37")
	case strings.EqualFold("zksync", targetChain):
		sourceContractAddress = common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337")
		targetContractAddress = common.HexToAddress("0x767Bc046c989f5e63683fB530f939DD34b91ceAC")
	default:
		sourceContractAddress = common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337")
		targetContractAddress = common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337")
	}
	return
}

func (v *V2Default) TransferAndLockMargin(_snapshot abi_v2_default.LnDefaultBridgeSourceSnapshot,
	_amount *big.Int, _receiver common.Address) ([]byte, error) {

	abiObj, err := abi_v2_default.HelixMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "GetAbi")
	}
	return abiObj.Pack("transferAndLockMargin", _snapshot, _amount, _receiver)
}

func (v *V2Default) Do(_ context.Context, opts TransferOptions) (tx *types.LegacyTx, err error) {
	var (
		sourceContractAddress, _ = v.getContract(v.opts.SourceChain, v.opts.TargetChain)
		targetChain              = v.opts.Config.GetChainConfig(v.opts.TargetChain)
		targetToken              = v.opts.Config.GetTokenInfoOnChain(v.opts.TargetTokenName, v.opts.TargetChain)
		sourceToken              = v.opts.Config.GetTokenInfoOnChain(v.opts.SourceTokenName, v.opts.SourceChain)
		isNativeToken            = v.opts.Config.IsNativeToken(v.opts.SourceChain, sourceToken.Name)
		transferId               [32]byte
		_                        = copy(transferId[:], opts.TransferId.Bytes())
		snapshot                 = abi_v2_default.LnDefaultBridgeSourceSnapshot{
			RemoteChainId: big.NewInt(int64(targetChain.Id)),
			Provider:      opts.Relayer,
			SourceToken:   common.HexToAddress(sourceToken.ContractAddress),
			TargetToken:   common.HexToAddress(targetToken.ContractAddress),
			TransferId:    transferId,
			TotalFee:      opts.TotalFee.BigInt(),
			WithdrawNonce: opts.WithdrawNonce,
		}
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(v.opts.Amount, sourceToken.Decimals), 0)
	data, err := v.TransferAndLockMargin(snapshot, amount.BigInt(), v.opts.Recipient)
	if err != nil {
		return nil, errors.Wrap(err, "transferAndLockMargin")
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
