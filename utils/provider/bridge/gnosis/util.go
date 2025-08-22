package gnosis

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"omni-balance/utils/chains"
	gnosis_claim "omni-balance/utils/enclave/router/gnosis/claim"
	gnosis_deposit "omni-balance/utils/enclave/router/gnosis/deposit"
	gnosis_transmuter "omni-balance/utils/enclave/router/gnosis/transmuter"
	"omni-balance/utils/erc20"
	"omni-balance/utils/wallets"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

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

func TransmuterWithdraw(ctx context.Context, value decimal.Decimal) ([]byte, error) {
	routerAbi, err := gnosis_transmuter.GnosisTransmuterMetaData.ParseABI()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return routerAbi.Pack("withdraw", value.BigInt())
}

func Deposit(ctx context.Context, token, receiver common.Address, value decimal.Decimal, data []byte) ([]byte, error) {
	routerAbi, err := gnosis_deposit.GnosisDepositMetaData.ParseABI()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return routerAbi.Pack("relayTokensAndCall", token, receiver, value.BigInt(), data)
}

func Withdraw(ctx context.Context, to common.Address, value decimal.Decimal, data []byte) ([]byte, error) {
	erc1363Abi, err := erc20.Erc1363MetaData.GetAbi()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return erc1363Abi.Pack("transferAndCall0", to, value.BigInt(), data)
}

