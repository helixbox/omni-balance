package configs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/jinzhu/copier"
	constant "omni-balance/utils/constant"
)

var (
	// see https://docs.uniswap.org/contracts/v3/reference/deployments/
	WETH_RECIPIENT_ADDRESS        = common.HexToAddress("0x0000000000000000000000000000000000000002")
	UNWRAP_WETH_RECIPIENT_ADDRESS = common.HexToAddress("0x0000000000000000000000000000000000000001")
	ethereumContractAddress       = ContractAddress{
		ChainId:                            1,
		Factory:                            common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
		Multicall:                          common.HexToAddress("0x1F98415757620B543A52E61c46B32eB19261F984"),
		Multicall2:                         common.HexToAddress("0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"),
		ProxyAdmin:                         common.HexToAddress("0xB753548F6E010e7e680BA186F9Ca1BdAB2E90cf2"),
		TickLens:                           common.HexToAddress("0xbfd8137f7d1516D3ea5cA83523914859ec47F573"),
		Quoter:                             common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6"),
		SwapRouter:                         common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
		NFTDescriptor:                      common.HexToAddress("0x42B24A95702b9986e82d421cC3568932790A48Ec"),
		NonfungibleTokenPositionDescriptor: common.HexToAddress("0x91ae842A5Ffd8d12023116943e72A606179294f3"),
		TransparentUpgradeableProxy:        common.HexToAddress("0xEe6A57eC80ea46401049E92587E52f5Ec1c24785"),
		NonfungiblePositionManager:         common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88"),
		V3Migrator:                         common.HexToAddress("0xA5644E29708357803b5A882D272c41cC0dF92B34"),
		QuoterV2:                           common.HexToAddress("0x61fFE014bA17989E743c5F6cB21bF9697530B21e"),
		SwapRouter02:                       common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"),
		Permit2:                            common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
		UniversalRouter:                    common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
		v3StakerAddress:                    common.HexToAddress("0xe34139463bA50bD61336E0c446Bd8C0867c6fE65"),
	}

	WNativeTokenAddress = map[string]common.Address{
		constant.Ethereum:  common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
		constant.Polygon:   common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270"),
		constant.Bnb:       common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"),
		constant.Avalanche: common.HexToAddress("0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7"),
		constant.Arbitrum:  common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1"),
		constant.Optimism:  common.HexToAddress("0x4200000000000000000000000000000000000006"),
		constant.Base:      common.HexToAddress("0x4200000000000000000000000000000000000006"),
		constant.Celo:      common.HexToAddress("0x471EcE3750Da237f93B8E339c536989b8978a438"),
		constant.Blast:     common.HexToAddress("0x4300000000000000000000000000000000000004"),
	}
	stableTokens = map[string]StableToken{
		constant.Ethereum: {
			USDT: common.HexToAddress("0xdAC17F958D2ee523a2206206994597C13D831ec7"),
			USDC: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
		},
		constant.Polygon: {
			USDT: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
			USDC: common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
		},
		constant.Bnb: {
			USDT: common.HexToAddress("0x524bC91Dc82d6b90EF29F76A3ECAaBAffFD490Bc"),
			USDC: common.HexToAddress("0xc1f47175d96fe7c4cd5370552e5954f384e3c791"),
		},
		constant.Avalanche: {
			USDT: common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
			USDC: common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
		},
		constant.Arbitrum: {
			USDT: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
			USDC: common.HexToAddress("0xaf88d065e77c8cc2239327c5edb3a432268e5831"),
		},
		constant.Optimism: {
			USDT: common.HexToAddress("0x94b008aa00579c1307b0ef2c499ad98a8ce58e58"),
			USDC: common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"),
		},
		constant.Base: {
			USDC: common.HexToAddress("0x833589fcd6edb6e08f4c7c32d4f71b54bda02913"),
		},
		constant.Celo: {
			USDC: common.HexToAddress("0xcebA9300f2b948710d2653dD7B07f33A8B32118C"),
		},
	}
)

