package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"omni-balance/utils/chains"
	base_deposit "omni-balance/utils/enclave/router/base/deposit"
	base_withdraw "omni-balance/utils/enclave/router/base/withdraw"
	"omni-balance/utils/wallets"

	log "omni-balance/utils/logging"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	// Superbridge API配置
	superbridgeBaseURL = "https://api.superbridge.app"
	superbridgeHomeURL = "https://superbridge.app"

	// API路径
	apiURL  = superbridgeBaseURL + "/api/v6/bridge/activity"
	homeURL = superbridgeHomeURL + "/" // 用浏览器打开这个页面以触发 CF 挑战

	// 本地API配置
	// baseURL = "http://localhost:3009"
	baseURL = "http://common-rebalance"

	// API路径
	rebalanceBaseERC20DepositPath = "/rebalance/base-erc20-deposit"
)

var jsonBody = []byte(`{"id":{"tokensId":"895f6697-9cef-41d6-96ee-f3d9926f7a02"},"evmAddress":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","cursor":null,"filters":{"type":"mainnets"},"multichainTokens":[]}`)

func Approve(ctx context.Context, chainId int64, tokenAddress, spender common.Address, owner wallets.Wallets,
	amount decimal.Decimal, client simulated.Client,
) error {
	return chains.TokenApprove(ctx, chains.TokenApproveParams{
		ChainId:         chainId,
		TokenAddress:    tokenAddress,
		Owner:           owner.GetAddress(true),
		SendTransaction: owner.SendTransaction,
		WaitTransaction: owner.WaitTransaction,
		Spender:         spender,
		AmountWei:       amount,
		Client:          client,
	})
}

func Deposit(ctx context.Context, l1Address, l2Address, receiver common.Address, amount decimal.Decimal) ([]byte, error) {
	routerAbi, err := base_deposit.BaseDepositMetaData.GetAbi()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return routerAbi.Pack("depositERC20To", l1Address, l2Address, receiver, amount.BigInt(), uint32(200000), []byte{})
}

func Withdraw(ctx context.Context, l2Address, receiver common.Address, amount decimal.Decimal) ([]byte, error) {
	routerAbi, err := base_withdraw.BaseWithdrawMetaData.GetAbi()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return routerAbi.Pack("withdrawTo", l2Address, receiver, amount.BigInt(), uint32(200000), []byte{})
}