// 第一步请求 API ，通过 depositTxHash == transactionHash 获取对应的 id
//
//	curl 'https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/9W7Ye5xFfefNYDxXD4StqAuj7TU8eLq5PLmuPUnhFbeQ' \
//	  -H 'accept: */*' \
//	  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
//	  -H 'content-type: application/json' \
//	  -H 'origin: https://bridge.gnosischain.com' \
//	  -H 'priority: u=1, i' \
//	  -H 'referer: https://bridge.gnosischain.com/' \
//	  -H 'sec-ch-ua: "Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'sec-fetch-dest: empty' \
//	  -H 'sec-fetch-mode: cors' \
//	  -H 'sec-fetch-site: cross-site' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36' \
//	  --data-raw '{"query":"query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {\n  transactions(\n    where: $where\n    orderBy: $orderBy\n    orderDirection: $orderDirection\n    first: $first\n    skip: $skip\n  ) {\n    ...TransactionFragment\n  }\n}\n\nfragment TransactionFragment on Transaction {\n  id\n  bridgeName\n  transactionHash\n  initiator\n  initiatorAmount\n  initiatorNetwork\n  initiatorToken\n  receiver\n  receiverToken\n  receiverAmount\n  receiverNetwork\n  transactionStatus\n  timestamp\n  execution {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n  validations {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n}","variables":{"orderBy":"timestamp","orderDirection":"desc","first":1000,"skip":0,"where":{"and":[{"transactionHash":"0xaee189b6f370812009d852efbb5cc04a1e15d85887fb758c28e7c32a5826ab87"}]}},"operationName":"Transactions"}'
//
// {"data":{"transactions":[{"bridgeName":"AMB","execution":null,"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61","initiator":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","initiatorAmount":"1000000000000000000","initiatorNetwork":"gnosis","initiatorToken":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","receiver":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","receiverAmount":"1000000000000000000","receiverNetwork":"mainnet","receiverToken":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","timestamp":"1754473600","transactionHash":"0xaee189b6f370812009d852efbb5cc04a1e15d85887fb758c28e7c32a5826ab87","transactionStatus":"UNCLAIMED","validations":[{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61-0x258667e543c913264388b33328337257af208a8f","timestamp":"1754474135","transactionHash":"0x647fc7e1888720f5858bce9a2f2b1a584903747685a256752e9bc34a0ef450ce","validatorAddr":"0x258667e543c913264388b33328337257af208a8f"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61-0x3e0a20099626f3d4d4ea7b0ce0330e88d1fe65d6","timestamp":"1754474135","transactionHash":"0x11a365b00daea841b3baf14afda15cbf91432e47cb1129da8c75ceaeb3204eab","validatorAddr":"0x3e0a20099626f3d4d4ea7b0ce0330e88d1fe65d6"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61-0x459a3bd49f1ff109bc90b76125533699aaaaf9a6","timestamp":"1754474135","transactionHash":"0x1b63412bd91fc3be90dd13fca4404d7a61381e36195ab78af0016d43d5d2c861","validatorAddr":"0x459a3bd49f1ff109bc90b76125533699aaaaf9a6"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61-0xbdc141c8d2343f33f40cb9edd601ccf460cd0dde","timestamp":"1754474135","transactionHash":"0xa1865e33459b23ec5917d72677e6f89fb189fb4ed3117a70ef13b240ee77515f","validatorAddr":"0xbdc141c8d2343f33f40cb9edd601ccf460cd0dde"}]}]}}
// 第二步请求 API, 通过 id 获取对应的 execution.transactionHash 作为 childTransactionHash
//
//	curl 'https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ' \
//	  -H 'accept: */*' \
//	  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
//	  -H 'cache-control: no-cache' \
//	  -H 'content-type: application/json' \
//	  -H 'origin: https://bridge.gnosischain.com' \
//	  -H 'pragma: no-cache' \
//	  -H 'priority: u=1, i' \
//	  -H 'referer: https://bridge.gnosischain.com/' \
//	  -H 'sec-ch-ua: "Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'sec-fetch-dest: empty' \
//	  -H 'sec-fetch-mode: cors' \
//	  -H 'sec-fetch-site: cross-site' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36' \
//	  --data-raw '{"query":"query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {\n  transactions(\n    where: $where\n    orderBy: $orderBy\n    orderDirection: $orderDirection\n    first: $first\n    skip: $skip\n  ) {\n    ...TransactionFragment\n  }\n}\n\nfragment TransactionFragment on Transaction {\n  id\n  bridgeName\n  transactionHash\n  initiator\n  initiatorAmount\n  initiatorNetwork\n  initiatorToken\n  receiver\n  receiverToken\n  receiverAmount\n  receiverNetwork\n  transactionStatus\n  timestamp\n  execution {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n  validations {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n}","variables":{"where":{"id":"0x000500004ac82b41bd819dd871590b510316f2385cb196fb000000000002bd61"},"skip":0,"first":1000},"operationName":"Transactions"}'
//
// {"data":{"transactions":[{"bridgeName":"AMB","execution":{"id":"0x000500004ac82b41bd819dd871590b510316f2385cb196fb000000000002bd61-0xf6a78083ca3e2a662d6dd1703c939c8ace2e268d","timestamp":"1755675245","transactionHash":"0xebbaa005387b4e91110fa25e24e6be0bd1c4529e2c37cd34d0fd1318ab5a0864","validatorAddr":"0x105cd22ed3d089bf5589c59b452f9de0796ca52d"},"id":"0x000500004ac82b41bd819dd871590b510316f2385cb196fb000000000002bd61","initiator":null,"initiatorAmount":null,"initiatorNetwork":"mainnet","initiatorToken":null,"receiver":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","receiverAmount":"1000000000000000000","receiverNetwork":"gnosis","receiverToken":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","timestamp":null,"transactionHash":null,"transactionStatus":"COMPLETED","validations":[{"id":"0xb21cf89b23eb95188ae9a05f38b215a2a2094c5ecb5b155b03c8626dd49537e0","timestamp":"1755675235","transactionHash":"0xb21cf89b23eb95188ae9a05f38b215a2a2094c5ecb5b155b03c8626dd49537e0","validatorAddr":"0x459a3bd49f1ff109bc90b76125533699aaaaf9a6"},{"id":"0xbdb83bed982a5d223721dee8bdfecc0e194d90be553b876d43dbb92b3d8dccd9","timestamp":"1755675235","transactionHash":"0xbdb83bed982a5d223721dee8bdfecc0e194d90be553b876d43dbb92b3d8dccd9","validatorAddr":"0x3e0a20099626f3d4d4ea7b0ce0330e88d1fe65d6"},{"id":"0xebbaa005387b4e91110fa25e24e6be0bd1c4529e2c37cd34d0fd1318ab5a0864","timestamp":"1755675245","transactionHash":"0xebbaa005387b4e91110fa25e24e6be0bd1c4529e2c37cd34d0fd1318ab5a0864","validatorAddr":"0x105cd22ed3d089bf5589c59b452f9de0796ca52d"}]}]}}
// GraphQL 查询结构体
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type GraphQLResponse struct {
	Data struct {
		Transactions []struct {
			ID                string `json:"id"`
			BridgeName        string `json:"bridgeName"`
			TransactionHash   string `json:"transactionHash"`
			TransactionStatus string `json:"transactionStatus"`
			Execution         *struct {
				ID              string `json:"id"`
				Timestamp       string `json:"timestamp"`
				TransactionHash string `json:"transactionHash"`
				ValidatorAddr   string `json:"validatorAddr"`
			} `json:"execution"`
		} `json:"transactions"`
	} `json:"data"`
}

