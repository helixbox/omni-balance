package safe

import (
	"omni-balance/utils/constant"
	"omni-balance/utils/wallets"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestSafe_GetAddress(t *testing.T) {
	// 创建 Safe 实例和相关变量
	c := wallets.WalletConfig{
		Address: common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
		Operator: wallets.Operator{
			Address: common.HexToAddress(constant.TestWalletAddress),
		}}

	assert.Equal(t, c.Address.Hex(), (&Safe{conf: c}).GetAddress().Hex(), "should return the configured address")

	assert.Equal(t, c.Address.Hex(), (&Safe{conf: c}).GetAddress(true).Hex(), "should return the configured address")

	c.Operator.MultiSignType = "safe"
	assert.Equal(t, c.Address.Hex(), (&Safe{conf: c}).GetAddress().Hex(), "should return the configured address")
	assert.Equal(t, c.Operator.Address.Hex(), (&Safe{conf: c}).GetAddress(true).Hex(), "should return the configured address")

	c.Operator.MultiSignType = ""
	assert.Equal(t, c.Address.Hex(), (&Safe{conf: c}).GetAddress().Hex(), "should return the configured address")
	assert.Equal(t, c.Address.Hex(), (&Safe{conf: c}).GetAddress(true).Hex(), "should return the configured address")
}
