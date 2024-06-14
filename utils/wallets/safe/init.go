package safe

import (
	"omni-balance/utils/wallets"
)

func init() {
	wallets.Register("safe", func(conf wallets.WalletConfig) wallets.Wallets {
		return NewSafe(conf)
	})
}
