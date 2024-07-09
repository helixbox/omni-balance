package bot

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/shopspring/decimal"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"
)

type ProcessType string

const (
	Parallel ProcessType = "parallel"
	Queue    ProcessType = "queue"
)

func (p ProcessType) String() string {
	return string(p)
}

type Params struct {
	Conf   configs.Config
	Info   Config
	Client simulated.Client
}

type Bot interface {
	Check(ctx context.Context, args Params) ([]Task, ProcessType, error)
	Name() string
}

type Config struct {
	Wallet    wallets.Wallets
	TokenName string `json:"token_name"`
	Chain     string `json:"chains"`
}

type Task struct {
	Wallet            string               `json:"wallet" gorm:"type:varchar(64)"`
	TokenInName       string               `json:"token_in_name"`
	TokenOutName      string               `json:"token_out_name" gorm:"type:varchar(64)"`
	TokenInChainName  string               `json:"source_chain_name"`
	TokenOutChainName string               `json:"target_chain_name"`
	CurrentChainName  string               `json:"current_chain_name" gorm:"type:varchar(64)"`
	Amount            decimal.Decimal      `json:"amount" gorm:"type:decimal(32,16); default:0"`
	Status            provider.TxStatus    `json:"status" gorm:"type:int; default:0;index"`
	ProviderType      configs.ProviderType `json:"provider_type" gorm:"type:varchar(64)"`
	ProviderName      string               `json:"provider_name" gorm:"type:varchar(64)"`
	Order             interface{}          `json:"order"`
	Remark            string               `json:"remark" grom:"type:varchar(32)"`
}
