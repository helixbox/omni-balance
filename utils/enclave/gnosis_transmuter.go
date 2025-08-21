package enclave

import (
	"math/big"

	gnosis_transmuter "omni-balance/utils/enclave/router/gnosis/transmuter"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type GnosisTransmuterRequest struct {
	GnosisTransmuter GnosisTransmuter `json:"gnosisUSDCTransmuter"`
}

func (a GnosisTransmuterRequest) GetRequestType() string {
	return "gnosis_transmuter"
}

type GnosisTransmuter struct {
	Amount string `json:"amount"`
	Meta   Meta   `json:"meta"`
}

func (c *Client) SignGnosisTransmuter(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	amount, err := getGnosisTransmuterInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	withdraw := GnosisTransmuter{
		Amount: amount.String(),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := GnosisTransmuterRequest{
		GnosisTransmuter: withdraw,
	}

	return c.signRequest(req, tx, chainID)
}

func getGnosisTransmuterInfo(input []byte) (*big.Int, error) {
	if len(input) < 4 {
		return nil, errors.New("data too short")
	}

	routerAbi, err := gnosis_transmuter.GnosisTransmuterMetaData.ParseABI()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["withdraw"].Inputs.Unpack(input[4:])
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(args) != 1 {
		return nil, errors.New("invalid number of args")
	}

	return args[0].(*big.Int), nil
}
