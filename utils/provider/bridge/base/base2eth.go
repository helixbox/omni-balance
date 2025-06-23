package base

//
// import (
// 	"context"
// 	"strings"
//
// 	"omni-balance/utils/configs"
// 	"omni-balance/utils/constant"
// 	"omni-balance/utils/provider"
//
// 	"github.com/shopspring/decimal"
// )
//
// type Base2Ethereum struct {
// 	config configs.Config
// }
//
// func NewL2ToL1(conf configs.Config, noInit ...bool) (provider.Provider, error) {
// 	if len(noInit) > 0 && noInit[0] {
// 		return &Base2Ethereum{}, nil
// 	}
// 	return &Base2Ethereum{config: config}, nil
// }
//
// func (b *Arbitrum2Ethereum) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
// 	_ decimal.Decimal,
// ) (bool, error) {
// 	if strings.ToLower(tokenInChainName) == constant.Base && strings.ToLower(tokenOutChainName) == constant.Ethereum {
// 		if strings.ToUpper(tokenName) == "COW" {
// 			return true, nil
// 		}
// 	}
// 	return false, nil
// }
//
// func (b *Arbitrum2Ethereum) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
// 	if strings.ToLower(args.TargetChain) == constant.Ethereum {
// 		return provider.TokenInCosts{
// 			provider.TokenInCost{
// 				TokenName:  "ETH",
// 				CostAmount: decimal.NewFromInt(0),
// 			},
// 		}, nil
// 	}
// 	return nil, nil
// }