//	curl 'https://api.superbridge.app/api/v6/bridge/activity' \
//	  -H 'accept: application/json, text/plain, */*' \
//	  -H 'content-type: application/json' \
//	  -H 'origin: https://superbridge.app' \
//	  -H 'referer: https://superbridge.app/' \
//	  -H 'sec-fetch-mode: cors' \
//	  -H 'sec-fetch-site: same-site' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36' \
//	  --data-raw '{"id":{"tokensId":"895f6697-9cef-41d6-96ee-f3d9926f7a02"},"evmAddress":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","cursor":null,"filters":{"type":"mainnets"},"multichainTokens":[]}'
//
// {"total":16,"transactions":[{"id":"8e0f2bf1-1e88-49c3-912b-0a3aa0bc7c27","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1754463209000,"status":"confirmed","transactionHash":"0x25bd9b4d51d93d698a6fe97143a1c9c3f77f0541ecfee16b7b158e5e23da66ca"},"receive":{"timestamp":1754463227000,"status":"confirmed","transactionHash":"0xd5f3f03fe6e1a881eafb46c340a00b93be1e8e1f68c7563bb264b8b512dfa416"},"fromChainId":42161,"toChainId":1,"duration":60000,"token":"0xaf88d065e77c8cC2239327C5EDb3A432268e5831","receiveToken":"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48","amount":"30000000","receiveAmount":"29812202","fromToken":{"address":"0xaf88d065e77c8cC2239327C5EDb3A432268e5831","decimals":6,"name":"USD Coin","symbol":"USDC","coinGeckoId":"usd-coin","chainId":42161,"logoURI":"https://djvebdd83rbuw.cloudfront.net/tokens/usdc.png","bridges":[130,11155111,8453,10,84532,11155420,1,480,59144,56,1868,33979,42220,7777777,57073],"usd":0.999774},"toToken":{"address":"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48","decimals":6,"name":"USD Coin","symbol":"USDC","coinGeckoId":"usd-coin","chainId":1,"logoURI":"https://djvebdd83rbuw.cloudfront.net/tokens/usdc.png","bridges":[8866,7777777,185,7560,1750,1135,360,183,33979,1868,177,42161,130,11155111,8453,10,84532,11155420,480,59144,690,56,42220,57073],"usd":0.999774},"type":"across-bridge","provider":"Across"},{"id":"8ab78903-d24f-463f-a8aa-d439a1a6e156","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33","send":{"timestamp":1753336155000,"status":"confirmed","transactionHash":"0x15bdd2eb9b9660d5ffbf2aad2ab4592477ef78b51256bbb0f172dcd1d5d84be2"},"receive":{"timestamp":1753336175000,"status":"confirmed","transactionHash":"0x9c549799dd2a390b157073ec4b9ea51cc5b6c4f743b63e4e07c47b6304fa161f"},"fromChainId":42161,"toChainId":1,"duration":60000,"token":"0x82aF49447D8a07e3bd95BD0d56f35241523fBab1","receiveToken":"0x0000000000000000000000000000000000000000","amount":"8000000000000000","receiveAmount":"7925088682481150","fromToken":{"address":"0x82af49447d8a07e3bd95bd0d56f35241523fbab1","decimals":18,"name":"Wrapped Ether","symbol":"WETH","chainId":42161,"logoURI":"","bridges":[1,10,8453,59144,34443,1135,690,7777777,480,57073,1868,130,56],"usd":4246.72},"toToken":{"address":"0x0000000000000000000000000000000000000000","decimals":18,"name":"Ether","symbol":"ETH","coinGeckoId":"ethereum","chainId":1,"logoURI":"https://djvebdd83rbuw.cloudfront.net/tokens/eth.png","bridges":[690,8866,7777777,185,8453,34443,480,10,291,7560,1750,957,1135,360,254,60808,7897,888888888,183,33979,8008,5330,57073,1868,130,1923,42220,5371,42161,59144,56],"usd":4246.72},"type":"across-bridge","provider":"Across"},{"id":"67b9a7bd-69b3-4532-88f5-9725966b08aa","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1753150091000,"status":"confirmed","transactionHash":"0xa2483083f29142a30b0a0059aa9891885856fefee942c5fe722a38b814db263c"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0x2246cd996fb80771f5a63d07eba096a4331f15af32847685738199605f42ba80"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0x81e32d4652Be82AE225DEdd1bD0bf3BCba8FEE07","receiveToken":"0x77Eb290DF0a5aaB15f681085FaeA1F653A3fc9b7","amount":"50000000000000000000000","fromToken":{"address":"0x81e32d4652Be82AE225DEdd1bD0bf3BCba8FEE07","decimals":18,"name":"Darwinia Network xRING","symbol":"xRING","chainId":1,"logoURI":"https://ethereum-optimism.github.io/data/XRING/logo.svg","bridges":[10,8453]},"toToken":{"address":"0x77Eb290DF0a5aaB15f681085FaeA1F653A3fc9b7","decimals":18,"name":"Darwinia Network xRING","symbol":"xRING","chainId":8453,"logoURI":"https://ethereum-optimism.github.io/data/XRING/logo.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0x2246cd996fb80771f5a63d07eba096a4331f15af32847685738199605f42ba80","status":6},{"id":"2b216b33-2dff-419a-b6bf-991977a4519a","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1751594899000,"status":"confirmed","transactionHash":"0x544b3255958ab236241e81a70e2e93498068b60b884aaf164b85b2f19a8d2513"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":3,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb"},{"id":"b22b39d9-5909-47de-8da5-ff062260d467","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1751538045000,"status":"confirmed","transactionHash":"0x7855917130ebf82dc4c7fbe03d23febfab793d81dfed8ec72f10bc9bcda7fbd6"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":5,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb","prove":{"blockNumber":22842779,"transactionHash":"0x33eb8b346b4ea21b56e9cb3d7d55e2cb3c8aae9711013bbdc6be7a63335480d9","timestamp":1751595023000,"status":"confirmed","game":{"address":"0xa14805B3D667130f83dd5545d606C30b9cF5ac40","resolvedAt":1751843159,"createdAt":1751540699,"index":5874,"status":2,"maxClockDuration":302400}}},{"id":"fbfe3375-7cda-4f37-aa61-d2767c53d5d0","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1751533237000,"status":"confirmed","transactionHash":"0x418742d34e0797e823490aff5e4a26b57fab4e2c271eb8961fbeba4d6fbbec8e"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":5,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb","prove":{"blockNumber":22838089,"transactionHash":"0xcac2001fbddb14889b7c0566bac2e768d147163dc0ac4abbf4495e8c50a8563e","timestamp":1751538299000,"status":"confirmed","game":{"address":"0x184922A2c680B5aac9187C4210Ec27F4c24A4e1b","resolvedAt":1751839547,"createdAt":1751537087,"index":5873,"status":2,"maxClockDuration":302400}}},{"id":"6ea71fbc-06ca-4c1c-941c-628e5632410f","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750756883000,"status":"confirmed","transactionHash":"0x565bf0a49ef28529b9de2ab836b7cd760a7992a670f49bef6b1b2693fa257a63"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0x162b3d98146231f5225f6ab919b589563689832b7c2274f3acf86d7a5ac7880e"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"1000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0x162b3d98146231f5225f6ab919b589563689832b7c2274f3acf86d7a5ac7880e","status":6},{"id":"a4436864-9bef-4fe8-b446-d582c293dfb7","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750756583000,"status":"confirmed","transactionHash":"0x4779af585a1b0dd26fd0a06d56b4d480f5197448baed9ed5cb40b24aef058f22"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0x5d2853fa7426f9ce3142b615ee43a7734fba630e2b1df27079d99dee87c5f4b7"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"1000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0x5d2853fa7426f9ce3142b615ee43a7734fba630e2b1df27079d99dee87c5f4b7","status":6},{"id":"255b57a9-6a5f-41fd-b3ad-72113198292d","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750756271000,"status":"confirmed","transactionHash":"0xfcbc40ff52a463b903bb52a3976f9d3ea59aeeeadf3085dcf764c6e39ea97eab"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0xa2e4c20854664ff50a8f427b49d2803d982fcffa5386d9b21f9cd46203b00c4b"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"1000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0xa2e4c20854664ff50a8f427b49d2803d982fcffa5386d9b21f9cd46203b00c4b","status":6},{"id":"33771ec3-0e93-4729-9b2d-e655da8950af","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750755683000,"status":"confirmed","transactionHash":"0x6ea1c72f7ed661e5dc47920eca44c4315c9c30421d7d08c8727d95c5eb021c28"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0x27b2e1fd90fc48dc76b2d7618866c6dfd82370c9baf04c1035b923eabae59efa"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"10000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0x27b2e1fd90fc48dc76b2d7618866c6dfd82370c9baf04c1035b923eabae59efa","status":6},{"id":"1cdc43dc-f6d8-4343-a47a-33266218f84a","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750755647000,"status":"confirmed","transactionHash":"0x945fc7f1d36c757b057a69ae97db5e0652c5bcf5e82ca358b7e52ada788c4367"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0xea602114b6903a46902c4c1896289ec898ca5ccf1da04277399ecd5a0e63bd34"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"10000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0xea602114b6903a46902c4c1896289ec898ca5ccf1da04277399ecd5a0e63bd34","status":6},{"id":"7a750963-5363-4ebd-b7ab-6a0be8795bf5","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750744387000,"status":"confirmed","transactionHash":"0x089eff1047085ebce995ac49f756a757e1c7e02d2ef1ac315e29900bb8d1c617"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":5,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb","prove":{"blockNumber":22773401,"transactionHash":"0x4f92186ccfa910171dde7cef896ef55496af7ba287ee2b206d1ca552c9c28d02","timestamp":1750757399000,"status":"confirmed","game":{"address":"0x5a3CF23F47486380098173174Aab26cCc26B060C","resolvedAt":1751048507,"createdAt":1750746047,"index":5654,"status":2,"maxClockDuration":302400}}},{"id":"dcbabefe-f203-4b1f-8420-4051a2af51b1","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750743535000,"status":"confirmed","transactionHash":"0x3350bcd4783e2cd932f90d9517c4fc4e1eca993d45bfba7db8d1af04d4209d88"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":5,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb","prove":{"blockNumber":22837909,"transactionHash":"0xb3ebf9a5f31669f52c59d1e7d08edd28e2460c6beb084dd3cc40c8c3a466bf6d","timestamp":1751536127000,"status":"confirmed","game":{"address":"0x15fA9f1AAC4C1dB1EA6F5b6A8534f0580c9a1892","resolvedAt":1751532563,"createdAt":1751230079,"index":5788,"status":2,"maxClockDuration":302400}}},{"id":"afd58c4f-c47b-40b9-97f2-1fa8e66e9a96","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","send":{"timestamp":1750742541000,"status":"confirmed","transactionHash":"0x1ae17c18157bedc4af66441b2113ca160082d24a561b4c20c863c105f85a8bb8"},"receive":{"timestamp":22842935000,"status":"confirmed","transactionHash":"0x10841a7e47a0ba2f8f192f4cac7757c58cf9226fd31f0acfd07b17d7a8a65e5e"},"fromChainId":8453,"toChainId":1,"duration":1,"token":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","receiveToken":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","amount":"1000000000000000000","fromToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"toToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"type":"withdrawal","provider":"OptimismWithdrawal","status":6,"deploymentId":"81883861-df09-4a49-816e-7268435d27eb","prove":{"blockNumber":22773056,"transactionHash":"0x53e9607c8487b7571220d15372cc86bf3f6ad76322bc0b2c6223ebe886faddef","timestamp":1750753247000,"status":"confirmed","game":{"address":"0x5a3CF23F47486380098173174Aab26cCc26B060C","resolvedAt":1751048507,"createdAt":1750746047,"index":5654,"status":2,"maxClockDuration":302400}}},{"id":"a536a539-9f9c-44ab-a401-95fccddd5832","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33","send":{"timestamp":1749708299000,"status":"confirmed","transactionHash":"0x6b1115f3ba8f76b42e483d1c0eb1300f18f750dd5404644c6dd20bd7b21771f6"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"24999000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db","status":6},{"id":"050e500e-8b66-415e-ae6a-db6868c1b741","from":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","to":"0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33","send":{"timestamp":1749707663000,"status":"confirmed","transactionHash":"0xf8f34f4ad0f934be684e1a0d1cda6c12c5af28c23b1c737f0bc4f1d121bf5c15"},"receive":{"timestamp":1,"status":"confirmed","transactionHash":"0xa04e242ad0e8ea0a74fa590dcf0e7a2213c6ee20a23adbd6865661b2ba5c601f"},"fromChainId":1,"toChainId":8453,"duration":1,"token":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","receiveToken":"0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69","amount":"1000000000000000000","fromToken":{"address":"0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":1,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[8453]},"toToken":{"address":"0xc694a91e6b071bf030a18bd3053a7fe09b6dae69","decimals":18,"name":"CoW Protocol Token","symbol":"COW","chainId":8453,"logoURI":"https://res.cloudinary.com/supermigrate/image/upload/v1729784219/migrations/aajuoiceu0sikvsjcnfn.svg","bridges":[1]},"type":"deposit","provider":"OptimismDeposit","deploymentId":"81883861-df09-4a49-816e-7268435d27eb","l2TransactionHash":"0xa04e242ad0e8ea0a74fa590dcf0e7a2213c6ee20a23adbd6865661b2ba5c601f","status":6}],"actionRequiredCount":5,"inProgressCount":5,"hasWithdrawalReadyToFinalize":null,"recipients":["1:0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","1:0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33","8453:0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","8453:0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33"]}
type ActivityRequest struct {
	ID            string   `json:"id,omitempty"`
	EvmAddress    string   `json:"evmAddress"`
	DeploymentIds []string `json:"deploymentIds"`
}

