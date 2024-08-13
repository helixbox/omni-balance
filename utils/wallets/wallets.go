package wallets

import (
	"context"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type MultiSignType string

const (
	MultiSignTypeSafe MultiSignType = "safe"
)

type Wallets interface {
	Name(ctx context.Context) string
	CheckFullAccess(ctx context.Context) error
	GetAddress(isReal ...bool) common.Address
	IsDifferentAddress() bool
	IsSupportEip712() bool
	SignRawMessage(msg []byte) (sig []byte, err error)

	MarshalJSON() ([]byte, error)

	GetExternalBalance(ctx context.Context, tokenAddress common.Address, decimals int32, client simulated.Client) (decimal.Decimal, error)
	GetBalance(ctx context.Context, tokenAddress common.Address, decimals int32, client simulated.Client) (decimal.Decimal, error)
	GetNonce(ctx context.Context, client simulated.Client) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.LegacyTx, client simulated.Client) (common.Hash, error)
	WaitTransaction(ctx context.Context, txHash common.Hash, client simulated.Client) error
	GetRealHash(ctx context.Context, txHash common.Hash, client simulated.Client) (common.Hash, error)
}

type WalletConfig struct {
	PrivateKey    string         `json:"private_key" yaml:"private_key" comment:"If Operator is empty and not multi sign wallet, then it is Address's private key, otherwise it is Operator's private key"`
	Address       common.Address `json:"address" yaml:"address" comment:"Need to monitor address"`
	Operator      Operator       `json:"operator" yaml:"operator" comment:"Used to isolate the monitoring address and the operation address, preventing the leakage of the monitoring address private key. If Operator is empty, it is not enabled. If it is a multi-sign wallet, it must be set"`
	MultiSignType string         `json:"multi_sign_type" yaml:"multi_sign_type" comment:"MultiSign type, ex: safe"`
}

type Operator struct {
	Address       common.Address `json:"address"`
	Operator      common.Address `json:"operator"`
	PrivateKey    string         `json:"private_key"`
	MultiSignType string         `json:"multi_sign_type" yaml:"multi_sign_type" comment:"MultiSign type, ex: safe"`
}

func (o Operator) IsMultiSign() bool {
	return o.MultiSignType == string(MultiSignTypeSafe)
}

func NewWallets(conf WalletConfig) Wallets {
	wallet := Get(conf.MultiSignType)
	if wallet == nil {
		debug.PrintStack()
		logrus.Fatalf("wallet type '%s' not found", conf.MultiSignType)
	}
	return wallet(conf)
}
