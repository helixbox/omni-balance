package helix_liquidity_claim

import (
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
)

func init() {
	provider.Register(configs.Bridge, New)
}

func New(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	return Claim{conf: conf}, nil
}
