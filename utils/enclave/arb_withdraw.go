package enclave

import (
	"math/big"

	"omni-balance/utils/enclave/router/withdraw"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ArbitrumWithdrawRequest struct {
	ArbitrumWithdraw ArbitrumWithdraw `json:"arbitrumWithdraw"`
}

func (a ArbitrumWithdrawRequest) GetRequestType() string {
	return "arbitrum_withdraw"
}

type ArbitrumWithdraw struct {
	Token  common.Address `json:"l1Token"`
	To     common.Address `json:"to"`
	Amount *big.Int       `json:"amount"`
	Data   []byte         `json:"data"`
	Meta   Meta           `json:"meta"`
}

func (c *Client) SignArbitrumWithdraw(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	token, receiver, amount, data, err := GetWithdrawInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	withdraw := ArbitrumWithdraw{
		Token:  token,
		To:     receiver,
		Amount: amount,
		Data:   data,
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := ArbitrumWithdrawRequest{
		ArbitrumWithdraw: withdraw,
	}

	return c.signRequest(req, tx, chainID)
}

func GetWithdrawInfo(input []byte) (token, recever common.Address, amount *big.Int, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, common.Address{}, nil, nil, errors.New("data too short")
	}

	routerAbi, err := withdraw.WithdrawMetaData.GetAbi()
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["outboundTransfer"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, errors.Wrap(err, "unpack")
	}

	if len(args) != 4 {
		return common.Address{}, common.Address{}, nil, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(common.Address), args[2].(*big.Int), args[5].([]byte), nil
}
