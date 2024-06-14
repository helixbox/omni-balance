package snowswap

import (
	"context"
	"github.com/shopspring/decimal"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
)

type Snowswap struct {
}

func (s Snowswap) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	//TODO implement me
	panic("implement me")
}

func (s Snowswap) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s Snowswap) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	//TODO implement me
	panic("implement me")
}

func (s Snowswap) Help() []string {
	//TODO implement me
	panic("implement me")
}

func (s Snowswap) Name() string {
	//TODO implement me
	panic("implement me")
}

func (s Snowswap) Type() configs.LiquidityProviderType {
	//TODO implement me
	panic("implement me")
}
