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
	return "base_prove"
}

type WithdrawalTransaction struct {
	Nonce    string         `json:"nonce"`
	Sender   common.Address `json:"sender"`
	Target   common.Address `json:"target"`
	Value    string         `json:"value"`
	GasLimit string         `json:"gasLimit"`
	Data     string         `json:"data"`
}

type OutputRootProof struct {
	Version                  common.Hash `json:"version"`
	StateRoot                common.Hash `json:"stateRoot"`
	MessagePasserStorageRoot common.Hash `json:"messagePasserStorageRoot"`
	LatestBlockhash          common.Hash `json:"latestBlockhash"`
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
	Tx               WithdrawalTransaction `json:"tx"`
	DisputeGameIndex string                `json:"disputeGameIndex"`
	OutputRootProof  OutputRootProof       `json:"outputRootProof"`
	WithdrawalProof  []string              `json:"withdrawalProof"`
	Meta             Meta                  `json:"meta"`
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

	txw := args[0].(base_portal.TypesWithdrawalTransaction)
	proof := args[2].(base_portal.TypesOutputRootProof)

	prove := BaseProve{
		Tx: WithdrawalTransaction{
			Nonce:    txw.Nonce.String(),
			Sender:   txw.Sender,
			Target:   txw.Target,
			Value:    txw.Value.String(),
			GasLimit: txw.GasLimit.String(),
			Data:     common.Bytes2Hex(txw.Data[:]),
		},
		DisputeGameIndex: args[1].(*big.Int).String(),
		OutputRootProof: OutputRootProof{
			Version:                  common.Hash(proof.Version),
			StateRoot:                common.Hash(proof.StateRoot),
			MessagePasserStorageRoot: common.Hash(proof.MessagePasserStorageRoot),
			LatestBlockhash:          common.Hash(proof.LatestBlockhash),
		},
		WithdrawalProof: convertToSlice(args[3].([][]byte)),
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
