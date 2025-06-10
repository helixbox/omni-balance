package enclave

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"omni-balance/utils/erc20"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func TestClient_SignErc20Transfer(t *testing.T) {
	// 设置测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法
		assert.Equal(t, "POST", r.Method)

		// 验证请求路径
		assert.Equal(t, "/sign/42161", r.URL.Path)

		// 验证Content-Type
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

		// 解码请求体
		var req SignRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		assert.NoError(t, err)

		// 返回测试响应
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignResponse{
			Signature: "d91bda9e6bc3195a7cc53a30ebe7cce9a96e1ad91e6b4559d72600afae3e48b911e9788938768bc3465c6c5030a0f03dda4a9c533985164bf78b4ac2665273d21b",
		})
	}))
	defer server.Close()

	// 创建客户端
	client := NewClient(server.URL)

	token := common.HexToAddress("0xaf88d065e77c8cc2239327c5edb3a432268e5831")
	receiver := common.HexToAddress("0x3304791b5034c82790167dbebdd7ca1bfc8c9dcf")
	amount := big.NewInt(100000)

	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	assert.NoError(t, err)
	input, err := erc20Abi.Pack("transfer", receiver, amount)
	assert.NoError(t, err)
	assert.NoError(t, err)

	// 创建待签名交易
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(42161),
		Nonce:     8,
		GasTipCap: big.NewInt(1350000),
		GasFeeCap: big.NewInt(13500000),
		Gas:       200000,
		To:        &token,
		Value:     big.NewInt(0),
		Data:      input,
	})

	// 执行签名
	signedTx, err := client.SignErc20Transfer(tx, 42161)
	encoded, err := signedTx.MarshalBinary()

	// 验证结果
	assert.NoError(t, err)
	assert.NotNil(t, signedTx)
	assert.Equal(t, signedTx.Hash().Hex(), "0xbf9d11d478eaba72fd13850d7c573232e17e419fa8e5ef1124aea961a4c521a5")
	assert.Equal(t, encoded, common.FromHex("02f8b082a4b1088314997083cdfe6083030d4094af88d065e77c8cc2239327c5edb3a432268e583180b844a9059cbb0000000000000000000000003304791b5034c82790167dbebdd7ca1bfc8c9dcf00000000000000000000000000000000000000000000000000000000000186a0c080a0d91bda9e6bc3195a7cc53a30ebe7cce9a96e1ad91e6b4559d72600afae3e48b9a011e9788938768bc3465c6c5030a0f03dda4a9c533985164bf78b4ac2665273d2"))
}