type ActivityResponse struct {
	Total        int           `json:"total"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
	Send struct {
		Timestamp       int64  `json:"timestamp"`
		Status          string `json:"status"`
		TransactionHash string `json:"transactionHash"`
	} `json:"send"`
	Receive struct {
		Timestamp       int64  `json:"timestamp"`
		Status          string `json:"status"`
		TransactionHash string `json:"transactionHash"`
	} `json:"receive"`
	FromChainID   int    `json:"fromChainId"`
	ToChainID     int    `json:"toChainId"`
	Duration      int    `json:"duration"`
	Token         string `json:"token"`
	ReceiveToken  string `json:"receiveToken"`
	Amount        string `json:"amount"`
	ReceiveAmount string `json:"receiveAmount"`
	FromToken     Token  `json:"fromToken"`
	ToToken       Token  `json:"toToken"`
	Type          string `json:"type"`
	Provider      string `json:"provider"`
}

type Token struct {
	Address     string  `json:"address"`
	Decimals    int     `json:"decimals"`
	Name        string  `json:"name"`
	Symbol      string  `json:"symbol"`
	CoinGeckoID string  `json:"coinGeckoId"`
	ChainID     int     `json:"chainId"`
	LogoURI     string  `json:"logoURI"`
	Bridges     []int   `json:"bridges"`
	USD         float64 `json:"usd"`
}

// curl http://localhost:3009/rebalance/base-erc20-deposit/0x7fdb54d91973eed12b2de36d165c9e2ee3f9e54871325f0fd544a6e3a534b1e1
// 0x7adc7f454b38f4df4c16e9b07ba6d02215f728348b5770d0e1b9f1b18cb1b381
func WaitForChildTransactionReceipt(ctx context.Context, depositTxHash, trader string) (string, error) {
	// 根据注释，这个函数应该调用本地API来获取子交易收据
	// 注释显示调用: http://localhost:3009/rebalance/base-erc20-deposit/{depositTxHash}
	// 返回: 0x7adc7f454b38f4df4c16e9b07ba6d02215f728348b5870d0e1b9f1b18cb1b381

	localAPIURL := fmt.Sprintf("%s%s/%s", baseURL, rebalanceBaseERC20DepositPath, depositTxHash)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", localAPIURL, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 打印响应体
	fmt.Println(string(respBody))

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API返回错误状态码: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	// 解析响应，根据注释返回的是交易哈希字符串
	// 去除可能的空白字符
	childTxHash := strings.TrimSpace(string(respBody))

	// 验证返回的是否为有效的以太坊地址格式
	if !strings.HasPrefix(childTxHash, "0x") || len(childTxHash) != 66 {
		return "", fmt.Errorf("返回的子交易哈希格式无效: %s", childTxHash)
	}

	log.Infof("成功获取子交易收据: %s", childTxHash)
	return childTxHash, nil
}

func WaitForProve(ctx context.Context, withdrawTx, trader string) (string, error) {
	// 先尝试一次
	proveId, err := getProve(ctx, withdrawTx, trader)
	if err != nil {
		fmt.Println("getProve error:", err)
	} else if proveId != "" {
		proveData, err := getProveData(ctx, proveId, trader)
		if err != nil {
			fmt.Println("getProveDta error:", err)
		} else {
			return proveData, nil
		}
	}

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			proveId, err := getProve(ctx, withdrawTx, trader)
			if err != nil {
				fmt.Println("getProve error:", err)
				continue
			}
			if proveId != "" {
				proveData, err := getProveData(ctx, proveId, trader)
				if err != nil {
					fmt.Println("getProveDta error:", err)
					continue
				}
				return proveData, nil
			}
		}
	}
}

func getProveData(ctx context.Context, proveId, trader string) (string, error) {
	return getData(ctx, proveId, "op_prove")
}

func getClaimData(ctx context.Context, proveId, trader string) (string, error) {
	return getData(ctx, proveId, "op_finalise")
}

//	curl 'https://api.superbridge.app/api/bridge/op_prove' \
//		-H 'content-type: application/json' \
//		-H 'origin: https://superbridge.app' \
//		--data-raw '{"id":"dcbabefe-f203-4b1f-8420-4051a2af51b1"}'
//
// {"to":"0x49048044D57e1C92A77f79988d21Fa8fAF74E97e","data":"0x4870496f","chainId":1}⏎
func getData(ctx context.Context, proveId string, method string) (string, error) {
	// 获取绕过 Cloudflare 的客户端
	client, err := BypassCloudflare()
	if err != nil {
		log.Errorf("绕过 Cloudflare 失败: %v", err)
		// 如果绕过失败，回退到普通客户端
		client = &http.Client{}
	}

	url := superbridgeBaseURL + "/api/bridge/op_prove"
	body := fmt.Sprintf(`{"id":"%s"}`, proveId)

	for {
		req, err := http.NewRequest("POST", url, strings.NewReader(body))
		if err != nil {
			return "", err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Origin", superbridgeHomeURL)
		req.Header.Set("Referer", superbridgeHomeURL+"/")
		req.Header.Set("Accept", "application/json, text/plain, */*")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

		resp, err := client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", err
			}
			var result struct {
				To      string `json:"to"`
				Data    string `json:"data"`
				ChainId uint32 `json:"chainId"`
			}
			if err := json.Unmarshal(respBytes, &result); err != nil {
				return "", err
			}
			return result.Data, nil
		}

		if err != nil {
			fmt.Println("getProveData http error:", err)
		} else {
			fmt.Println("getProveData status:", resp.Status)
			resp.Body.Close()
		}

		// 10分钟后重试，或ctx被取消
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(10 * time.Minute):
		}
	}
}

func getProve(ctx context.Context, withdrawTx, trader string) (string, error) {
	return getId(ctx, withdrawTx, trader, 3)
}

func getClaim(ctx context.Context, proveTx, trader string) (string, error) {
	return getId(ctx, proveTx, trader, 5)
}

func getId(ctx context.Context, txHash, trader string, status uint32) (string, error) {
	// 获取绕过 Cloudflare 的客户端
	client, err := BypassCloudflare()
	if err != nil {
		log.Errorf("绕过 Cloudflare 失败: %v", err)
		// 如果绕过失败，回退到普通客户端
		client = &http.Client{}
	}

	requestBody := map[string]interface{}{
		"id": map[string]interface{}{
			"tokensId": "895f6697-9cef-41d6-96ee-f3d9926f7a02",
		},
		"evmAddress":       trader,
		"cursor":           nil,
		"filters":          map[string]interface{}{"type": "mainnets"},
		"multichainTokens": []interface{}{},
	}
	body, _ := json.Marshal(requestBody)

	log.Infof("request: %s", string(body))

	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", superbridgeHomeURL)
	req.Header.Set("Referer", superbridgeHomeURL+"/")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var activity ActivityResponse
	if err := json.Unmarshal(respBody, &activity); err != nil {
		return "", err
	}
	log.Infof("response: %s", respBody)

	for _, tx := range activity.Transactions {
		if status == 3 {
			if tx.Send.TransactionHash == txHash {
				if tx.Send.Status == "confirmed" {
					return tx.ID, nil
				}
			}
		} else if status == 5 {
			if tx.Receive.TransactionHash == txHash {
				if tx.Receive.Status == "confirmed" {
					return tx.ID, nil
				}
			}
		} else {
			return "", errors.New("unknown state")
		}
	}
	log.Infof("still waiting for get status %d", status)
	return "", errors.New("still waiting for get status")
}

func WaitForClaim(ctx context.Context, proveTx, trader string) (string, error) {
	// 先尝试一次
	claimId, err := getClaim(ctx, proveTx, trader)
	if err != nil {
		fmt.Println("getClaim error:", err)
	} else if claimId != "" {
		claimData, err := getClaimData(ctx, claimId, trader)
		if err != nil {
			fmt.Println("getClaimDta error:", err)
		} else {
			return claimData, nil
		}
	}

	ticker := time.NewTicker(6 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			claimId, err := getClaim(ctx, proveTx, trader)
			if err != nil {
				fmt.Println("getClaim error:", err)
				continue
			}
			if claimId != "" {
				claimData, err := getClaimData(ctx, claimId, trader)
				if err != nil {
					fmt.Println("getClaimDta error:", err)
					continue
				}
				return claimData, nil
			}
		}
	}
}

// BypassCloudflare 使用 chromedp 绕过 Cloudflare 挑战并获取 cookies
func BypassCloudflare() (*http.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 90*time.Second)
	defer cancel()

	// 启动 chromedp（无头模式）。如果想看浏览器过程，把 chromedp.Flag("headless", true) 改为 false。
	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		// chromedp.Headless, // 无头。要可视化调试请注释掉这行
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("disable-setuid-sandbox", true),
	}

	// 检查环境变量中的 Chrome 路径（Docker 环境）
	if chromePath := os.Getenv("CHROME_PATH"); chromePath != "" {
		if _, err := os.Stat(chromePath); err == nil {
			opts = append(opts, chromedp.ExecPath(chromePath))
		}
	}

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	cctx, ccancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Infof))
	defer ccancel()

	// 访问主页并等待 Cloudflare 验证
	if err := chromedp.Run(cctx,
		chromedp.Navigate(homeURL),
	); err != nil {
		return nil, fmt.Errorf("导航失败: %v", err)
	}

	// 等待并轮询 cookies，直到出现 cf_clearance 或超时
	var cookies []*network.Cookie
	found := false
	waitCtx, waitCancel := context.WithTimeout(cctx, 45*time.Second)
	defer waitCancel()

	for {
		if err := chromedp.Run(waitCtx, chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			cookies, err = network.GetCookies().Do(ctx)
			return err
		})); err != nil {
			// 超时或其他错误
			break
		}

		// 检查是否包含 cf_clearance 或其他 CF cookie
		for _, c := range cookies {
			if c.Name == "cf_clearance" || c.Name == "__cf_bm" {
				found = true
				break
			}
		}
		if found {
			break
		}
		// 若还没出现，短暂 sleep 再轮询
		select {
		case <-time.After(1 * time.Second):
			// 再试
		case <-waitCtx.Done():
			break
		}
		if waitCtx.Err() != nil {
			break
		}
	}

	if !found {
		log.Infof("未检测到 cf_clearance/__cf_bm 等 Cloudflare cookie，可能需要手动完成挑战或延长超时。程序仍将尝试使用现有 cookie（若有）继续请求。")
	} else {
		log.Infof("检测到 Cloudflare cookie，准备用它重放 API 请求。")
	}

	// 将 chromedp 的 cookies 转换并放入 http.Client 的 CookieJar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:     jar,
		Timeout: 30 * time.Second,
	}

	// set cookies for api domain
	u, _ := url.Parse(apiURL)
	var httpCookies []*http.Cookie
	log.Infof("cookies: %v", cookies)
	for _, c := range cookies {
		log.Infof("cookie: %s = %s", c.Name, c.Value)
		// 只添加与目标域相关或通用的 cookie
		// 注意：chromedp.Cookies 返回的 Domain 可能是 .superbridge.app, 需要 set 到目标URL
		// 转换 Expires 字段（从 float64 到 time.Time）
		var expires time.Time
		if c.Expires > 0 {
			expires = time.Unix(int64(c.Expires), 0)
		}

		httpCookies = append(httpCookies, &http.Cookie{
			Name:     c.Name,
			Value:    c.Value,
			Path:     c.Path,
			Domain:   c.Domain,
			Expires:  expires,
			Secure:   c.Secure,
			HttpOnly: c.HTTPOnly,
		})
	}
	jar.SetCookies(u, httpCookies)

	return client, nil
}

// MakeRequestWithCloudflareBypass 使用绕过 Cloudflare 的客户端发送请求
func MakeRequestWithCloudflareBypass() error {
	// 获取绕过 Cloudflare 的客户端
	client, err := BypassCloudflare()
	if err != nil {
		return fmt.Errorf("绕过 Cloudflare 失败: %v", err)
	}

	// 构造 POST 请求（包含常见浏览器头）
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("构造请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", superbridgeHomeURL)
	req.Header.Set("Referer", superbridgeHomeURL+"/")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")
	// 你还可以设置 sec-ch-ua 等 header，但一般不必要，关键是带上 cookie 和常见头

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求 API 失败: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Println("状态码:", resp.Status)
	fmt.Println("Response body:")
	fmt.Println(string(bodyBytes))

	return nil
}
