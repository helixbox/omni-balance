package enclave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"time"

	"omni-balance/utils/erc20"

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
	Meta     Meta           `json:"meta"`
}

type Meta struct {
	Nonce                uint64   `json:"nonce"`
	GasLimit             uint64   `json:"gasLimit"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
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
func (c *Client) SignErc20Transfer(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	if tx.Type() != types.DynamicFeeTxType {
		return nil, errors.New("only support dynamic fee tx")
	}

	receiver, amount, err := GetTransferInfo(tx.Data())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	transfer := &Erc20Transfer{
		Token:    *tx.To(),
		Receiver: receiver,
		Amount:   amount,
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

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

// GetTransferInfo input ex: a9059cbb0000000000000000000000000350101f2cb6aa65caab7954246a56f906a3f57d0000000000000000000000000000000000000000033b2e3c9fd0803ce8000000
func GetTransferInfo(input []byte) (to common.Address, amount *big.Int, err error) {
	if len(input) < 4 {
		return common.Address{}, nil, errors.New("invalid input")
	}
	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get abi")
	}
	args, err := erc20Abi.Methods["transfer"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "unpack")
	}
	if len(args) != 2 {
		return common.Address{}, nil, errors.New("invalid args")
	}
	return args[0].(common.Address), args[1].(*big.Int), nil
}
