package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"omni-balance/utils/chains"
	"omni-balance/utils/provider/bridge/helix/abi_v2_opposite"
)

type V2Opposite struct {
	opts Options
}

func NewV2Opposite(opts Options) Transfer {
	return &V2Opposite{opts: opts}
}

func (v2 *V2Opposite) getContract(_, _ string) (common.Address, common.Address) {
	return common.HexToAddress("0x48d769d5C7ff75703cDd1543A1a2ed9bC9044A23"), common.HexToAddress("0x48d769d5C7ff75703cDd1543A1a2ed9bC9044A23")
}

func (v2 *V2Opposite) Do(_ context.Context, opts TransferOptions) (tx *types.LegacyTx, err error) {
	var (
		sourceContractAddress, _ = v2.getContract(v2.opts.SourceChain, v2.opts.TargetChain)
		targetChain              = v2.opts.Config.GetChainConfig(v2.opts.TargetChain)
		targetToken              = v2.opts.Config.GetTokenInfoOnChain(v2.opts.TargetTokenName, v2.opts.TargetChain)
		sourceToken              = v2.opts.Config.GetTokenInfoOnChain(v2.opts.SourceTokenName, v2.opts.SourceChain)
		isNativeToken            = v2.opts.Config.IsNativeToken(v2.opts.SourceChain, sourceToken.Name)
		transferId               [32]byte
		_                        = copy(transferId[:], opts.TransferId.Bytes())
		snapshot                 = abi_v2_opposite.LnOppositeBridgeSourceSnapshot{
			RemoteChainId:   big.NewInt(int64(targetChain.Id)),
			Provider:        opts.Relayer,
			SourceToken:     common.HexToAddress(sourceToken.ContractAddress),
			TargetToken:     common.HexToAddress(targetToken.ContractAddress),
			TransferId:      transferId,
			TotalFee:        opts.TotalFee.BigInt(),
			DepositedMargin: opts.DepositedMargin.BigInt(),
		}
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(v2.opts.Amount, sourceToken.Decimals), 0)
	data, err := v2.transferAndLockMargin(snapshot, amount.BigInt(), v2.opts.Recipient)
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

func (v2 *V2Opposite) transferAndLockMargin(_snapshot abi_v2_opposite.LnOppositeBridgeSourceSnapshot,
	_amount *big.Int, _receiver common.Address) ([]byte, error) {

	abiObj, err := abi_v2_opposite.AbiV2OppositeMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "GetAbi")
	}
	return abiObj.Pack("transferAndLockMargin", _snapshot, _amount, _receiver)
}
