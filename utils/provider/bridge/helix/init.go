package helix

import (
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
)

func init() {
	provider.Register(configs.Bridge, New)
}

type BridgeType string

const (
	LnV2DefaultType           BridgeType = "lnv2-default"
	LnV2OppositeType          BridgeType = "lnv2-opposite"
	LnV3Type                  BridgeType = "lnv3"
	sourceChainSendingAction             = "sourceChainSending"
	sourceChainSentAction                = "sourceChainSent"
	targetChainSendingAction             = "targetChainSending"
	targetChainReceivedAction            = "targetChainReceived"
)

var (
	BRIDGES = map[BridgeType]func(opts Options) Transfer{
		LnV2DefaultType:  NewV2Default,
		LnV2OppositeType: NewV2Opposite,
		LnV3Type:         NewV3,
	}
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
