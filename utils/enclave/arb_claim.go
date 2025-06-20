package enclave

import (
	"math/big"

	"omni-balance/utils/enclave/router/outbox"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ArbitrumClaimRequest struct {
	ArbitrumClaim ArbitrumClaim `json:"arbitrumClaim"`
}

func (a ArbitrumClaimRequest) GetRequestType() string {
	return "arbitrum_claim"
}

// bytes32[] proof,uint256 index,address l2Sender,address to,uint256 l2Block,uint256 l1Block,uint256 l2Timestamp,uint256 value,bytes data
type ArbitrumClaim struct {
	Proof       []common.Hash `json:"proof"`
	Index       *big.Int      `json:"index"`
	L2Block     *big.Int      `json:"l2Block"`
	L1Block     *big.Int      `json:"l1Block"`
	L2Timestamp *big.Int      `json:"l2Timestamp"`
	Value       *big.Int      `json:"value"`
	Data        []byte        `json:"data"`
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
		Proof:       args[0].([]common.Hash),
		Index:       args[1].(*big.Int),
		L2Block:     args[4].(*big.Int),
		L1Block:     args[5].(*big.Int),
		L2Timestamp: args[6].(*big.Int),
		Value:       args[7].(*big.Int),
		Data:        args[8].([]byte),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}
	return ArbitrumClaimRequest{ArbitrumClaim: claim}, nil
}
