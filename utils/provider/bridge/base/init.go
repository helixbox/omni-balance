package base

import (
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"

	"github.com/ethereum/go-ethereum/common"
)

func init() {
	provider.Register(configs.Bridge, NewL1ToL2)
	provider.Register(configs.Bridge, NewL2ToL1)
}

const (
	sourceChainSendingAction  = "sourceChainSending"
	sourceChainSentAction     = "sourceChainSent"
	targetChainSendingAction  = "targetChainSending"
	targetChainReceivedAction = "targetChainReceived"
)

func Action2Int(action string) int {
	switch action {
	case sourceChainSendingAction:
		return 1
	case sourceChainSentAction:
		return 2
	case targetChainSendingAction:
		return 3
	case targetChainReceivedAction:
		return 4
	default:
		return 0
	}
}

type tokenConfig struct {
	l1Address common.Address
	l2Address common.Address
}
