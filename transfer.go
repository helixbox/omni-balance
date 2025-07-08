package enclave

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type TransferRequest struct {
	Erc20Transfer Erc20Transfer `json:"erc20Transfer"`
}

func (s TransferRequest) GetRequestType() string {
	return "erc20_transfer"
}

type Erc20Transfer struct {
	Token    common.Address `json:"token"`
	Receiver common.Address `json:"receiver"`
	Amount   *big.Int       `json:"amount"`
	Meta     Meta           `json:"meta"`
}

func buildTransferRequest(tx *types.Transaction) (*TransferRequest, error) {
	receiver, amount, err := GetTransferInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	transfer := &Erc20Transfer{
		Token:    *tx.To(),
		Receiver: receiver,
		Amount:   amount,
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	return &TransferRequest{
		Erc20Transfer: *transfer,
	}, nil
}
