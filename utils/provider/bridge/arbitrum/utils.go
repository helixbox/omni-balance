package arbitrum

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/wallets"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const GATEWAY_ROUTER_ABI = `[{"inputs":[{"internalType":"address","name":"_l1Token","type":"address"},{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"outboundTransfer","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"nonpayable","type":"function"}]`

type TxRequest struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Data  string `json:"data"`
}

type WaitResponse struct {
	Tx       string `json:"tx"`
	Complete bool   `json:"complete"`
}

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

func Deposit(ctx context.Context, l1Address, receiver common.Address, amount decimal.Decimal) (TxRequest, error) {
	u, err := url.Parse("http://common-rebalance/rebalance/arb1-erc20-deposit")
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "url parse error")
	}
	type DepositRequest struct {
		L1Token string `json:"l1Token"`
		Amount  string `json:"amount"`
	}
	dr := DepositRequest{
		L1Token: l1Address.Hex(),
		Amount:  amount.BigInt().String(),
	}
	body, err := json.Marshal(dr)
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "marshal deposit request error")
	}

	var response TxRequest
	err = utils.Request(ctx, "POST", u.String(), bytes.NewReader(body), &response)
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "get deposit tx request error")
	}

	if !strings.EqualFold(receiver.Hex(), response.From) {
		return TxRequest{}, errors.New("receiver address mismatch")
	}
	return response, nil
}

func WaitForChildTransactionReceipt(ctx context.Context, txHash string) (WaitResponse, error) {
	u, err := url.Parse("http://common-rebalance/rebalance/wait-l2-tx-receipt")
	if err != nil {
		return WaitResponse{}, errors.Wrap(err, "url parse error")
	}
	type waitRequest struct {
		Tx string `json:"tx"`
	}
	dr := waitRequest{
		Tx: txHash,
	}
	body, err := json.Marshal(dr)
	if err != nil {
		return WaitResponse{}, errors.Wrap(err, "marshal wait request error")
	}

	var result WaitResponse
	err = utils.Request(ctx, "POST", u.String(), bytes.NewReader(body), &result)
	if err != nil {
		return WaitResponse{}, errors.Wrap(err, "get deposit tx request error")
	}

	return result, nil
}

type ClaimRequest struct{}

func WaitForClaim(ctx context.Context, txHash string) (TxRequest, error) {
	u, err := url.Parse("http://common-rebalance/rebalance/arb1-erc20-claim")
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "url parse error")
	}
	type waitRequest struct {
		Tx string `json:"tx"`
	}
	dr := waitRequest{
		Tx: txHash,
	}
	body, err := json.Marshal(dr)
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "marshal wait request error")
	}

	var result TxRequest
	err = utils.Request(ctx, "POST", u.String(), bytes.NewReader(body), &result)
	if err != nil {
		return TxRequest{}, errors.Wrap(err, "get claim tx request error")
	}

	return result, nil
}

func Withdraw(ctx context.Context, l1Address, receiver common.Address, amount decimal.Decimal) ([]byte, error) {
	contractAbi, err := abi.JSON(strings.NewReader(GATEWAY_ROUTER_ABI))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("outboundTransfer", l1Address, receiver, amount.BigInt(), "")
}
