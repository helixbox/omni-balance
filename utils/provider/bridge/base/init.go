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

	state1 = "state1"
	state2 = "state2"
	state3 = "state3"
	state4 = "state4"
	state5 = "state5"
	state6 = "state6"
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
	case state1:
		return 1
	case state2:
		return 2
	case state3:
		return 3
	case state4:
		return 4
	case state5:
		return 5
	case state6:
		return 6
	default:
		return 0
	}
}

type tokenConfig struct {
	l1Address common.Address
	l2Address common.Address
}
