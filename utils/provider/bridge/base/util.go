package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"omni-balance/utils/chains"
	base_deposit "omni-balance/utils/enclave/router/base/deposit"
	base_withdraw "omni-balance/utils/enclave/router/base/withdraw"
	"omni-balance/utils/wallets"

	log "omni-balance/utils/logging"

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
	client := &http.Client{}

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
	withdrawer, err := Withdrawer(common.HexToHash(withdrawTx))
	if err != nil {
		return "", errors.Wrap(err, "init withdrawer")
	}

	err = withdrawer.CheckIfProvable()
	if err != nil {
		return "", errors.Wrap(err, "check provable")
	}

	proofTime, err := withdrawer.GetProvenWithdrawalTime()
	if err != nil {
		return "", errors.Wrap(err, "Error querying withdrawal proof")
	}

	if proofTime == 0 {
		withdrawer.ProveWithdrawal()
		// if err != nil {
		// 	return "", errors.Wrap(err, "Error proving withdrawal")
		// }

		fmt.Println("The withdrawal has been successfully proven, finalization of the withdrawal can be done once the dispute game has finished and the finalization period has elapsed")
		return "", nil
	}
	return "", nil
}

func getClaim(ctx context.Context, proveTx, trader string) (string, error) {
	return getId(ctx, proveTx, trader, 5)
}

func getId(ctx context.Context, txHash, trader string, status uint32) (string, error) {
	client := &http.Client{}

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
