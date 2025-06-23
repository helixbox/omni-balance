package enclave

import (
	"math/big"

	deposit "omni-balance/utils/enclave/router/deposit"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type ArbitrumDepositRequest struct {
	ArbitrumDeposit ArbitrumDeposit `json:"arbDeposit"`
}

func (a ArbitrumDepositRequest) GetRequestType() string {
	return "arbitrum_deposit"
}

type ArbitrumDeposit struct {
	Token       common.Address `json:"token"`
	To          common.Address `json:"to"`
	Amount      *big.Int       `json:"amount"`
	MaxGas      *big.Int       `json:"maxGas"`
	GasPriceBid *big.Int       `json:"gasPriceBid"`
	Data        string         `json:"data"`
	Meta        Meta           `json:"meta"`
}

func (c *Client) SignArbitrumDeposit(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	token, receiver, amount, maxGas, gasPrice, data, err := GetDepositInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	deposit := ArbitrumDeposit{
		Token:       token,
		To:          receiver,
		Amount:      amount,
		MaxGas:      maxGas,
		GasPriceBid: gasPrice,
		Data:        common.Bytes2Hex(data),
		Meta: Meta{
			Nonce:                tx.Nonce(),
			Value:                tx.Value(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	req := ArbitrumDepositRequest{
		ArbitrumDeposit: deposit,
	}

	return c.signRequest(req, tx, chainID)
}

func GetDepositInfo(input []byte) (token, receiver common.Address, amount, maxGas, gasPrice *big.Int, data []byte, err error) {
	if len(input) < 4 {
		return common.Address{}, common.Address{}, nil, nil, nil, nil, errors.New("data too short")
	}

	routerAbi, err := deposit.DepositMetaData.GetAbi()
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, nil, nil, errors.WithStack(err)
	}

	args, err := routerAbi.Methods["outboundTransfer"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, common.Address{}, nil, nil, nil, nil, errors.Wrap(err, "unpack")
	}

	if len(args) != 6 {
		return common.Address{}, common.Address{}, nil, nil, nil, nil, errors.New("invalid number of args")
	}

	return args[0].(common.Address), args[1].(common.Address), args[2].(*big.Int), args[3].(*big.Int), args[4].(*big.Int), args[5].([]byte), nil
}
