package bot

import (
	"context"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"

	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/shopspring/decimal"
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

// Bot defines the core interface for balance monitoring and task generation bots.
// Implementations should handle specific balance checking logic for different scenarios.
type Bot interface {
	// Check verifies current conditions and generates required tasks.
	// Parameters:
	//   - ctx: Context for cancellation/timeouts
	//   - args: Execution parameters including wallet config and chain information
	// Returns:
	//   - []Task: List of required balance adjustment tasks
	//   - ProcessType: Execution strategy (parallel/sequential) for the tasks
	//   - error: Validation errors or provider communication failures
	Check(ctx context.Context, args Params) ([]Task, ProcessType, error)

	// Balance retrieves the current token balance for monitoring purposes.
	// Parameters:
	//   - ctx: Context for cancellation/timeouts
	//   - args: Parameters containing wallet and token configuration
	// Returns:
	//   - decimal.Decimal: Current token balance
	//   - error: Balance check failures or chain communication errors
	Balance(ctx context.Context, args Params) (decimal.Decimal, error)

	// Name returns the unique identifier for this bot implementation.
	// Returns:
	//   - string: Machine-readable name used in configuration and logging
	Name() string
}

type Config struct {
	Wallet     wallets.Wallets
	TokenName  string          `json:"token_name"`
	TokenPrice decimal.Decimal `json:"token_price"`
	Chain      string          `json:"chains"`
}

type Task struct {
	Wallet            string               `json:"wallet" gorm:"type:varchar(64)"`
	CurrentBalance    decimal.Decimal      `json:"current_balance"`
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