// WaitForChildTransactionReceipt 等待并获取 Gnosis 跨链桥的子交易收据
// 第一步：通过 depositTxHash 获取交易 ID
// 第二步：通过 ID 获取 execution.transactionHash 作为 childTransactionHash
func WaitForChildTransactionReceipt(ctx context.Context, depositTxHash, trader string) (string, error) {
	return WaitForChildTransactionReceiptWithCookie(ctx, depositTxHash, trader, "")
}

// WaitForChildTransactionReceiptWithCookie 带自定义 Cookie 的版本
func WaitForChildTransactionReceiptWithCookie(ctx context.Context, depositTxHash, trader, customCookies string) (string, error) {
	const (
		// GraphQL 端点
		graphqlEndpoint1 = "https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/9W7Ye5xFfefNYDxXD4StqAuj7TU8eLq5PLmuPUnhFbeQ"
		graphqlEndpoint2 = "https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ"
		// 查询间隔
		pollInterval = 10 * time.Minute
		// 最大等待时间
		maxWaitTime = 10 * time.Minute
	)

	// 第一步：通过 depositTxHash 获取交易 ID
	transactionID, err := getTransactionID(ctx, graphqlEndpoint1, depositTxHash, customCookies)
	if err != nil {
		return "", errors.Wrap(err, "failed to get transaction ID")
	}

	fmt.Println("transactionID:", transactionID)

	// 第二步：等待并获取子交易哈希
	childTxHash, err := waitForChildTransaction(ctx, graphqlEndpoint2, transactionID, pollInterval, maxWaitTime, customCookies)
	if err != nil {
		return "", errors.Wrap(err, "failed to get child transaction hash")
	}

	return childTxHash, nil
}

// getTransactionID 通过 depositTxHash 获取交易 ID
func getTransactionID(ctx context.Context, endpoint, depositTxHash, customCookies string) (string, error) {
	query := `
		query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {
			transactions(
				where: $where
				orderBy: $orderBy
				orderDirection: $orderDirection
				first: $first
				skip: $skip
			) {
				id
				bridgeName
				transactionHash
				initiator
				initiatorAmount
				initiatorNetwork
				initiatorToken
				receiver
				receiverToken
				receiverAmount
				receiverNetwork
				transactionStatus
				timestamp
				execution {
					id
					timestamp
					transactionHash
					validatorAddr
				}
				validations {
					id
					timestamp
					transactionHash
					validatorAddr
				}
			}
		}`

	variables := map[string]interface{}{
		"orderBy":        "timestamp",
		"orderDirection": "desc",
		"first":          1000,
		"skip":           0,
		"where": map[string]interface{}{
			"and": []map[string]interface{}{
				{"transactionHash": depositTxHash},
			},
		},
	}

	req := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	resp, err := makeGraphQLRequest(ctx, endpoint, req, customCookies)
	if err != nil {
		return "", err
	}

	if len(resp.Data.Transactions) == 0 {
		return "", errors.Errorf("no transaction found for hash: %s", depositTxHash)
	}

	return resp.Data.Transactions[0].ID, nil
}

// waitForChildTransaction 等待并获取子交易哈希
func waitForChildTransaction(ctx context.Context, endpoint, transactionID string, pollInterval, maxWaitTime time.Duration, customCookies string) (string, error) {
	query := `
		query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {
			transactions(
				where: $where
				orderBy: $orderBy
				orderDirection: $orderDirection
				first: $first
				skip: $skip
			) {
				id
				bridgeName
				transactionHash
				initiator
				initiatorAmount
				initiatorNetwork
				initiatorToken
				receiver
				receiverToken
				receiverAmount
				receiverNetwork
				transactionStatus
				timestamp
				execution {
					id
					timestamp
					transactionHash
					validatorAddr
				}
				validations {
					id
					timestamp
					transactionHash
					validatorAddr
				}
			}
		}`

	variables := map[string]interface{}{
		"where": map[string]interface{}{
			"id": transactionID,
		},
		"skip":  0,
		"first": 1000,
	}

	req := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	startTime := time.Now()
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			if time.Since(startTime) > maxWaitTime {
				return "", errors.New("timeout waiting for child transaction")
			}

			resp, err := makeGraphQLRequest(ctx, endpoint, req, customCookies)
			if err != nil {
				continue // 继续等待
			}

			if len(resp.Data.Transactions) > 0 {
				tx := resp.Data.Transactions[0]
				if tx.Execution != nil && tx.Execution.TransactionHash != "" {
					return tx.Execution.TransactionHash, nil
				}
			}
		}
	}
}

