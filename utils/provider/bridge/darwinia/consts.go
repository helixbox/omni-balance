package darwinia

import (
	"context"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SwapFunc func(ctx context.Context, args SwapParams) (tx *types.DynamicFeeTx, err error)

var (
	SUPPORT_TOKENS = []supportToken{
		{Config: crab2darwinia, Fn: Crab2Darwinia},
		{Config: darwinia2Crab, Fn: Darwinia2Crab},
		{Config: darwinia2ethereum, Fn: Darwinia2ethereum},
		{Config: ethereum2darwinia, Fn: Ethereum2darwinia},
	}
	SUPPORTS         = make(map[int64]map[int64]map[string]struct{})
	BRIDGE_DIRECTION = make(map[int64]map[int64]SwapFunc)
)

type supportToken struct {
	Config map[string]Params
	Fn     func(ctx context.Context, args SwapParams) (tx *types.DynamicFeeTx, err error)
}

type Params struct {
	TokenAddress     common.Address
	contractAddress  common.Address
	remoteAppAddress common.Address
	localAppAddress  common.Address
	xTokenAddress    common.Address
	recipient        common.Address
	sourceMessager   common.Address
	targetMessager   common.Address
	originalToken    common.Address
	sourceChainId    int64
	targetChainId    int64
	extData          string
}

func init() {
	for _, supportTokens := range SUPPORT_TOKENS {
		for tokenName, params := range supportTokens.Config {
			if _, ok := BRIDGE_DIRECTION[params.sourceChainId]; !ok {
				BRIDGE_DIRECTION[params.sourceChainId] = make(map[int64]SwapFunc)
			}
			if _, ok := BRIDGE_DIRECTION[params.sourceChainId][params.targetChainId]; !ok {
				BRIDGE_DIRECTION[params.sourceChainId][params.targetChainId] = supportTokens.Fn
			}
			if _, ok := SUPPORTS[params.sourceChainId]; !ok {
				SUPPORTS[params.sourceChainId] = make(map[int64]map[string]struct{})
			}

			if _, ok := SUPPORTS[params.sourceChainId][params.targetChainId]; !ok {
				SUPPORTS[params.sourceChainId][params.targetChainId] = make(map[string]struct{}, 0)
			}

			SUPPORTS[params.sourceChainId][params.targetChainId][tokenName] = struct{}{}
		}
	}
}

func GetSourceChains(targetChainId int64, TokenName string) (sourceChainIds []int64) {
	TokenName = strings.ToUpper(TokenName)
	for sourceChainId := range SUPPORTS {
		if _, ok := SUPPORTS[sourceChainId][targetChainId]; !ok {
			continue
		}
		if _, ok := SUPPORTS[sourceChainId][targetChainId][TokenName]; !ok {
			continue
		}
		sourceChainIds = append(sourceChainIds, sourceChainId)
	}
	return sourceChainIds
}
