package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"math/big"
)

type PermitSingle struct {
	Details           PermitDetails         `json:"details"`
	Spender           common.Address        `json:"spender"`
	SigDeadline       *big.Int              `json:"sig_deadline"`
	VerifyingContract common.Address        `json:"verifying_contract"`
	ChainId           *math.HexOrDecimal256 `json:"chain_id"`
}

type PermitDetails struct {
	Token      common.Address `json:"token"`
	Amount     *big.Int       `json:"amount"`
	Expiration *big.Int       `json:"expiration"`
	Nonce      *big.Int       `json:"nonce"`
}