// makeGraphQLRequest 发送 GraphQL 请求
func makeGraphQLRequest(ctx context.Context, endpoint string, req GraphQLRequest, customCookies string) (*GraphQLResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request")
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}

	// 设置完整的浏览器请求头，保持与浏览器一致
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "*/*")
	httpReq.Header.Set("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	httpReq.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36")
	httpReq.Header.Set("Referer", "https://bridge.gnosischain.com/")
	httpReq.Header.Set("Origin", "https://bridge.gnosischain.com")
	httpReq.Header.Set("Sec-Ch-Ua", `"Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"`)
	httpReq.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	httpReq.Header.Set("Sec-Ch-Ua-Platform", `"macOS"`)
	httpReq.Header.Set("Sec-Fetch-Dest", "empty")
	httpReq.Header.Set("Sec-Fetch-Mode", "cors")
	httpReq.Header.Set("Sec-Fetch-Site", "cross-site")
	httpReq.Header.Set("Priority", "u=1, i")

	// 设置 Cookie，优先使用自定义的 Cookie
	if customCookies != "" {
		httpReq.Header.Set("Cookie", customCookies)
	} else {
		// 设置默认的 Cookie
		// defaultCookies := []string{
		// 	"__Secure-next-auth.callback-url=https://bridge.gnosischain.com",
		// 	"__Secure-next-auth.csrf-token=random_token_here",
		// 	"__Host-next-auth.csrf-token=random_token_here",
		// }
		// httpReq.Header.Set("Cookie", strings.Join(defaultCookies, "; "))
	}

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var graphqlResp GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&graphqlResp); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return &graphqlResp, nil
}

