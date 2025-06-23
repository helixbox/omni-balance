package base

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"omni-balance/utils/chains"
	"omni-balance/utils/wallets"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
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

// func Deposit(ctx context.Context, l1Address, l2Address, receiver common.Address, amount decimal.Decimal) ([]byte, error) {
// 	routerAbi, err := deposit.depositMetaData.GetAbi()
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
//
// 	return routerAbi.Pack("depositERC20To", l1Address, l2Address, receiver, big.NewInt(200000), amount.BigInt(), []byte{})
// }

// ~ curl 'https://api.superbridge.app/api/v5/bridge/activity' \
//
//	        -H 'content-type: application/json' \
//	        -H 'origin: https://superbridge.app' \
//	        --data-raw '{"evmAddress":"0x9003d8731df107aA5E3FEADdFC165787b910Ff1e","deploymentIds":["81883861-df09-4a49-816e-7268435d27eb"]}' | jq
//
//		{
//		  "total": 2,
//		  "transactions": [
//		    {
//		      "id": "a536a539-9f9c-44ab-a401-95fccddd5832",
//		      "createdAt": "2025-06-12T06:05:03.836Z",
//		      "updatedAt": "2025-06-12T06:08:05.115Z",
//		      "type": "deposit",
//		      "provider": "OptimismDeposit",
//		      "deploymentId": "81883861-df09-4a49-816e-7268435d27eb",
//		      "status": 6,
//		      "metadata": {
//		        "to": "0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33",
//		        "data": {
//		          "amount": "24999000000000000000000",
//		          "l1TokenAddress": "0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB",
//		          "l2TokenAddress": "0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69"
//		        },
//		        "from": "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e",
//		        "type": "token-deposit"
//		      },
//		      "relay": {
//		        "transactionHash": "0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db",
//		        "timestamp": 1,
//		        "status": "confirmed"
//		      },
//		      "receive": {
//		        "transactionHash": "0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db",
//		        "timestamp": 1,
//		        "status": "confirmed"
//		      },
//		      "deposit": {
//		        "blockNumber": 22686578,
//		        "transactionHash": "0x6b1115f3ba8f76b42e483d1c0eb1300f18f750dd5404644c6dd20bd7b21771f6",
//		        "timestamp": 1749708299000,
//		        "status": "confirmed"
//		      },
//		      "send": {
//		        "blockNumber": 22686578,
//		        "transactionHash": "0x6b1115f3ba8f76b42e483d1c0eb1300f18f750dd5404644c6dd20bd7b21771f6",
//		        "timestamp": 1749708299000,
//		        "status": "confirmed"
//		      },
//		      "l2TransactionHash": "0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db",
//		      "toChainId": 8453,
//		      "fromChainId": 1,
//		      "from": "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e",
//		      "to": "0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33",
//		      "amount": "24999000000000000000000",
//		      "token": "0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB",
//		      "receiveToken": "0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69"
//		    },
//		    {
//		      "id": "050e500e-8b66-415e-ae6a-db6868c1b741",
//		      "createdAt": "2025-06-12T05:54:32.170Z",
//		      "updatedAt": "2025-06-12T05:57:32.437Z",
//		      "type": "deposit",
//		      "provider": "OptimismDeposit",
//		      "deploymentId": "81883861-df09-4a49-816e-7268435d27eb",
//		      "status": 6,
//		      "metadata": {
//		        "to": "0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33",
//		        "data": {
//		          "amount": "1000000000000000000",
//		          "l1TokenAddress": "0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB",
//		          "l2TokenAddress": "0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69"
//		        },
//		        "from": "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e",
//		        "type": "token-deposit"
//		      },
//		      "relay": {
//		        "transactionHash": "0xa04e242ad0e8ea0a74fa590dcf0e7a2213c6ee20a23adbd6865661b2ba5c601f",
//		        "timestamp": 1,
//		        "status": "confirmed"
//		      },
//		      "receive": {
//		        "transactionHash": "0xa04e242ad0e8ea0a74fa590dcf0e7a2213c6ee20a23adbd6865661b2ba5c601f",
//		        "timestamp": 1,
//		        "status": "confirmed"
//		      },
//		      "deposit": {
//		        "blockNumber": 22686525,
//		        "transactionHash": "0xf8f34f4ad0f934be684e1a0d1cda6c12c5af28c23b1c737f0bc4f1d121bf5c15",
//		        "timestamp": 1749707663000,
//		        "status": "confirmed"
//		      },
//		      "send": {
//		        "blockNumber": 22686525,
//		        "transactionHash": "0xf8f34f4ad0f934be684e1a0d1cda6c12c5af28c23b1c737f0bc4f1d121bf5c15",
//		        "timestamp": 1749707663000,
//		        "status": "confirmed"
//		      },
//		      "l2TransactionHash": "0xa04e242ad0e8ea0a74fa590dcf0e7a2213c6ee20a23adbd6865661b2ba5c601f",
//		      "toChainId": 8453,
//		      "fromChainId": 1,
//		      "from": "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e",
//		      "to": "0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33",
//		      "amount": "1000000000000000000",
//		      "token": "0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB",
//		      "receiveToken": "0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69"
//		    }
//		  ],
//		  "actionRequiredCount": 0,
//		  "inProgressCount": 0,
//		  "hasWithdrawalReadyToFinalize": null,
//		  "recipients": [
//		    "8453:0xD1Fc331dBF956e21DA5c2D89CAA2f98c80317D33"
//		  ]
//		}
type ActivityRequest struct {
	ID            string   `json:"id,omitempty"`
	EvmAddress    string   `json:"evmAddress"`
	DeploymentIds []string `json:"deploymentIds"`
}

type ActivityResponse struct {
	Transactions []struct {
		Deposit struct {
			TransactionHash string `json:"transactionHash"`
		} `json:"deposit"`
		Receive struct {
			TransactionHash string `json:"transactionHash"`
		} `json:"receive"`
		// 其他字段...
	} `json:"transactions"`
}

func WaitForChildTransactionReceipt(ctx context.Context, depositTxHash, trader string) (string, error) {
	apiUrl := "https://api.superbridge.app/api/v5/bridge/activity"
	requestBody := map[string]interface{}{
		"evmAddress":    trader,
		"deploymentIds": []string{"81883861-df09-4a49-816e-7268435d27eb"},
	}
	body, _ := json.Marshal(requestBody)

	log.Infof("request: %s", string(body))

	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://superbridge.app")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var activity struct {
		Transactions []struct {
			Deposit struct {
				TransactionHash string `json:"transactionHash"`
			} `json:"deposit"`
			Receive struct {
				TransactionHash string `json:"transactionHash"`
			} `json:"receive"`
		} `json:"transactions"`
	}
	if err := json.Unmarshal(respBody, &activity); err != nil {
		return "", err
	}
	log.Infof("response: %s", respBody)

	for _, tx := range activity.Transactions {
		if tx.Deposit.TransactionHash == depositTxHash {
			return tx.Receive.TransactionHash, nil
		}
	}
	return "", errors.New("receive transaction hash not found after polling")
}
