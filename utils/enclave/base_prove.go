package enclave

import (
	"math/big"

	base_portal "omni-balance/utils/enclave/router/base/portal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type BaseProveRequest struct {
	BaseProve BaseProve `json:"baseProve"`
}

func (a BaseProveRequest) GetRequestType() string {
	return "base_claim"
}

// function proveWithdrawalTransaction(
//
//	WithdrawalTransaction memory tx,
//	uint256 disputeGameIndex,
//	OutputRootProof calldata outputRootProof,
//	bytes[] calldata withdrawalProof
//
// );
type BaseProve struct {
	Tx               base_portal.TypesWithdrawalTransaction `json:"tx"`
	DisputeGameIndex string                                 `json:"disputeGameIndex"`
	OutputRootProof  base_portal.TypesOutputRootProof       `json:"outputRootProof"`
	WithdrawalProof  []string                               `json:"withdrawalProof"`
	Meta             Meta                                   `json:"meta"`
}

func (c *Client) SignBaseProve(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	req, err := BuildProveRequest(tx.Data(), tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c.signRequest(req, tx, chainID)
}

func BuildProveRequest(input []byte, tx *types.Transaction) (BaseProveRequest, error) {
	if len(input) < 4 {
		return BaseProveRequest{}, errors.New("data too short")
	}

	portalAbi, err := base_portal.BasePortalMetaData.ParseABI()
	if err != nil {
		return BaseProveRequest{}, errors.WithStack(err)
	}

	args, err := portalAbi.Methods["proveWithdrawalTransaction"].Inputs.Unpack(input[4:])
	if err != nil {
		return BaseProveRequest{}, errors.Wrap(err, "unpack")
	}

	if len(args) != 4 {
		return BaseProveRequest{}, errors.New("invalid number of args")
	}

	prove := BaseProve{
		Tx:               args[0].(base_portal.TypesWithdrawalTransaction),
		DisputeGameIndex: args[1].(*big.Int).String(),
		OutputRootProof:  args[2].(base_portal.TypesOutputRootProof),
		WithdrawalProof:  convertToSlice(args[3].([][]byte)),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}
	return BaseProveRequest{BaseProve: prove}, nil
}

func convertToSlice(byteSlices [][]byte) []string {
	result := make([]string, len(byteSlices))
	for i, bytes := range byteSlices {
		result[i] = common.Bytes2Hex(bytes[:])
	}
	return result
}
