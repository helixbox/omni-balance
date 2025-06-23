package base

import (
	"context"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"
	"strings"

	"github.com/shopspring/decimal"
)

type Ethereum2Base struct {
	config configs.Config
}

func NewL1ToL2(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Ethereum2Base{}, nil
	}
	return &Ethereum2Base{config: conf}, nil
}

func (b *Ethereum2Base) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	_ decimal.Decimal,
) (bool, error) {
	if strings.ToLower(tokenInChainName) == constant.Ethereum && strings.ToLower(tokenOutChainName) == constant.Base {
		if strings.ToUpper(tokenName) == "COW" {
			return true, nil
		}
	}
	return false, nil
}

func (b *Ethereum2Base) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if strings.ToLower(args.TargetChain) == constant.Base {
		return provider.TokenInCosts{
			provider.TokenInCost{
				TokenName:  "ETH",
				CostAmount: decimal.NewFromInt(0),
			},
		}, nil
	}
	return nil, nil
}