// 第一步 查询API，根据data.transactions[0].transactionStatus == "UNCLAIMED" 状态后可进行下一步操作
//
//	curl 'https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ' \
//	  -H 'accept: */*' \
//	  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
//	  -H 'cache-control: no-cache' \
//	  -H 'content-type: application/json' \
//	  -H 'origin: https://bridge.gnosischain.com' \
//	  -H 'pragma: no-cache' \
//	  -H 'priority: u=1, i' \
//	  -H 'referer: https://bridge.gnosischain.com/' \
//	  -H 'sec-ch-ua: "Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'sec-fetch-dest: empty' \
//	  -H 'sec-fetch-mode: cors' \
//	  -H 'sec-fetch-site: cross-site' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36' \
//	  --data-raw '{"query":"query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {\n  transactions(\n    where: $where\n    orderBy: $orderBy\n    orderDirection: $orderDirection\n    first: $first\n    skip: $skip\n  ) {\n    ...TransactionFragment\n  }\n}\n\nfragment TransactionFragment on Transaction {\n  id\n  bridgeName\n  transactionHash\n  initiator\n  initiatorAmount\n  initiatorNetwork\n  initiatorToken\n  receiver\n  receiverToken\n  receiverAmount\n  receiverNetwork\n  transactionStatus\n  timestamp\n  execution {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n  validations {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n}","variables":{"where":{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57"},"skip":0,"first":1000},"operationName":"Transactions"}'
//
// {"data":{"transactions":[{"bridgeName":"AMB","execution":null,"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57","initiator":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","initiatorAmount":"1000000000000000000","initiatorNetwork":"gnosis","initiatorToken":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","receiver":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","receiverAmount":"1000000000000000000","receiverNetwork":"mainnet","receiverToken":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","timestamp":"1755684880","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","transactionStatus":"UNCLAIMED","validations":[{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57-0x3e0a20099626f3d4d4ea7b0ce0330e88d1fe65d6","timestamp":"1755685420","transactionHash":"0x23135e6624425e4a91d303162d4e97fbe08540fed9c6712f9a84c40de0e2e9cf","validatorAddr":"0x3e0a20099626f3d4d4ea7b0ce0330e88d1fe65d6"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57-0x459a3bd49f1ff109bc90b76125533699aaaaf9a6","timestamp":"1755685420","transactionHash":"0x5e619bc2d712e57c8baadfef509ed797dbcdfaae6ad2662abb9a04b3c6a3f6cb","validatorAddr":"0x459a3bd49f1ff109bc90b76125533699aaaaf9a6"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57-0x674c97db4ce6cac04a124d745979f3e4cba0e9f0","timestamp":"1755685425","transactionHash":"0x091f5581a73a3b48e3348a543baf2722a33989e8ee9eeefc6f8a9a2c5b7b7ca2","validatorAddr":"0x674c97db4ce6cac04a124d745979f3e4cba0e9f0"},{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57-0xfa98b60e02a61b6590f073cad56e68326652d094","timestamp":"1755685425","transactionHash":"0xbcfe78437b15d7770e6c572fe5584eaf2db3a3f7ac2d59d0e5f66f0c2720fab9","validatorAddr":"0xfa98b60e02a61b6590f073cad56e68326652d094"}]}]}}
// 第二步，通过查询 ethClient.getTransactionReceipt 查询 withdrawTx 的交易receipt，
// 根据 receipt logs 中获取 topic `UserRequestForSignature (index_topic_1 bytes32 messageId, bytes encodedData)` 也即 0x520d2afde79cbd5db58755ac9480f81bc658e5c517fcae7365a3d832590b0183
// 获取 encodedData 就是用户所要签署的 message
//
//	curl 'https://rpc.gnosischain.com/' \
//	  -H 'accept: */*' \
//	  -H 'accept-language: en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7' \
//	  -H 'cache-control: no-cache' \
//	  -H 'content-type: application/json' \
//	  -H 'origin: https://bridge.gnosischain.com' \
//	  -H 'pragma: no-cache' \
//	  -H 'priority: u=1, i' \
//	  -H 'referer: https://bridge.gnosischain.com/' \
//	  -H 'sec-ch-ua: "Not;A=Brand";v="99", "Google Chrome";v="139", "Chromium";v="139"' \
//	  -H 'sec-ch-ua-mobile: ?0' \
//	  -H 'sec-ch-ua-platform: "macOS"' \
//	  -H 'sec-fetch-dest: empty' \
//	  -H 'sec-fetch-mode: cors' \
//	  -H 'sec-fetch-site: same-site' \
//	  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36' \
//	  --data-raw '{"method":"eth_getTransactionReceipt","params":["0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b"],"id":43,"jsonrpc":"2.0"}'
//
// {"jsonrpc":"2.0","result":{"transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","transactionIndex":"0x5","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","cumulativeGasUsed":"0x1a30d8","gasUsed":"0x56f54","effectiveGasPrice":"0xca","from":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","to":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","contractAddress":null,"logs":[{"removed":false,"logIndex":"0x19","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","data":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x0000000000000000000000009003d8731df107aa5e3feaddfc165787b910ff1e","0x000000000000000000000000f6a78083ca3e2a662d6dd1703c939c8ace2e268d"]},{"removed":false,"logIndex":"0x1a","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","data":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000149003d8731df107aa5e3feaddfc165787b910ff1e","topics":["0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16","0x0000000000000000000000009003d8731df107aa5e3feaddfc165787b910ff1e","0x000000000000000000000000f6a78083ca3e2a662d6dd1703c939c8ace2e268d"]},{"removed":false,"logIndex":"0x1b","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","data":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000","topics":["0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5","0x000000000000000000000000f6a78083ca3e2a662d6dd1703c939c8ace2e268d"]},{"removed":false,"logIndex":"0x1c","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0x177127622c4a00f3d409b75571e12cb3c8973d3c","data":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000","topics":["0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef","0x000000000000000000000000f6a78083ca3e2a662d6dd1703c939c8ace2e268d","0x0000000000000000000000000000000000000000000000000000000000000000"]},{"removed":false,"logIndex":"0x1d","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0x75df5af045d91108662d8080fd1fefad6aa0bb59","data":"0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000b500050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57f6a78083ca3e2a662d6dd1703c939c8ace2e268d88ad09518695c6c3712ac10a214be5109a655671000927c00101806401272255bb000000000000000000000000def1ca1fb7fbcdc777520aa7f396b4e015f497ab0000000000000000000000009003d8731df107aa5e3feaddfc165787b910ff1e0000000000000000000000000000000000000000000000000de0b6b3a764000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000a3bc83d557e3f2ddff4d44966a96397760159d8b00000000000000000000000000000000000000000000000000000000000000010000000000000000000000007606e9d8655e48159e7bec8541c2e71a7aa3e418","topics":["0x218247aabc759e65b5bb92ccc074f9d62cd187259f2a0984c3c9cf91f67ff7cf","0x3c8ff829064102a66faf249012a71c95e21641e68009aa5fcd9e8308aaff8453"]},{"removed":false,"logIndex":"0x1f","transactionIndex":"0x5","transactionHash":"0xbf8c7779887487872ec2314a2f599564fb3a4259b58f516233aa34dd2a92970b","blockHash":"0xf9b167c2c6aaecdb5ca47f34fd96fc2150903724cb6e26d6e6a58540b6ef35c3","blockNumber":"0x27c43ad","address":"0xf6a78083ca3e2a662d6dd1703c939c8ace2e268d","data":"0x0000000000000000000000000000000000000000000000000de0b6b3a7640000","topics":["0x59a9a8027b9c87b961e254899821c9a276b5efc35d1f7409ea4f291470f1629a","0x000000000000000000000000177127622c4a00f3d409b75571e12cb3c8973d3c","0x0000000000000000000000009003d8731df107aa5e3feaddfc165787b910ff1e","0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001ad57"]}],"logsBloom":"0x0002000000000000000200000000000000000000000000100000000002200000000000000000000000000000000000000000000000000000000000000000000000002000000000000800000800400000000000000010000000002000000000201000081002000000008000000000080018000000000000000000001000000000000400002000000002000010000000800000002000001000808000000020000100000000000000000001000001000000000000000000000000000000000000000000000600000020000000000000000000200000000000282004000000013000000000000010000000000000000000000000000020000008000000000000000","status":"0x1","type":"0x2"},"id":43}
// 第三步， 通过 read 合约 AMBBridgeHelper,地址0x7d94ece17e81355326e3359115D4B02411825EdD, 的 getSignatures(bytes message) 获取签名
// 第四步， 通过 gnosis_cliam package 中 safeExecuteSignaturesWithAutoGasLimit 对 message 和 signagture 做 abi 编码
func getClaim(ctx context.Context, withdrawTx, trader string) ([]byte, error) {
	const (
		// GraphQL 端点
		graphqlEndpoint = "https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ"
		// AMBBridgeHelper 合约地址
		ambBridgeHelperAddress = "0x7d94ece17e81355326e3359115D4B02411825EdD"
		// UserRequestForSignature 事件 topic
		userRequestForSignatureTopic = "0x520d2afde79cbd5db58755ac9480f81bc658e5c517fcae7365a3d832590b0183"
	)

	// 第一步：查询 API，获取交易状态为 "UNCLAIMED" 的交易
	query := `
		query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {
			transactions(
				where: $where
				orderBy: $orderBy
				orderDirection: $orderDirection
				first: $first
				skip: $skip
			) {
				id
				bridgeName
				transactionHash
				initiator
				initiatorAmount
				initiatorNetwork
				initiatorToken
				receiver
				receiverToken
				receiverAmount
				receiverNetwork
				transactionStatus
				timestamp
				execution {
					id
					timestamp
					transactionHash
					validatorAddr
				}
				validations {
					id
					timestamp
					transactionHash
					validatorAddr
				}
			}
		}`

	variables := map[string]interface{}{
		"where": map[string]interface{}{
			"transactionHash": withdrawTx,
		},
		"skip":  0,
		"first": 1000,
	}

	req := GraphQLRequest{
		Query:     query,
		Variables: variables,
	}

	resp, err := makeGraphQLRequest(ctx, graphqlEndpoint, req, "")
	if err != nil {
		return nil, errors.Wrap(err, "failed to query GraphQL API")
	}

	if len(resp.Data.Transactions) == 0 {
		return nil, errors.Errorf("no transaction found for hash: %s", withdrawTx)
	}

	tx := resp.Data.Transactions[0]
	if tx.TransactionStatus != "UNCLAIMED" {
		return nil, errors.Errorf("transaction status is not UNCLAIMED: %s", tx.TransactionStatus)
	}

	// 第二步：通过 RPC 查询交易收据，获取 UserRequestForSignature 事件的 encodedData
	encodedData, err := getEncodedDataFromReceipt(ctx, withdrawTx, userRequestForSignatureTopic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get encoded data from receipt")
	}

	// 第三步：调用 AMBBridgeHelper 合约的 getSignatures 方法获取签名
	signatures, err := getSignaturesFromContract(ctx, ambBridgeHelperAddress, encodedData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get signatures from contract")
	}

	// 第四步：使用 gnosis_claim 包中的 safeExecuteSignaturesWithAutoGasLimit 对 message 和 signature 做 ABI 编码
	claimData, err := packSafeExecuteSignaturesWithAutoGasLimit(encodedData, signatures)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack safe execute signatures")
	}

	return claimData, nil
}

