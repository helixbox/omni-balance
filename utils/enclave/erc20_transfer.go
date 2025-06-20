package enclave

import (
	"math/big"

	"omni-balance/utils/erc20"

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

func (c *Client) SignErc20Transfer(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	req, err := buildTransferRequest(tx)
	if err != nil {
		return nil, err
	}

	return c.signRequest(req, tx, chainID)
}

// GetTransferInfo input ex: a9059cbb0000000000000000000000000350101f2cb6aa65caab7954246a56f906a3f57d0000000000000000000000000000000000000000033b2e3c9fd0803ce8000000
func GetTransferInfo(input []byte) (to common.Address, amount *big.Int, err error) {
	if len(input) < 4 {
		return common.Address{}, nil, errors.New("invalid input")
	}
	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get abi")
	}
	args, err := erc20Abi.Methods["transfer"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "unpack")
	}
	if len(args) != 2 {
		return common.Address{}, nil, errors.New("invalid args")
	}
	return args[0].(common.Address), args[1].(*big.Int), nil
}