var (
	V2Pool = "v2-pool"
	V3Pool = "v3-pool"
)

type StableToken struct {
	USDT common.Address
	USDC common.Address
}

type ContractAddress struct {
	ChainId                            int `json:"chain_id"`
	Factory                            common.Address
	Multicall                          common.Address
	Multicall2                         common.Address
	ProxyAdmin                         common.Address
	TickLens                           common.Address
	Quoter                             common.Address
	SwapRouter                         common.Address
	NFTDescriptor                      common.Address
	NonfungibleTokenPositionDescriptor common.Address
	TransparentUpgradeableProxy        common.Address
	NonfungiblePositionManager         common.Address
	V3Migrator                         common.Address
	QuoterV2                           common.Address
	SwapRouter02                       common.Address
	Permit2                            common.Address
	UniversalRouter                    common.Address
	v3StakerAddress                    common.Address
}

func (c ContractAddress) Clone() *ContractAddress {
	var newContractAddress = new(ContractAddress)
	_ = copier.Copy(newContractAddress, &c)
	return newContractAddress
}

func GetContractAddress(chainName string) ContractAddress {
	var result = ethereumContractAddress.Clone()
	result.ChainId = constant.GetChainId(chainName)
	switch chainName {
	case constant.Ethereum:
		return ethereumContractAddress
	case constant.Arbitrum:
		result.Multicall = common.HexToAddress("0xadF885960B47eA2CD9B55E6DAc6B42b7Cb2806dB")
		result.UniversalRouter = common.HexToAddress("0x5E325eDA8064b456f4781070C0738d849c824258")
		result.Multicall2 = common.HexToAddress("")
		return *result
	case constant.Optimism:
		result.Multicall2 = common.HexToAddress("")
		result.V3Migrator = common.HexToAddress("")
		result.UniversalRouter = common.HexToAddress("0xCb1355ff08Ab38bBCE60111F1bb2B784bE25D7e8")
		return *result
	case constant.Polygon:
		result.Multicall2 = common.HexToAddress("")
		result.UniversalRouter = common.HexToAddress("0xec7BE89e9d109e7e3Fec59c222CF297125FEFda2")
		return *result
	case constant.Base:
		return ContractAddress{
			ChainId:                            constant.GetChainId(chainName),
			Factory:                            common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD"),
			Multicall:                          common.HexToAddress("0x091e99cb1C49331a94dD62755D168E941AbD0693"),
			ProxyAdmin:                         common.HexToAddress("0x3334d83e224aF5ef9C2E7DDA7c7C98Efd9621fA9"),
			TickLens:                           common.HexToAddress("0x0CdeE061c75D43c82520eD998C23ac2991c9ac6d"),
			NFTDescriptor:                      common.HexToAddress("0xF9d1077fd35670d4ACbD27af82652a8d84577d9F"),
			NonfungibleTokenPositionDescriptor: common.HexToAddress("0x4f225937EDc33EFD6109c4ceF7b560B2D6401009"),
			TransparentUpgradeableProxy:        common.HexToAddress("0x4615C383F85D0a2BbED973d83ccecf5CB7121463"),
			NonfungiblePositionManager:         common.HexToAddress("0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1"),
			V3Migrator:                         common.HexToAddress("0x23cF10b1ee3AdfCA73B0eF17C07F7577e7ACd2d7"),
			QuoterV2:                           common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a"),
			SwapRouter02:                       common.HexToAddress("0x2626664c2603336E57B271c5C0b26F421741e481"),
			Permit2:                            common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
			UniversalRouter:                    common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD"),
			v3StakerAddress:                    common.HexToAddress("0x42bE4D6527829FeFA1493e1fb9F3676d2425C3C1"),
		}
	case constant.Bnb:
		return ContractAddress{
			ChainId:                            constant.GetChainId(chainName),
			Factory:                            common.HexToAddress("0xdB1d10011AD0Ff90774D0C6Bb92e5C5c8b4461F7"),
			Multicall:                          common.HexToAddress("0x963Df249eD09c358A4819E39d9Cd5736c3087184"),
			ProxyAdmin:                         common.HexToAddress("0xC9A7f5b73E853664044ab31936D0E6583d8b1c79"),
			TickLens:                           common.HexToAddress("0xD9270014D396281579760619CCf4c3af0501A47C"),
			NFTDescriptor:                      common.HexToAddress("0x831d93E55AF23A2977E4DA892d5005f4F2995071"),
			NonfungibleTokenPositionDescriptor: common.HexToAddress("0x0281E98322e4e8E53491D576Ee6A2BFCE644C55C"),
			TransparentUpgradeableProxy:        common.HexToAddress("0xAec98e489AE35F243eB63452f6ad233A6c97eE97"),
			NonfungiblePositionManager:         common.HexToAddress("0x7b8A01B39D58278b5DE7e48c8449c9f4F5170613"),
			V3Migrator:                         common.HexToAddress("0x32681814957e0C13117ddc0c2aba232b5c9e760f"),
			QuoterV2:                           common.HexToAddress("0x78D78E420Da98ad378D7799bE8f4AF69033EB077"),
			SwapRouter02:                       common.HexToAddress("0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2"),
			Permit2:                            common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
			UniversalRouter:                    common.HexToAddress("0x4Dae2f939ACf50408e13d58534Ff8c2776d45265"),
			v3StakerAddress:                    common.HexToAddress("0x49B53C35AF9072fC71767577BF6380a88EE32C71"),
		}
	case constant.Avalanche:
		return ContractAddress{
			ChainId:                            constant.GetChainId(chainName),
			Factory:                            common.HexToAddress("0x740b1c1de25031C31FF4fC9A62f554A55cdC1baD"),
			Multicall:                          common.HexToAddress("0x0139141Cd4Ee88dF3Cdb65881D411bAE271Ef0C2"),
			ProxyAdmin:                         common.HexToAddress("0x9AdA7D7879214073F40183F3410F2b3f088c6381"),
			TickLens:                           common.HexToAddress("0xEB9fFC8bf81b4fFd11fb6A63a6B0f098c6e21950"),
			NFTDescriptor:                      common.HexToAddress("0x27Dd7eE7fE723e83Bf3612a75a034951fe299E99"),
			NonfungibleTokenPositionDescriptor: common.HexToAddress("0xe89B7C295d73FCCe88eF263F86e7310925DaEBAF"),
			TransparentUpgradeableProxy:        common.HexToAddress("0xE1f93a7cB6fFa2dB4F9d5A2FD43158A428993C09"),
			NonfungiblePositionManager:         common.HexToAddress("0x655C406EBFa14EE2006250925e54ec43AD184f8B"),
			V3Migrator:                         common.HexToAddress("0x44f5f1f5E452ea8d29C890E8F6e893fC0f1f0f97"),
			QuoterV2:                           common.HexToAddress("0xbe0F5544EC67e9B3b2D979aaA43f18Fd87E6257F"),
			SwapRouter02:                       common.HexToAddress("0xbb00FF08d01D300023C629E8fFfFcb65A5a578cE"),
			Permit2:                            common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
			UniversalRouter:                    common.HexToAddress("0x4Dae2f939ACf50408e13d58534Ff8c2776d45265"),
			v3StakerAddress:                    common.HexToAddress("0xCA9D0668C600c4dd07ca54Be1615FE5CDFd76Ac3"),
		}
	case constant.Celo:
		return ContractAddress{
			ChainId:                            constant.GetChainId(chainName),
			Factory:                            common.HexToAddress("0xAfE208a311B21f13EF87E33A90049fC17A7acDEc"),
			Multicall2:                         common.HexToAddress("0x633987602DE5C4F337e3DbF265303A1080324204"),
			ProxyAdmin:                         common.HexToAddress("0xc1b262Dd7643D4B7cA9e51631bBd900a564BF49A"),
			TickLens:                           common.HexToAddress("0x5f115D9113F88e0a0Db1b5033D90D4a9690AcD3D"),
			NFTDescriptor:                      common.HexToAddress("0xa9Fd765d85938D278cb0b108DbE4BF7186831186"),
			NonfungibleTokenPositionDescriptor: common.HexToAddress("0x644023b316bB65175C347DE903B60a756F6dd554"),
			TransparentUpgradeableProxy:        common.HexToAddress("0x505B43c452AA4443e0a6B84bb37771494633Fde9"),
			NonfungiblePositionManager:         common.HexToAddress("0x3d79EdAaBC0EaB6F08ED885C05Fc0B014290D95A"),
			V3Migrator:                         common.HexToAddress("0x3cFd4d48EDfDCC53D3f173F596f621064614C582"),
			QuoterV2:                           common.HexToAddress("0x82825d0554fA07f7FC52Ab63c961F330fdEFa8E8"),
			SwapRouter02:                       common.HexToAddress("0x5615CDAb10dc425a742d643d949a7F474C01abc4"),
			Permit2:                            common.HexToAddress("0x000000000022D473030F116dDEE9F6B43aC78BA3"),
			UniversalRouter:                    common.HexToAddress("0x643770E279d5D0733F21d6DC03A8efbABf3255B4"),
			v3StakerAddress:                    common.HexToAddress("0x6586FB35393abF7Ff454977a9b3c912d218791C6"),
		}
	case constant.Blast:
		return ContractAddress{
			ChainId:                            constant.GetChainId(chainName),
			Factory:                            common.HexToAddress("0x792edAdE80af5fC680d96a2eD80A44247D2Cf6Fd"),
			Multicall:                          common.HexToAddress("0xdC7f370de7631cE9e2c2e1DCDA6B3B5744Cf4705"),
			ProxyAdmin:                         common.HexToAddress("0x7C9cAa4ac84C8FAD8Bd504DBF90e791F91f41705"),
			TickLens:                           common.HexToAddress("0x2E95185bCdD928a3e984B7e2D6560Ab1b17d7274"),
			NFTDescriptor:                      common.HexToAddress("0xAa32bD3926097fd04d22b4433e9867417EE79333"),
			NonfungibleTokenPositionDescriptor: common.HexToAddress("0x497089D9450BB58f536c38c1C0d0A37472303508"),
			TransparentUpgradeableProxy:        common.HexToAddress("0xB22Ef02E13B1900EBF10391e57162402c11BfF05"),
			NonfungiblePositionManager:         common.HexToAddress("0xB218e4f7cF0533d4696fDfC419A0023D33345F28"),
			V3Migrator:                         common.HexToAddress("0x15CA7043CD84C5D21Ae76Ba0A1A967d42c40ecE0"),
			QuoterV2:                           common.HexToAddress("0x6Cdcd65e03c1CEc3730AeeCd45bc140D57A25C77"),
			SwapRouter02:                       common.HexToAddress("0x549FEB8c9bd4c12Ad2AB27022dA12492aC452B66"),
			Permit2:                            common.HexToAddress("0x000000000022d473030f116ddee9f6b43ac78ba3"),
			UniversalRouter:                    common.HexToAddress("0x643770E279d5D0733F21d6DC03A8efbABf3255B4"),
			v3StakerAddress:                    common.HexToAddress("0xEcAF7c276f746170642e97De961f2f0361e1aCc8"),
		}

	}
	return ContractAddress{}
}

func GetRouterAddress(chainName string, poolType string) common.Address {
	contracts := map[string]map[string]common.Address{
		constant.Ethereum: {
			V2Pool: common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"),
			V3Pool: common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564"),
		},
	}
	return contracts[chainName][poolType]
}

func GetNativeTokenWrapperAddress(chainName string) common.Address {
	return WNativeTokenAddress[chainName]
}

func GetStableTokensAddress(chainName string) StableToken {
	return stableTokens[chainName]
}