// getEncodedDataFromReceipt 从交易收据中获取 UserRequestForSignature 事件的 encodedData
func getEncodedDataFromReceipt(ctx context.Context, txHash, topic string) ([]byte, error) {
	// 这里需要实现 RPC 调用来获取交易收据
	// 由于没有直接的 ethclient，我们需要通过 HTTP RPC 调用来实现

	// 构建 RPC 请求
	rpcReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getTransactionReceipt",
		"params":  []string{txHash},
		"id":      1,
	}

	// 发送 RPC 请求到 Gnosis RPC 端点
	rpcResp, err := makeRPCRequest(ctx, "https://rpc.gnosischain.com/", rpcReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction receipt")
	}

	// 解析响应，查找 UserRequestForSignature 事件的 encodedData
	encodedData, err := parseEncodedDataFromReceipt(rpcResp, topic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse encoded data from receipt")
	}

	argType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, errors.Wrap(err, "new abi type")
	}
	message, err := abi.Arguments{{Type: argType}}.Unpack(encodedData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack signatures")
	}

	return message[0].([]byte), nil
}

// getSignaturesFromContract 调用 AMBBridgeHelper 合约的 getSignatures 方法
func getSignaturesFromContract(ctx context.Context, contractAddress string, message []byte) ([]byte, error) {
	// 这里需要实现合约调用来获取签名
	// 由于没有直接的 ethclient，我们需要通过 HTTP RPC 调用来实现

	data := "0x" + common.Bytes2Hex(packGetSignaturesCall(message))

	// 构建 RPC 请求调用合约的 getSignatures 方法
	rpcReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_call",
		"params": []interface{}{
			map[string]interface{}{
				"to":   contractAddress,
				"data": data,
			},
			"latest",
		},
		"id": 1,
	}

	// 发送 RPC 请求
	rpcResp, err := makeRPCRequest(ctx, "https://rpc.gnosischain.com/", rpcReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call contract")
	}

	// 解析响应获取签名
	signatures, err := parseSignaturesFromResponse(rpcResp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse signatures from response")
	}

	return signatures, nil
}

