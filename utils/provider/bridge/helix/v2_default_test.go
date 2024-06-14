package helix

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"testing"
)

var conf = &configs.Config{
	Chains: []configs.Chain{
		{
			Name:         "polygon",
			RpcEndpoints: nil,
			NativeToken:  "MATIC",
			Tokens: []configs.Token{
				{
					Name:            "RING",
					ContractAddress: "0x9C1C23E60B72Bc88a043bf64aFdb16A02540Ae8f",
					Decimals:        18,
				},
			},
			Id: 137,
		},

		{
			Name:         "ethereum",
			RpcEndpoints: nil,
			NativeToken:  "ETH",
			Tokens: []configs.Token{
				{
					Name:            "RING",
					ContractAddress: "0x9469d013805bffb7d3debe5e7839237e535ec483",
					Decimals:        18,
				},
			},
			Id: 1,
		},
		{
			Name:         "darwinia",
			RpcEndpoints: nil,
			NativeToken:  "RING",
			Tokens: []configs.Token{
				{
					Name:            "RING",
					ContractAddress: constant.ZeroAddress.Hex(),
					Decimals:        18,
				},
			},
			Id: 46,
		},
		{
			Name:         "arbitrum",
			RpcEndpoints: nil,
			NativeToken:  "ETH",
			Tokens: []configs.Token{
				{
					Name:            "ETH",
					ContractAddress: constant.ZeroAddress.Hex(),
					Decimals:        18,
				},
				{
					Name:            "RING",
					ContractAddress: "0x9e523234D36973f9e38642886197D023C88e307e",
					Decimals:        18,
				},
				{
					Name:            "USDT",
					ContractAddress: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
					Decimals:        6,
				},
			},
			Id: 42161,
		},
		{
			Name:         "op",
			RpcEndpoints: nil,
			Tokens: []configs.Token{
				{
					Name:            "USDT",
					ContractAddress: "0x94b008aA00579c1307B0EF2c499aD98a8ce58e58",
					Decimals:        6,
				},
			},
			Id: 10,
		},
		{
			Name:         "blast",
			RpcEndpoints: nil,
			Tokens: []configs.Token{
				{
					Name:            "ETH",
					ContractAddress: constant.ZeroAddress.Hex(),
					Decimals:        18,
				},
			},
			Id: 81457,
		},
	},
}

func init() {
	conf.Init()
}

func TestV2Default_Do(t *testing.T) {
	args := []struct {
		name          string
		expectedData  string
		sourceChain   string
		targetChain   string
		amount        decimal.Decimal
		tokenName     string
		recipient     common.Address
		sender        common.Address
		relayer       common.Address
		transferId    common.Hash
		withdrawNonce uint64
		totalFee      decimal.Decimal
		to            common.Address
	}{
		{
			name:         "ethereum->RING->darwinia-dvm",
			expectedData: "0be175a5000000000000000000000000000000000000000000000000000000000000002e0000000000000000000000003b9e571adecb0c277486036d6097e9c2cccfa9d90000000000000000000000009469d013805bffb7d3debe5e7839237e535ec4830000000000000000000000000000000000000000000000000000000000000000ce6cf916269ec90c387e8fa1ef25b09296a310ea99a525872f9cc1e3b1f1080500000000000000000000000000000000000000000000001043561a882930000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000178d8546c5f78e01133858958355b06ec3406a1a",
			sourceChain:  "ethereum",
			targetChain:  "darwinia-dvm",
			amount:       decimal.New(0, 0),
			tokenName:    "RING",
			recipient:    common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			sender:       common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			relayer:      common.HexToAddress("0x3b9e571adecb0c277486036d6097e9c2cccfa9d9"),
			transferId:   common.HexToHash("0xce6cf916269ec90c387e8fa1ef25b09296a310ea99a525872f9cc1e3b1f10805"),
			totalFee:     decimal.RequireFromString("300000000000000000000"),
			to:           common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337"),
		},
		{
			name:          "arbitrum->RING->polygon",
			expectedData:  "0be175a500000000000000000000000000000000000000000000000000000000000000890000000000000000000000003b9e571adecb0c277486036d6097e9c2cccfa9d90000000000000000000000009e523234d36973f9e38642886197d023c88e307e0000000000000000000000009c1c23e60b72bc88a043bf64afdb16a02540ae8f71e5800b0122a77638334639f447eb4d85ce9e8256a090cbeb1e7839de15211a00000000000000000000000000000000000000000000000ad78ebc5ac620000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000178d8546c5f78e01133858958355b06ec3406a1a",
			sourceChain:   "arbitrum",
			targetChain:   "polygon",
			amount:        decimal.New(0, 0),
			tokenName:     "RING",
			recipient:     common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			sender:        common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			relayer:       common.HexToAddress("0x3b9e571adecb0c277486036d6097e9c2cccfa9d9"),
			transferId:    common.HexToHash("0x71e5800b0122a77638334639f447eb4d85ce9e8256a090cbeb1e7839de15211a"),
			totalFee:      decimal.RequireFromString("200000000000000000000"),
			withdrawNonce: 1,
			to:            common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337"),
		},
		{
			name:          "arbitrum->RING->polygon",
			expectedData:  "0be175a500000000000000000000000000000000000000000000000000000000000000890000000000000000000000003b9e571adecb0c277486036d6097e9c2cccfa9d90000000000000000000000009e523234d36973f9e38642886197d023c88e307e0000000000000000000000009c1c23e60b72bc88a043bf64afdb16a02540ae8f71e5800b0122a77638334639f447eb4d85ce9e8256a090cbeb1e7839de15211a00000000000000000000000000000000000000000000000ad78ebc5ac620000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000178d8546c5f78e01133858958355b06ec3406a1a",
			sourceChain:   "arbitrum",
			targetChain:   "polygon",
			amount:        decimal.New(0, 0),
			tokenName:     "RING",
			recipient:     common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			sender:        common.HexToAddress("0x178D8546C5f78e01133858958355B06EC3406A1A"),
			relayer:       common.HexToAddress("0x3b9e571adecb0c277486036d6097e9c2cccfa9d9"),
			transferId:    common.HexToHash("0x71e5800b0122a77638334639f447eb4d85ce9e8256a090cbeb1e7839de15211a"),
			totalFee:      decimal.RequireFromString("200000000000000000000"),
			withdrawNonce: 1,
			to:            common.HexToAddress("0x94C614DAeFDbf151E1BB53d6A201ae5fF56A9337"),
		},
	}

	for _, arg := range args {
		t.Run(arg.name, func(t *testing.T) {
			v2 := NewV2Default(Options{
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
				Relayer:       arg.relayer,
				TransferId:    arg.transferId,
				TotalFee:      arg.totalFee,
				WithdrawNonce: arg.withdrawNonce,
			})
			assert.NoError(t, err)
			assert.Equal(t, arg.expectedData, common.Bytes2Hex(tx.Data))
			assert.Equal(t, arg.to.String(), tx.To.String())
		})
	}
}
