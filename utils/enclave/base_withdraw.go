package enclave

import (
	"encoding/hex"
	"math/big"

	base_withdraw "omni-balance/utils/enclave/router/base/withdraw"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type BaseWithdrawRequest struct {
	BaseWithdraw BaseWithdraw `json:"baseWithdraw"`
}

func (a BaseWithdrawRequest) GetRequestType() string {
	return "base_withdraw"
}

type BaseWithdraw struct {
	L2Token     common.Address `json:"l2Token"`
	To          common.Address `json:"to"`
	Amount      string         `json:"amount"`
	MinGasLimit uint32         `json:"minGasLimit"`
	ExtraData   string         `json:"extraData"`
	Meta        Meta           `json:"meta"`
}

func (c *Client) SignBaseWithdraw(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	l2token, receiver, amount, minGasLimit, data, err := getWithdrawInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	withdraw := BaseWithdraw{
		L2Token:     l2token,
		To:          receiver,
		Amount:      amount.String(),
		MinGasLimit: minGasLimit,
		ExtraData:   hex.EncodeToString(data),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := BaseWithdrawRequest{
		BaseWithdraw: withdraw,
	}

	return c.signRequest(req, tx, chainID)
}

func getWithdrawInfo(input []byte) (token, recever common.Address, amount *big.Int, minGasLimit uint32, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, common.Address{}, nil, 0, nil, errors.New("data too short")
	}

	routerAbi, err := base_withdraw.BaseWithdrawMetaData.GetAbi()
	if err != nil {
		return common.Address{}, common.Address{}, nil, 0, nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["withdrawTo"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, common.Address{}, nil, 0, nil, errors.Wrap(err, "unpack")
	}

	if len(args) != 5 {
		return common.Address{}, common.Address{}, nil, 0, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(common.Address), args[2].(*big.Int), args[3].(uint32), args[4].([]byte), nil
}