// packGetSignaturesCall 打包调用 getSignatures 方法的调用数据
func packGetSignaturesCall(message []byte) []byte {
	// 方法签名：getSignatures(bytes message)
	// 方法 ID：keccak256("getSignatures(bytes)") 的前4字节
	// 计算得到的方法 ID：0x6b45b4e3

	// 构建 ABI 编码
	// 对于 bytes 类型，需要先编码长度，然后编码数据
	// 格式：[4字节方法ID][32字节偏移量][32字节长度][实际数据]

	// 方法 ID
	methodID := []byte{0x6b, 0x45, 0xb4, 0xe3}

	// 偏移量（指向数据开始位置，固定为 0x20）
	offset := make([]byte, 32)
	offset[31] = 0x20

	// 数据长度
	length := make([]byte, 32)
	length[31] = byte(len(message))

	// 数据（需要填充到32字节的倍数）
	paddedData := make([]byte, 0)
	paddedData = append(paddedData, length...)
	paddedData = append(paddedData, message...)

	// 如果数据长度不是32的倍数，需要填充
	if len(paddedData)%32 != 0 {
		padding := 32 - (len(paddedData) % 32)
		for i := 0; i < padding; i++ {
			paddedData = append(paddedData, 0)
		}
	}

	// 组合所有部分
	result := make([]byte, 0)
	result = append(result, methodID...)
	result = append(result, offset...)
	result = append(result, paddedData...)

	return result
}

