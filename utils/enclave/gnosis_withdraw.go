package enclave

import (
	"math/big"

	"omni-balance/utils/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type GnosisWithdrawRequest struct {
	GnosisWithdraw GnosisWithdraw `json:"gnosisWithdraw"`
}

func (a GnosisWithdrawRequest) GetRequestType() string {
	return "gnosis_withdraw"
}

type GnosisWithdraw struct {
	Token common.Address `json:"token"`
	To    common.Address `json:"to"`
	Value string         `json:"value"`
	Data  string         `json:"data"`
	Meta  Meta           `json:"meta"`
}

func (c *Client) SignGnosisWithdraw(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	to, value, data, err := getGnosisWithdrawInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	withdraw := GnosisWithdraw{
		Token: *tx.To(),
		To:    to,
		Value: value.String(),
		Data:  common.Bytes2Hex(data),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := GnosisWithdrawRequest{
		GnosisWithdraw: withdraw,
	}

	return c.signRequest(req, tx, chainID)
}

func getGnosisWithdrawInfo(input []byte) (to common.Address, value *big.Int, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, nil, nil, errors.New("data too short")
	}

	erc1363Abi, err := erc20.Erc1363MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, errors.WithStack(err)
	}

	args, err := erc1363Abi.Methods["transferAndCall0"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, nil, nil, errors.WithStack(err)
	}

	if len(args) != 3 {
		return common.Address{}, nil, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(*big.Int), args[2].([]byte), nil
}
