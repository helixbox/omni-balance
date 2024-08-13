package multisig

import "omni-balance/utils/wallets"

func init() {
	wallets.Register("safe", NewMultisig)
}