// parseSignaturesFromResponse 从 RPC 响应中解析签名
func parseSignaturesFromResponse(response map[string]interface{}) ([]byte, error) {
	// 解析 RPC 响应获取签名
	result, ok := response["result"]
	if !ok {
		return nil, errors.New("no result in response")
	}

	resultStr, ok := result.(string)
	if !ok {
		return nil, errors.New("invalid result format")
	}

	// 移除 0x 前缀
	if strings.HasPrefix(resultStr, "0x") {
		resultStr = resultStr[2:]
	}

	// 转换为字节
	signaturesBytes := common.FromHex(resultStr)
	if len(signaturesBytes) == 0 {
		return nil, errors.New("empty signatures")
	}

	argType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, errors.Wrap(err, "new abi type")
	}
	signatures, err := abi.Arguments{{Type: argType}}.Unpack(signaturesBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack signatures")
	}

	return signatures[0].([]byte), nil
}

// packSafeExecuteSignaturesWithAutoGasLimit 使用 gnosis_claim 包打包调用数据
func packSafeExecuteSignaturesWithAutoGasLimit(message, signatures []byte) ([]byte, error) {
	// 导入 gnosis_claim 包
	gnosisClaimAbi, err := gnosis_claim.GnosisClaimMetaData.ParseABI()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get gnosis claim ABI")
	}

	// 打包 safeExecuteSignaturesWithAutoGasLimit 调用
	packedData, err := gnosisClaimAbi.Pack("safeExecuteSignaturesWithAutoGasLimit", message, signatures)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack safeExecuteSignaturesWithAutoGasLimit")
	}

	return packedData, nil
}

// 辅助函数

// makeRPCRequest 发送 RPC 请求
func makeRPCRequest(ctx context.Context, endpoint string, req map[string]interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal RPC request")
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", endpoint, strings.NewReader(string(jsonData)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create RPC request")
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send RPC request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("unexpected RPC status code: %d", resp.StatusCode)
	}

	var rpcResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, errors.Wrap(err, "failed to decode RPC response")
	}

	return rpcResp, nil
}

// parseEncodedDataFromReceipt 从交易收据中解析 encodedData
func parseEncodedDataFromReceipt(receipt map[string]interface{}, topic string) ([]byte, error) {
	// 解析交易收据，查找指定 topic 的日志
	result, ok := receipt["result"]
	if !ok {
		return nil, errors.New("no result in receipt")
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid result format")
	}

	logs, ok := resultMap["logs"]
	if !ok {
		return nil, errors.New("no logs in receipt")
	}

	logsArray, ok := logs.([]interface{})
	if !ok {
		return nil, errors.New("invalid logs format")
	}

	// 查找指定 topic 的日志
	for _, log := range logsArray {
		logMap, ok := log.(map[string]interface{})
		if !ok {
			continue
		}

		topics, ok := logMap["topics"]
		if !ok {
			continue
		}

		topicsArray, ok := topics.([]interface{})
		if !ok || len(topicsArray) == 0 {
			continue
		}

		// 检查第一个 topic 是否匹配
		if topicsArray[0] == topic {
			// 找到匹配的日志，提取 data 字段
			data, ok := logMap["data"]
			if !ok {
				continue
			}

			dataStr, ok := data.(string)
			if !ok {
				continue
			}

			// 移除 0x 前缀并转换为字节
			if strings.HasPrefix(dataStr, "0x") {
				dataStr = dataStr[2:]
			}

			encodedData := common.FromHex(dataStr)
			if len(encodedData) == 0 {
				return nil, errors.New("empty encoded data")
			}

			return encodedData, nil
		}
	}

	return nil, errors.New("UserRequestForSignature event not found in receipt")
}

func WaitForClaim(ctx context.Context, withdrawTx, trader string) ([]byte, error) {
	data, err := getClaim(ctx, withdrawTx, trader)
	if err == nil {
		return data, nil
	}

	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			data, err = getClaim(ctx, withdrawTx, trader)
			if err != nil {
				fmt.Println("getClaim error:", err)
				continue
			}
			return data, nil
		}
	}
}
