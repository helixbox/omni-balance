package enclave

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

type RequestType interface {
	GetRequestType() string
}

type Meta struct {
	Nonce                uint64   `json:"nonce"`
	GasLimit             uint64   `json:"gasLimit"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
	Value                *big.Int `json:"value,omitempty"`
}

type SignResponse struct {
	Signature string `json:"result"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) signRequest(request RequestType, tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	if tx.Type() != types.DynamicFeeTxType {
		return nil, errors.New("only support dynamic fee tx")
	}

	// 将请求体转换为JSON
	jsonBody, err := json.Marshal(request)
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

	if signResp.Signature == "" {
		return nil, errors.New("enclave not approved")
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

