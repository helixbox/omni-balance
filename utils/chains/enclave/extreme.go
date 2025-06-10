package enclave_exreme

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// SignRequest 定义签名请求的结构
type SignRequest struct {
	Erc20Transfer Erc20Transfer `json:"erc20Transfer"`
}

type Erc20Transfer struct {
	Token    common.Address `json:"token"`
	Receiver common.Address `json:"receiver"`
	Amount   *big.Int       `json:"acmount"`
	Meta     struct {
		Nonce                uint64   `json:"nonce"`
		GasLimit             uint64   `json:"gasLimit"`
		MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
		MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	} `json:"meta"`
}

// SignResponse 定义签名服务响应的结构
type SignResponse struct {
	// TODO: 根据实际签名服务的响应格式定义
	Signature string `json:"signature"`
}

// Client 签名服务客户端
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient 创建新的签名服务客户端
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SignTransaction 请求签名服务对交易进行签名
func (c *Client) SignErc20Transfer(tx *types.Transaction, transfer *Erc20Transfer, chainID int64) (*types.Transaction, error) {
	// 构造请求体
	req := SignRequest{
		Erc20Transfer: *transfer,
	}

	// 将请求体转换为JSON
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 构建URL
	url := fmt.Sprintf("%s/sign/%d", c.baseURL, chainID)

	// 创建HTTP请求
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("signing service returned status %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var signResp SignResponse
	if err := json.Unmarshal(body, &signResp); err != nil {
		return nil, errors.WithStack(err)
	}

	signer := types.NewLondonSigner(big.NewInt(chainID))
	// Decode hex signature string to bytes
	sig := common.FromHex(signResp.Signature)

	adjustSignature(sig)
	return tx.WithSignature(signer, sig)
}

func adjustSignature(sig []byte) {
	sig[64] = sig[64] - 27
}
