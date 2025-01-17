package helix_liquidity

import (
	"context"
	"omni-balance/utils/bot/helix_liquidity/lendingpool"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	aaveAddressBook = map[string]AvaeConfig{
		constant.Arbitrum: {
			Chain:     constant.Arbitrum,
			L2Pool:    common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
			Oracle:    common.HexToAddress("0xb56c2F0B653B2e0b10C9b928C8580Ac5Df02C7C7"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{
				"USDT": {
					Name:            "USDT",
					AToken:          common.HexToAddress("0x6ab707Aca953eDAeFBc4fD23bA73294241490620"),
					VToken:          common.HexToAddress("0xfb00AC187a8Eb5AFAE4eACE434F493Eb62672df7"),
					UnderlyingToken: common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"),
					Decimals:        6,
				},

				"USDC.E": {
					Name:            "USDC.E",
					AToken:          common.HexToAddress("0x625E7708f30cA75bfd92586e17077590C60eb4cD"),
					VToken:          common.HexToAddress("0xFCCf3cAbbe80101232d343252614b6A3eE81C989"),
					UnderlyingToken: common.HexToAddress("0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8"),
					Decimals:        6,
				},

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x724dc807b04555b71ed48a6896b6F41593b8C637"),
					VToken:          common.HexToAddress("0xf611aEb5013fD2c0511c9CD55c7dc5C1140741A6"),
					UnderlyingToken: common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831"),
					Decimals:        6,
				},

				"ETH": {
					Name:            "ETH",
					AToken:          common.HexToAddress("0xe50fA9b3c56FfB159cB0FCA61F5c9D750e8128c8"),
					VToken:          common.HexToAddress("0x0c84331e39d6658Cd6e6b9ba04736cC4c4734351"),
					UnderlyingToken: common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"),
					Decimals:        18,
				},

				"WBTC": {
					Name:            "WBTC",
					AToken:          common.HexToAddress("0x078f358208685046a11C85e8ad32895DED33A249"),
					VToken:          common.HexToAddress("0x92b42c66840C7AD907b4BF74879FF3eF7c529473"),
					UnderlyingToken: common.HexToAddress("0x2f2a2543B76A4166549F7aaB2e75Bef0aefC5B0f"),
					Decimals:        8,
				},

				"DAI": {
					Name:            "DAI",
					AToken:          common.HexToAddress("0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE"),
					VToken:          common.HexToAddress("0x8619d80FB0141ba7F184CbF22fd724116D9f7ffC"),
					UnderlyingToken: common.HexToAddress("0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"),
					Decimals:        18,
				},

				"LINK": {
					Name:            "LINK",
					AToken:          common.HexToAddress("0x191c10Aa4AF7C30e871E70C95dB0E4eb77237530"),
					VToken:          common.HexToAddress("0x953A573793604aF8d41F306FEb8274190dB4aE0e"),
					UnderlyingToken: common.HexToAddress("0xf97f4df75117a78c1A5a0DBb814Af92458539FB4"),
					Decimals:        18,
				},
			},
		},
		constant.Optimism: {
			Chain:     constant.Optimism,
			L2Pool:    common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
			Oracle:    common.HexToAddress("0xD81eb3728a631871a7eBBaD631b5f424909f0c77"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDT": {
					Name:            "USDT",
					AToken:          common.HexToAddress("0x6ab707Aca953eDAeFBc4fD23bA73294241490620"),
					VToken:          common.HexToAddress("0xfb00AC187a8Eb5AFAE4eACE434F493Eb62672df7"),
					UnderlyingToken: common.HexToAddress("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58"),
					Decimals:        6,
				},

				"USDC.E": {
					Name:            "USDC.E",
					AToken:          common.HexToAddress("0x625E7708f30cA75bfd92586e17077590C60eb4cD"),
					VToken:          common.HexToAddress("0xFCCf3cAbbe80101232d343252614b6A3eE81C989"),
					UnderlyingToken: common.HexToAddress("0x7F5c764cBc14f9669B88837ca1490cCa17c31607"),
					Decimals:        6,
				},

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x38d693cE1dF5AaDF7bC62595A37D667aD57922e5"),
					VToken:          common.HexToAddress("0x5D557B07776D12967914379C71a1310e917C7555"),
					UnderlyingToken: common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
					Decimals:        6,
				},

				"ETH": {
					Name:            "ETH",
					AToken:          common.HexToAddress("0xe50fA9b3c56FfB159cB0FCA61F5c9D750e8128c8"),
					VToken:          common.HexToAddress("0x0c84331e39d6658Cd6e6b9ba04736cC4c4734351"),
					UnderlyingToken: common.HexToAddress("0x4200000000000000000000000000000000000006"),
					Decimals:        18,
				},

				"WBTC": {
					Name:            "WBTC",
					AToken:          common.HexToAddress("0x078f358208685046a11C85e8ad32895DED33A249"),
					VToken:          common.HexToAddress("0x92b42c66840C7AD907b4BF74879FF3eF7c529473"),
					UnderlyingToken: common.HexToAddress("0x68f180fcCe6836688e9084f035309E29Bf0A2095"),
					Decimals:        8,
				},

				"DAI": {
					Name:            "DAI",
					AToken:          common.HexToAddress("0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE"),
					VToken:          common.HexToAddress("0x8619d80FB0141ba7F184CbF22fd724116D9f7ffC"),
					UnderlyingToken: common.HexToAddress("0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1"),
					Decimals:        18,
				},

				"LINK": {
					Name:            "LINK",
					AToken:          common.HexToAddress("0x191c10Aa4AF7C30e871E70C95dB0E4eb77237530"),
					VToken:          common.HexToAddress("0x953A573793604aF8d41F306FEb8274190dB4aE0e"),
					UnderlyingToken: common.HexToAddress("0x350a791Bfc2C21F9Ed5d10980Dad2e2638ffa7f6"),
					Decimals:        18,
				},
			},
		},
		constant.Scroll: {
			Chain:     constant.Scroll,
			L2Pool:    common.HexToAddress("0x11fCfe756c05AD438e312a7fd934381537D3cFfe"),
			Oracle:    common.HexToAddress("0x04421D8C506E2fA2371a08EfAaBf791F624054F3"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x1D738a3436A8C49CefFbaB7fbF04B660fb528CbD"),
					VToken:          common.HexToAddress("0x3d2E209af5BFa79297C88D6b57F89d792F6E28EE"),
					UnderlyingToken: common.HexToAddress("0x06eFdBFf2a14a7c8E15944D1F4A48F9F95F663A4"),
					Decimals:        6,
				},

				"ETH": {
					Name:            "ETH",
					AToken:          common.HexToAddress("0xf301805bE1Df81102C957f6d4Ce29d2B8c056B2a"),
					VToken:          common.HexToAddress("0xfD7344CeB1Df9Cf238EcD667f4A6F99c6Ef44a56"),
					UnderlyingToken: common.HexToAddress("0x5300000000000000000000000000000000000004"),
					Decimals:        18,
				},
			},
		},
		constant.Base: {
			Chain:     constant.Base,
			L2Pool:    common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5"),
			Oracle:    common.HexToAddress("0x2Cc0Fc26eD4563A5ce5e8bdcfe1A2878676Ae156"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x4e65fE4DbA92790696d040ac24Aa414708F5c0AB"),
					VToken:          common.HexToAddress("0x59dca05b6c26dbd64b5381374aAaC5CD05644C28"),
					UnderlyingToken: common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913"),
					Decimals:        6,
				},

				"ETH": {
					Name:            "ETH",
					AToken:          common.HexToAddress("0xD4a0e0b9149BCee3C920d2E00b5dE09138fd8bb7"),
					VToken:          common.HexToAddress("0x24e6e0795b3c7c71D965fCc4f371803d1c1DcA1E"),
					UnderlyingToken: common.HexToAddress("0x4200000000000000000000000000000000000006"),
					Decimals:        18,
				},
			},
		},
		constant.Polygon: {
			Chain:     constant.Polygon,
			L2Pool:    common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
			Oracle:    common.HexToAddress("0xb023e699F5a33916Ea823A16485e259257cA8Bd1"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC.E": {
					Name:            "USDC.E",
					AToken:          common.HexToAddress("0x625E7708f30cA75bfd92586e17077590C60eb4cD"),
					VToken:          common.HexToAddress("0xFCCf3cAbbe80101232d343252614b6A3eE81C989"),
					UnderlyingToken: common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
					Decimals:        6,
				},

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0xA4D94019934D8333Ef880ABFFbF2FDd611C762BD"),
					VToken:          common.HexToAddress("0xE701126012EC0290822eEA17B794454d1AF8b030"),
					UnderlyingToken: common.HexToAddress("0x3c499c542cEF5E3811e1192ce70d8cC03d5c3359"),
					Decimals:        6,
				},

				"USDT": {
					Name:            "USDT",
					AToken:          common.HexToAddress("0x6ab707Aca953eDAeFBc4fD23bA73294241490620"),
					VToken:          common.HexToAddress("0xfb00AC187a8Eb5AFAE4eACE434F493Eb62672df7"),
					UnderlyingToken: common.HexToAddress("0xc2132D05D31c914a87C6611C10748AEb04B58e8F"),
					Decimals:        6,
				},

				"WBTC": {
					Name:            "WBTC",
					AToken:          common.HexToAddress("0x078f358208685046a11C85e8ad32895DED33A249"),
					VToken:          common.HexToAddress("0x92b42c66840C7AD907b4BF74879FF3eF7c529473"),
					UnderlyingToken: common.HexToAddress("0x1BFD67037B42Cf73acF2047067bd4F2C47D9BfD6"),
					Decimals:        8,
				},

				"DAI": {
					Name:            "DAI",
					AToken:          common.HexToAddress("0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE"),
					VToken:          common.HexToAddress("0x8619d80FB0141ba7F184CbF22fd724116D9f7ffC"),
					UnderlyingToken: common.HexToAddress("0x8f3Cf7ad23Cd3CaDbD9735AFf958023239c6A063"),
					Decimals:        18,
				},

				"LINK": {
					Name:            "LINK",
					AToken:          common.HexToAddress("0x191c10Aa4AF7C30e871E70C95dB0E4eb77237530"),
					VToken:          common.HexToAddress("0x953A573793604aF8d41F306FEb8274190dB4aE0e"),
					UnderlyingToken: common.HexToAddress("0x53E0bca35eC356BD5ddDFebbD1Fc0fD03FaBad39"),
					Decimals:        18,
				},
			},
		},
		constant.Bsc: {
			Chain:     constant.Bsc,
			L2Pool:    common.HexToAddress("0x6807dc923806fE8Fd134338EABCA509979a7e0cB"),
			Oracle:    common.HexToAddress("0x39bc1bfDa2130d6Bb6DBEfd366939b4c7aa7C697"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x00901a076785e0906d1028c7d6372d247bec7d61"),
					VToken:          common.HexToAddress("0xcDBBEd5606d9c5C98eEedd67933991dC17F0c68d"),
					UnderlyingToken: common.HexToAddress("0x8AC76a51cc950d9822D68b83fE1Ad97B32Cd580d"),
					Decimals:        18,
				},

				"USDT": {
					Name:            "USDT",
					AToken:          common.HexToAddress("0xa9251ca9DE909CB71783723713B21E4233fbf1B1"),
					VToken:          common.HexToAddress("0xF8bb2Be50647447Fb355e3a77b81be4db64107cd"),
					UnderlyingToken: common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"),
					Decimals:        18,
				},

				"WBTC": {
					Name:            "WBTC",
					AToken:          common.HexToAddress("0x56a7ddc4e848EbF43845854205ad71D5D5F72d3D"),
					VToken:          common.HexToAddress("0x7b1E82F4f542fbB25D64c5523Fe3e44aBe4F2702"),
					UnderlyingToken: common.HexToAddress("0x7130d2A12B9BCbFAe4f2634d864A1Ee1Ce3Ead9c"),
					Decimals:        18,
				},
			},
		},
		constant.Avalanche: {
			Chain:     constant.Avalanche,
			L2Pool:    common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD"),
			Oracle:    common.HexToAddress("0xEBd36016B3eD09D4693Ed4251c67Bd858c3c7C9C"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x625E7708f30cA75bfd92586e17077590C60eb4cD"),
					VToken:          common.HexToAddress("0xFCCf3cAbbe80101232d343252614b6A3eE81C989"),
					UnderlyingToken: common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E"),
					Decimals:        6,
				},

				"USDT": {
					Name:            "USDT",
					AToken:          common.HexToAddress("0x6ab707Aca953eDAeFBc4fD23bA73294241490620"),
					VToken:          common.HexToAddress("0xfb00AC187a8Eb5AFAE4eACE434F493Eb62672df7"),
					UnderlyingToken: common.HexToAddress("0x9702230A8Ea53601f5cD2dc00fDBc13d4dF4A8c7"),
					Decimals:        6,
				},

				"WBTC": {
					Name:            "WBTC",
					AToken:          common.HexToAddress("0x8ffDf2DE812095b1D19CB146E4c004587C0A0692"),
					VToken:          common.HexToAddress("0xA8669021776Bc142DfcA87c21b4A52595bCbB40a"),
					UnderlyingToken: common.HexToAddress("0x152b9d0FdC40C096757F570A51E494bd4b943E50"),
					Decimals:        8,
				},

				"DAI": {
					Name:            "DAI",
					AToken:          common.HexToAddress("0x82E64f49Ed5EC1bC6e43DAD4FC8Af9bb3A2312EE"),
					VToken:          common.HexToAddress("0x8619d80FB0141ba7F184CbF22fd724116D9f7ffC"),
					UnderlyingToken: common.HexToAddress("0xd586E7F844cEa2F87f50152665BCbc2C279D8d70"),
					Decimals:        18,
				},

				"LINK": {
					Name:            "LINK",
					AToken:          common.HexToAddress("0x191c10Aa4AF7C30e871E70C95dB0E4eb77237530"),
					VToken:          common.HexToAddress("0x953A573793604aF8d41F306FEb8274190dB4aE0e"),
					UnderlyingToken: common.HexToAddress("0x5947BB275c521040051D82396192181b413227A3"),
					Decimals:        18,
				},
			},
		},
		constant.Gnosis: {
			Chain:     constant.Gnosis,
			L2Pool:    common.HexToAddress("0xb50201558B00496A145fE76f7424749556E326D8"),
			Oracle:    common.HexToAddress("0xeb0a051be10228213BAEb449db63719d6742F7c4"),
			Multicall: common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11"),
			DebtTokens: map[string]debtTokens{

				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0xc6B7AcA6DE8a6044E0e32d0c841a89244A10D284"),
					VToken:          common.HexToAddress("0x5F6f7B0a87CA3CF3d0b431Ae03EF3305180BFf4d"),
					UnderlyingToken: common.HexToAddress("0xDDAfbb505ad214D7b80b1f830fcCc89B60fb7A83"),
					Decimals:        6,
				},

				"DAI": {
					Name:            "DAI",
					AToken:          common.HexToAddress("0xd0Dd6cEF72143E22cCED4867eb0d5F2328715533"),
					VToken:          common.HexToAddress("0x281963D7471eCdC3A2Bd4503e24e89691cfe420D"),
					UnderlyingToken: common.HexToAddress("0xe91D153E0b41518A2Ce8Dd3D7944Fa863463a97d"),
					Decimals:        18,
				},
			},
		},
		constant.ArbitrumSepolia: {
			Chain:  constant.ArbitrumSepolia,
			L2Pool: common.HexToAddress("0xBfC91D59fdAA134A4ED45f7B584cAf96D7792Eff"),
			DebtTokens: map[string]debtTokens{
				"USDC": {
					Name:            "USDC",
					AToken:          common.HexToAddress("0x460b97BD498E1157530AEb3086301d5225b91216"),
					VToken:          common.HexToAddress("0x4fBE3A94C60A5085dA6a2D309965DcF34c36711d"),
					UnderlyingToken: common.HexToAddress("0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d"),
					Decimals:        6,
				},
			},
		},
	}
)

