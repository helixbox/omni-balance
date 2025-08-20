package enclave

import (
	"math/big"

	gnosis_deposit "omni-balance/utils/enclave/router/gnosis/deposit"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type GnosisDepositRequest struct {
	GnosisDeposit GnosisDeposit `json:"gnosisDeposit"`
}

func (a GnosisDepositRequest) GetRequestType() string {
	return "gnosis_deposit"
}

// pub token: Address,
// pub receiver: Address,
// pub value: U256,
// pub data: Bytes,
type GnosisDeposit struct {
	Token    common.Address `json:"token"`
	Receiver common.Address `json:"receiver"`
	Value    string         `json:"value"`
	Data     string         `json:"data"`
	Meta     Meta           `json:"meta"`
}

func (c *Client) SignGnosisDeposit(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	token, receiver, value, data, err := getGnosisDepositInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	deposit := GnosisDeposit{
		Token:    token,
		Receiver: receiver,
		Value:    value.String(),
		Data:     common.Bytes2Hex(data),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := GnosisDepositRequest{
		GnosisDeposit: deposit,
	}

	return c.signRequest(req, tx, chainID)
}

func getGnosisDepositInfo(input []byte) (token, receiver common.Address, value *big.Int, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, common.Address{}, nil, nil, errors.New("data too short")
	}

	routerAbi, err := gnosis_deposit.GnosisDepositMetaData.ParseABI()
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["relayTokensAndCall"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, errors.WithStack(err)
	}

	if len(args) != 4 {
		return common.Address{}, common.Address{}, nil, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(common.Address), args[2].(*big.Int), args[3].([]byte), nil
}
