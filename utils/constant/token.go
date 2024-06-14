package constant

import "github.com/ethereum/go-ethereum/common"

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	//stableTokens = map[string]StableToken{
	//	Ethereum: {
	//		USDT: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
	//		USDC: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
	//	},
	//	Polygon: {
	//		USDT: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
	//		USDC: common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
	//	},
	//	Bnb: {
	//		USDT: common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
	//		USDC: common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
	//	},
	//	Avalanche: {
	//		USDT: common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
	//		USDC: common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
	//	},
	//	Arbitrum: {
	//		USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0Ps0000"),
	//		USDC: common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
	//	},
	//	Optimism: {
	//		USDT: common.HexToAddress("0x94b008aA00579c1307B06f3CD2E8CEa38539Ef48"),
	//		USDC: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
	//	},
	//	Base: {
	//		USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0Ps0000"),
	//		USDC: common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
	//	},
	//	Celo: {
	//		USDT: common.HexToAddress("0x765DE816845861e75A25D541C82aB31593bB6eD8"),
	//		USDC: common.HexToAddress("0xF194afDf50B03e69Bd7D057c1Aa9e10c9954E4C9"),
	//	},
	//	Blast: {
	//		USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0Ps0000"),
	//		USDC: common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
	//	},
	//}
)

type StableToken struct {
	USDT common.Address
	USDC common.Address
}
