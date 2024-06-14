package darwinia

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallets"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GetSourceChains returns the list of source chain IDs that support transferring the specified token to the target chain.
// It checks the SUPPORTS map to identify the supported source chains based on the target chain ID and token name.
// The source chain IDs are returned as a slice of int64.
//
// Parameters:
//   - targetChainId: The ID of the target chain.
//   - TokenName: The name of the token.
//   - wallet: The wallet address.
//   - amount: The amount of the token to be transferred.
func (b *Bridge) FindValidSourceChains(ctx context.Context, targetChainId int, TokenName, wallet string,
	amount decimal.Decimal) (sourceChainIds []int64) {

	sourceChains := GetSourceChains(int64(targetChainId), TokenName)
	if len(sourceChains) == 0 {
		return nil
	}
	var (
		w sync.WaitGroup
		m sync.Mutex
	)
	for _, sourceChain := range sourceChains {
		w.Add(1)
		go b.checkAndAppendSourceChainIfBalanceSufficient(ctx, sourceChain, TokenName, wallet, amount, &sourceChainIds, &w, &m)
	}
	w.Wait()
	return sourceChainIds
}

func (b *Bridge) GetValidSourceChain(ctx context.Context, targetChainId int, TokenName, wallet string,
	amount decimal.Decimal) (sourceChainId int64) {

	sourceChains := b.FindValidSourceChains(ctx, targetChainId, TokenName, wallet, amount)
	if len(sourceChains) == 0 {
		return 0
	}
	return utils.Choose(sourceChains)
}

func (b *Bridge) GetBalance(ctx context.Context, args provider.BalanceParams) (decimal.Decimal, error) {
	chain := b.config.GetChainConfig(args.Chain)
	token := chain.GetToken(args.Token)
	if token.ContractAddress == "" {
		return decimal.Zero, errors.Errorf("token %s not found on %s", args.Token, args.Chain)
	}
	client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "dial rpc")
	}
	defer client.Close()

	return chains.GetTokenBalance(ctx, client, token.ContractAddress, args.Wallet, token.Decimals)
}

func (b *Bridge) checkAndAppendSourceChainIfBalanceSufficient(ctx context.Context, sourceChain int64,
	TokenName, wallet string, amount decimal.Decimal, sourceChainIds *[]int64, w *sync.WaitGroup, m *sync.Mutex) {

	defer w.Done()
	balance, err := b.GetBalance(ctx,
		provider.BalanceParams{Token: TokenName, Chain: constant.GetChainName(sourceChain), Wallet: wallet})
	if err != nil {
		logrus.Debugf("get %s balance error: %s", constant.GetChainName(sourceChain), err)
		return
	}
	if !balance.GreaterThanOrEqual(amount) {
		return
	}
	logrus.Debugf("check %s balance: %s >= %s", constant.GetChainName(sourceChain), balance, amount)
	m.Lock()
	defer m.Unlock()
	*sourceChainIds = append(*sourceChainIds, sourceChain)
}

func (b *Bridge) WaitForBridgeSuccess(ctx context.Context, tx, sender string) (HistoryRecord, error) {
	log := logrus.WithFields(logrus.Fields{
		"tx":      tx,
		"sender":  sender,
		"chain":   b.Name(),
		"type":    b.Type(),
		"message": "wait for bridge success",
	})
	requestData := GraphQLRequest{
		OperationName: "GetHistory",
		Variables: map[string]interface{}{
			"row":    20,
			"page":   0,
			"sender": sender,
		},
		Query: graphOlQuery,
	}
	jsonBytes, err := json.Marshal(requestData)
	if err != nil {

		return HistoryRecord{}, errors.Wrap(err, "marshal request error")
	}
	var t = time.NewTicker(time.Second)
	defer t.Stop()
	getRecords := func() (HistoryRecordResponse, error) {
		resp, err := http.Post("https://apollo.xtoken.box/graphql", "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			log.Warnf("get %s status error: %s", b.Name(), err)
			return HistoryRecordResponse{}, nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Warnf("get %s status error: %s", b.Name(), err)
			return HistoryRecordResponse{}, nil
		}

		var response HistoryRecordResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return HistoryRecordResponse{}, errors.Wrap(err, "unmarshal response error")
		}
		return response, nil
	}

	for {
		select {
		case <-ctx.Done():
			return HistoryRecord{}, nil
		case <-t.C:
			result, err := getRecords()
			if err != nil {
				return HistoryRecord{}, errors.Wrap(err, "get records error")
			}
			log.Debugf("get %s records: %d", b.Name(), len(result.Data.HistoryRecords.Records))
			for _, record := range result.Data.HistoryRecords.Records {
				if !strings.EqualFold(record.RequestTxHash, tx) {
					continue
				}
				if record.Result == 0 {
					log.Debugf("tx %s status: pending", tx)
					continue
				}
				if record.Result != 3 {
					return record, errors.Errorf("bridge failed: %d", record.Result)
				}
				return record, nil
			}
		}
	}
}

