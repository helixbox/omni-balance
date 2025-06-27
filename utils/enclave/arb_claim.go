package enclave

import (
	"math/big"

	"omni-balance/utils/enclave/router/arb/outbox"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ArbitrumClaimRequest struct {
	ArbitrumClaim ArbitrumClaim `json:"arbClaim"`
}

func (a ArbitrumClaimRequest) GetRequestType() string {
	return "arbitrum_claim"
}

// bytes32[] proof,uint256 index,address l2Sender,address to,uint256 l2Block,uint256 l1Block,uint256 l2Timestamp,uint256 value,bytes data
type ArbitrumClaim struct {
	Proof       []common.Hash `json:"proof"`
	Index       string        `json:"index"`
	L2Sender    string        `json:"l2Sender"`
	To          string        `json:"to"`
	L2Block     string        `json:"l2Block"`
	L1Block     string        `json:"l1Block"`
	L2Timestamp string        `json:"l2Timestamp"`
	Value       string        `json:"value"`
	Data        string        `json:"data"`
	Meta        Meta          `json:"meta"`
}

func (c *Client) SignArbitrumClaim(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	req, err := BuildClaimRequest(tx.Data(), tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c.signRequest(req, tx, chainID)
}

func BuildClaimRequest(input []byte, tx *types.Transaction) (ArbitrumClaimRequest, error) {
	if len(input) < 4 {
		return ArbitrumClaimRequest{}, errors.New("data too short")
	}

	outboxAbi, err := outbox.OutboxMetaData.GetAbi()
	if err != nil {
		return ArbitrumClaimRequest{}, errors.WithStack(err)
	}

	args, err := outboxAbi.Methods["executeTransaction"].Inputs.Unpack(input[4:])
	if err != nil {
		return ArbitrumClaimRequest{}, errors.Wrap(err, "unpack")
	}

	if len(args) != 9 {
		return ArbitrumClaimRequest{}, errors.New("invalid number of args")
	}

	claim := ArbitrumClaim{
		Proof:       convertToCommonHashSlice(args[0].([][32]uint8)),
		Index:       args[1].(*big.Int).String(),
		L2Sender:    args[2].(common.Address).Hex(),
		To:          args[3].(common.Address).Hex(),
		L2Block:     args[4].(*big.Int).String(),
		L1Block:     args[5].(*big.Int).String(),
		L2Timestamp: args[6].(*big.Int).String(),
		Value:       args[7].(*big.Int).String(),
		Data:        common.Bytes2Hex(args[8].([]byte)),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}
	return ArbitrumClaimRequest{ArbitrumClaim: claim}, nil
}

// convertToCommonHashSlice 将 [][32]uint8 转换为 []common.Hash
func convertToCommonHashSlice(byteSlices [][32]uint8) []common.Hash {
	result := make([]common.Hash, len(byteSlices))
	for i, bytes := range byteSlices {
		result[i] = common.BytesToHash(bytes[:])
	}
	return result
}
