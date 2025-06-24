package enclave

import (
	"math/big"

	base_deposit "omni-balance/utils/enclave/router/base/deposit"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type BaseDepositRequest struct {
	BaseDeposit BaseDeposit `json:"baseDeposit"`
}

func (a BaseDepositRequest) GetRequestType() string {
	return "base_deposit"
}

// pub l1_token: Address,
// pub l2_token: Address,
// pub to: Address,
// pub amount: U256,
// pub min_gas_limit: u32,
// pub extra_data: Bytes,
// pub meta: TxMeta,
type BaseDeposit struct {
	L1Token     common.Address `json:"l1Token"`
	L2Token     common.Address `json:"l2Token"`
	To          common.Address `json:"to"`
	Amount      string         `json:"amount"`
	MinGasLimit uint32         `json:"minGasLimit"`
	ExtraData   string         `json:"extraData"`
	Meta        Meta           `json:"meta"`
}

func (c *Client) SignBaseDeposit(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	l1Token, l2Token, receiver, amount, minGasLimit, data, err := getDepositInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	deposit := BaseDeposit{
		L1Token:     l1Token,
		L2Token:     l2Token,
		To:          receiver,
		Amount:      amount.String(),
		MinGasLimit: minGasLimit,
		ExtraData:   common.Bytes2Hex(data),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := BaseDepositRequest{
		BaseDeposit: deposit,
	}

	return c.signRequest(req, tx, chainID)
}

func getDepositInfo(input []byte) (l1Token, l2Token, receiver common.Address, amount *big.Int, minGasLimit uint32, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, common.Address{}, common.Address{}, nil, 0, nil, errors.New("data too short")
	}

	routerAbi, err := base_deposit.BaseDepositMetaData.GetAbi()
	if err != nil {
		return common.Address{}, common.Address{}, common.Address{}, nil, 0, nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["depositERC20To"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, common.Address{}, common.Address{}, nil, 0, nil, errors.Wrap(err, "unpack")
	}

	if len(args) != 6 {
		return common.Address{}, common.Address{}, common.Address{}, nil, 0, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(common.Address), args[2].(common.Address), args[3].(*big.Int), args[4].(uint32), args[5].([]byte), nil
}
