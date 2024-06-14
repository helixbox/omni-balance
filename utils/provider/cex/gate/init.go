package gate

import (
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
)

func init() {
	provider.Register(configs.CEX, New)
}
