package enclave

import (
	gnosis_claim "omni-balance/utils/enclave/router/gnosis/claim"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type GnosisClaimRequest struct {
	GnosisClaim GnosisClaim `json:"gnosisClaim"`
}

func (a GnosisClaimRequest) GetRequestType() string {
	return "gnosis_claim"
}

type GnosisClaim struct {
	Message    string `json:"message"`
	Signatures string `json:"signatures"`
	Meta       Meta   `json:"meta"`
}

func (c *Client) SignGnosisClaim(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	req, err := buildGnosisClaimRequest(tx.Data(), tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c.signRequest(req, tx, chainID)
}

func buildGnosisClaimRequest(input []byte, tx *types.Transaction) (GnosisClaimRequest, error) {
	if len(input) < 4 {
		return GnosisClaimRequest{}, errors.New("data too short")
	}

	routerAbi, err := gnosis_claim.GnosisClaimMetaData.ParseABI()
	if err != nil {
		return GnosisClaimRequest{}, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["safeExecuteSignaturesWithAutoGasLimit"].Inputs.Unpack(input[4:])
	if err != nil {
		return GnosisClaimRequest{}, errors.Wrap(err, "unpack")
	}

	if len(args) != 2 {
		return GnosisClaimRequest{}, errors.New("invalid number of args")
	}

	claim := GnosisClaim{
		Message:    common.Bytes2Hex(args[0].([]byte)),
		Signatures: common.Bytes2Hex(args[1].([]byte)),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	return GnosisClaimRequest{GnosisClaim: claim}, nil
}
