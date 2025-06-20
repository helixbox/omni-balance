package enclave

import (
	"math/big"

	"omni-balance/utils/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ApproveReqeust struct {
	Erc20Approve Erc20Approve `json:"erc20Approve"`
}

func (r ApproveReqeust) GetRequestType() string {
	return "erc20_approve"
}

type Erc20Approve struct {
	Token   common.Address `json:"token"`
	Spender common.Address `json:"spender"`
	Amount  *big.Int       `json:"amount"`
	Meta    Meta           `json:"meta"`
}

func (c *Client) SignErc20Approve(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	spender, amount, err := GetApproveInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	approve := &Erc20Approve{
		Token:   *tx.To(),
		Spender: spender,
		Amount:  amount,
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := ApproveReqeust{
		Erc20Approve: *approve,
	}
	return c.signRequest(req, tx, chainID)
}

func GetApproveInfo(input []byte) (spender common.Address, amount *big.Int, err error) {
	if len(input) < 4 {
		return common.Address{}, nil, errors.New("invalid input")
	}
	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get abi")
	}
	args, err := erc20Abi.Methods["approve"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "unpack")
	}
	if len(args) != 2 {
		return common.Address{}, nil, errors.New("invalid args")
	}
	return args[0].(common.Address), args[1].(*big.Int), nil
}
