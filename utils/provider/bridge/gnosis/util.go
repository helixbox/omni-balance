package gnosis

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"omni-balance/utils/chains"
	gnosis_deposit "omni-balance/utils/enclave/router/gnosis/deposit"
	"omni-balance/utils/erc20"
	"omni-balance/utils/wallets"
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
	return erc1363Abi.Pack("transferAndCall", to, value.BigInt(), data)
}

// 第一步请求 API ，通过 depositTxHash == transactionHash 获取对应的 id
//
//	curl 'https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ' \
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
//	  --data-raw '{"query":"query Transactions($where: Transaction_filter, $orderBy: Transaction_orderBy, $orderDirection: OrderDirection, $first: Int, $skip: Int) {\n  transactions(\n    where: $where\n    orderBy: $orderBy\n    orderDirection: $orderDirection\n    first: $first\n    skip: $skip\n  ) {\n    ...TransactionFragment\n  }\n}\n\nfragment TransactionFragment on Transaction {\n  id\n  bridgeName\n  transactionHash\n  initiator\n  initiatorAmount\n  initiatorNetwork\n  initiatorToken\n  receiver\n  receiverToken\n  receiverAmount\n  receiverNetwork\n  transactionStatus\n  timestamp\n  execution {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n  validations {\n    id\n    timestamp\n    transactionHash\n    validatorAddr\n  }\n}","variables":{"where":{"id_in":["0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61"]},"skip":0,"first":1000},"operationName":"Transactions"}'
//
// {"data":{"transactions":[{"bridgeName":"AMB","execution":{"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61-0xf6a78083ca3e2a662d6dd1703c939c8ace2e268d","timestamp":"1755152207","transactionHash":"0xe3007d614ff5fdb88592f55e7ed60bf6a2d00010aaadac01bed4a26092377e26","validatorAddr":null},"id":"0x00050000a7823d6f1e31569f51861e345b30c6bebf70ebe7000000000001aa61","initiator":null,"initiatorAmount":null,"initiatorNetwork":"gnosis","initiatorToken":null,"receiver":"0x9003d8731df107aa5e3feaddfc165787b910ff1e","receiverAmount":"1000000000000000000","receiverNetwork":"mainnet","receiverToken":"0xdef1ca1fb7fbcdc777520aa7f396b4e015f497ab","timestamp":null,"transactionHash":null,"transactionStatus":"COMPLETED","validations":[]}]}}
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
		graphqlEndpoint = "https://gateway-arbitrum.network.thegraph.com/api/a077e395ce795c9f990504ec1ec6c8ba/subgraphs/id/2ths6FTZhCBggnyakh7PL5KH91zjRv8xPNfzaCRKogJ"
		// 查询间隔
		pollInterval = 10 * time.Second
		// 最大等待时间
		maxWaitTime = 10 * time.Minute
	)

	// 第一步：通过 depositTxHash 获取交易 ID
	transactionID, err := getTransactionID(ctx, graphqlEndpoint, depositTxHash, customCookies)
	if err != nil {
		return "", errors.Wrap(err, "failed to get transaction ID")
	}

	// 第二步：等待并获取子交易哈希
	childTxHash, err := waitForChildTransaction(ctx, graphqlEndpoint, transactionID, pollInterval, maxWaitTime, customCookies)
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
			"id_in": []string{transactionID},
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
		defaultCookies := []string{
			"__Secure-next-auth.callback-url=https://bridge.gnosischain.com",
			"__Secure-next-auth.csrf-token=random_token_here",
			"__Host-next-auth.csrf-token=random_token_here",
		}
		httpReq.Header.Set("Cookie", strings.Join(defaultCookies, "; "))
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

func WaitForClaim(ctx context.Context, depositTxHash, trader string) ([]byte, error) {
	return nil, nil
}
