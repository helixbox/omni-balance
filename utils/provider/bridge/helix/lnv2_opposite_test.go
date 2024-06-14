package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"omni-balance/utils/constant"
	"testing"
)

func TestV2Opposite_Do(t *testing.T) {
	args := []struct {
		name            string
		expectedData    string
		sourceChain     string
		targetChain     string
		amount          decimal.Decimal
		tokenName       string
		recipient       common.Address
		sender          common.Address
		relayer         common.Address
		transferId      common.Hash
		withdrawNonce   uint64
		DepositedMargin decimal.Decimal
		totalFee        decimal.Decimal
		to              common.Address
	}{
		{
			name:            "arbitrum->RING->ethereum",
			expectedData:    "2656c14700000000000000000000000000000000000000000000000000000000000000010000000000000000000000003b9e571adecb0c277486036d6097e9c2cccfa9d90000000000000000000000009e523234d36973f9e38642886197d023c88e307e0000000000000000000000009469d013805bffb7d3debe5e7839237e535ec48376a404f716bf0df99f9224f551175a2ab1d549b76d70f694b6b999a16ff3676100000000000000000000000000000000000000000000006cf0be29b1954b6400000000000000000000000000000000000000000000003f870857a3e0e38000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000178d8546c5f78e01133858958355b06ec3406a1a",
			sourceChain:     constant.Arbitrum,
			targetChain:     constant.Ethereum,
			amount:          decimal.New(0, 0),
			tokenName:       "RING",
			recipient:       common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			sender:          common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			relayer:         common.HexToAddress("0x3b9e571adecb0c277486036d6097e9c2cccfa9d9"),
			transferId:      common.HexToHash("0x76a404f716bf0df99f9224f551175a2ab1d549b76d70f694b6b999a16ff36761"),
			totalFee:        decimal.RequireFromString("2009595708618000000000"),
			to:              common.HexToAddress("0x48d769d5C7ff75703cDd1543A1a2ed9bC9044A23"),
			DepositedMargin: decimal.RequireFromString("300000000000000000000000"),
		},
		{
			name:            "darwinia-dvm->RING->ethereum",
			expectedData:    "2656c14700000000000000000000000000000000000000000000000000000000000000010000000000000000000000003b9e571adecb0c277486036d6097e9c2cccfa9d900000000000000000000000000000000000000000000000000000000000000000000000000000000000000009469d013805bffb7d3debe5e7839237e535ec4838cb641af960bee6a76f153defd3a6c19908f6f1e5aa68041f030e62b47d2d1b50000000000000000000000000000000000000000000000a3aedf842ccc11da00000000000000000000000000000000000000000000002a5a058fc295ed0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000178d8546c5f78e01133858958355b06ec3406a1a",
			sourceChain:     constant.DarwiniaDvm,
			targetChain:     constant.Ethereum,
			amount:          decimal.New(0, 0),
			tokenName:       "RING",
			recipient:       common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			sender:          common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			relayer:         common.HexToAddress("0x3b9e571adecb0c277486036d6097e9c2cccfa9d9"),
			transferId:      common.HexToHash("0x8cb641af960bee6a76f153defd3a6c19908f6f1e5aa68041f030e62b47d2d1b5"),
			totalFee:        decimal.RequireFromString("3019420219625000000000"),
			to:              common.HexToAddress("0x48d769d5C7ff75703cDd1543A1a2ed9bC9044A23"),
			DepositedMargin: decimal.RequireFromString("200000000000000000000000"),
		},
	}

	for _, arg := range args {
		t.Run(arg.name, func(t *testing.T) {
			v2 := NewV2Opposite(Options{
				SourceTokenName: arg.tokenName,
				TargetTokenName: arg.tokenName,
				SourceChain:     arg.sourceChain,
				TargetChain:     arg.targetChain,
				Config:          *conf,
				Sender:          arg.sender,
				Recipient:       arg.recipient,
				Amount:          arg.amount,
			})
			tx, err := v2.Do(context.Background(), TransferOptions{
				Relayer:         arg.relayer,
				TransferId:      arg.transferId,
				TotalFee:        arg.totalFee,
				WithdrawNonce:   arg.withdrawNonce,
				DepositedMargin: arg.DepositedMargin,
			})
			assert.NoError(t, err)
			assert.Equal(t, arg.expectedData, common.Bytes2Hex(tx.Data))
			assert.Equal(t, arg.to.String(), tx.To.String())
		})
	}
}