type AvaeConfig struct {
	Chain      string
	L2Pool     common.Address
	Oracle     common.Address
	Multicall  common.Address
	DebtTokens map[string]debtTokens
}

type debtTokens struct {
	Name            string
	AToken          common.Address
	VToken          common.Address
	UnderlyingToken common.Address
	Decimals        int32
}

type Aave struct {
}

func (a Aave) BalanceOf(ctx context.Context, args DebtParams) (decimal.Decimal, error) {
	conf, ok := aaveAddressBook[args.Chain]
	if !ok || conf.Chain == "" || conf.DebtTokens[args.Token].Name == "" {
		log.Warnf("chain %s token %s not support", args.Chain, args.Token)
		return decimal.Zero, nil
	}
	atokenBalance, err := chains.GetTokenBalance(ctx, args.Client, conf.DebtTokens[args.Token].AToken.Hex(),
		args.Address.Hex(), conf.DebtTokens[args.Token].Decimals)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get atoken balance error")
	}

	vtokenBalance, err := chains.GetTokenBalance(ctx, args.Client, conf.DebtTokens[args.Token].VToken.Hex(),
		args.Address.Hex(), conf.DebtTokens[args.Token].Decimals)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get vtoken balance error")
	}
	if args.TokenPrice.InexactFloat64() <= 0 {
		return atokenBalance.Sub(vtokenBalance), nil
	}
	pool, err := lendingpool.NewLendingpoolCaller(conf.L2Pool, args.Client)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "new lendingpool caller error")
	}
	user, err := pool.GetUserAccountData(nil, args.Address)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get user account data error")
	}
	if user.AvailableBorrowsBase.Int64() <= 0 || args.TokenPrice.Equal(decimal.RequireFromString("1")) {
		return atokenBalance.Sub(vtokenBalance), nil
	}
	AvailableBorrowsBalance := chains.WeiToEth(user.AvailableBorrowsBase, 8)
	return AvailableBorrowsBalance.Div(args.TokenPrice), nil

}

func (a Aave) Name() string {
	return "aave"
}