func approve(ctx context.Context, chainId int64, tokenAddress, spender common.Address, owner wallets.Wallets,
	amount decimal.Decimal, client simulated.Client) error {

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

func Issue(remoteChainId int64, originalToken, originalSender, recipient, rollbackAccount common.Address,
	amount decimal.Decimal, nonce int64, extData []byte) ([]byte, error) {
	contractAbi, err := abi.JSON(strings.NewReader(XTOKEN_ISSUING_NEXT))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack(
		"issue",
		big.NewInt(remoteChainId),
		originalToken,
		originalSender,
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		extData)
}

func ReceiveMessage(srcAppChainId int64, remoteAppAddress, localAppAddress common.Address, message string) ([]byte, error) {
	contractAbi, err := abi.JSON(strings.NewReader(MSGLINE_MESSAGER))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("receiveMessage", big.NewInt(srcAppChainId), remoteAppAddress, localAppAddress, common.Hex2Bytes(message))
}

func FetchMsglineFeeAndParams(ctx context.Context, sourceChainId, targetChainId int64, sourceMessager, targetMessager,
	sender common.Address, payload string) (fee decimal.Decimal, extParams []byte, gas decimal.Decimal, err error) {
	u, err := url.Parse("https://api.msgport.xyz/ormp/fee")
	if err != nil {
		return decimal.Zero, nil, decimal.Zero, errors.Wrap(err, "url parse error")
	}
	query := u.Query()
	query.Set("from_chain_id", strconv.Itoa(int(sourceChainId)))
	query.Set("to_chain_id", strconv.Itoa(int(targetChainId)))
	if !strings.HasPrefix(payload, "0x") {
		payload = "0x" + payload
	}
	query.Set("payload", payload)
	query.Set("from_address", sourceMessager.Hex())
	query.Set("to_address", targetMessager.Hex())
	query.Set("refund_address", sender.Hex())
	u.RawQuery = query.Encode()
	type Fee struct {
		Code int `json:"code"`
		Data struct {
			Fee    string `json:"fee"`
			Params string `json:"params"`
			Gas    struct {
				GasForMessagingLayer decimal.Decimal `json:"gasForMessagingLayer"`
				GasForMsgport        decimal.Decimal `json:"gasForMsgport"`
				Multiplier           decimal.Decimal `json:"multiplier"`
				Total                decimal.Decimal `json:"total"`
			} `json:"gas"`
		} `json:"data"`
	}
	var result Fee
	err = utils.Request(ctx, "GET", u.String(), nil, &result)
	if err != nil {
		return decimal.Zero, nil, decimal.Zero, errors.Wrap(err, "request error")
	}
	if result.Code != 0 {
		return decimal.Zero, nil, decimal.Zero, errors.Errorf("fetch fee error: %d", result.Code)
	}
	return decimal.RequireFromString(result.Data.Fee),
		common.Hex2Bytes(strings.TrimPrefix(result.Data.Params, "0x")), result.Data.Gas.Total, nil
}

func WTokenLockAndXIssue(remoteChainId int64, recipient, rollbackAccount common.Address, amount decimal.Decimal,
	nonce int64, extData, extParams string) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(WTOKEN_CONVERTOR))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("lockAndXIssue",
		big.NewInt(remoteChainId),
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		common.Hex2Bytes(extData),
		common.Hex2Bytes(extParams),
	)
}

func XTokenLockAndXIssue(remoteChainId int64, originalToken, recipient, rollbackAccount common.Address,
	amount decimal.Decimal, nonce int64, extData, extParams string) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(XTOKEN_BACKING_NEXT))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("lockAndXIssue",
		big.NewInt(remoteChainId),
		originalToken,
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		common.Hex2Bytes(strings.TrimPrefix(extData, "0x")),
		common.Hex2Bytes(strings.TrimPrefix(extParams, "0x")))
}

func Unlock(remoteChainId int64, originalToken, originSender, recipient, rollbackAccount common.Address,
	amount decimal.Decimal, nonce int64, extData []byte) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(XTOKEN_BACKING_NEXT))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("unlock",
		big.NewInt(remoteChainId),
		originalToken,
		originSender,
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		extData,
	)
}

func BurnAndXUnlock(xToken, recipient, rollbackAccount common.Address, amount decimal.Decimal, nonce int64,
	extData, extParams []byte) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(XTOKEN_ISSUING_NEXT))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("burnAndXUnlock",
		xToken,
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		extData,
		extParams,
	)
}

func XTokenBurnAndXUnlock(recipient, rollbackAccount common.Address, amount decimal.Decimal, nonce int64,
	extData []byte, extParams []byte) ([]byte, error) {

	contractAbi, err := abi.JSON(strings.NewReader(XTOKEN_CONVERTOR))
	if err != nil {
		return nil, err
	}
	return contractAbi.Pack("burnAndXUnlock",
		recipient,
		rollbackAccount,
		amount.BigInt(),
		big.NewInt(nonce),
		extData,
		extParams,
	)
}

func ReplaceExtData(extData string, sender string) string {
	return strings.ReplaceAll(extData, "{sender}", strings.TrimPrefix(sender, "0x"))
}
