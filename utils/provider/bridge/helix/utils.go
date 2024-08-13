package helix

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

// WaitForBridge waits for the bridge transaction to be confirmed and retrieves the history record
// for the transaction. It uses a time ticker to periodically check for the record. If the context
// is canceled, it returns a context canceled error. Otherwise, it retrieves the history record and
// checks if it is a valid record. If the record is not valid, it returns the record and an error
// indicating the reason. Otherwise, it returns the record and nil error.
//
// Parameters:
// - ctx: The context for the operation
// - sender: The sender address
// - tx: The transaction hash
//
// Returns:
// - record: The history record
// - error: The error, or nil if the record is valid.
func (b *Bridge) WaitForBridge(ctx context.Context, sender common.Address, tx common.Hash) (HistoryRecord, error) {
	var (
		t        = time.NewTicker(time.Second * 3)
		maxTry   = 60
		tryCount int
	)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return HistoryRecord{}, context.Canceled
		case <-t.C:
			tryCount++
			if tryCount > maxTry {
				return HistoryRecord{}, errors.New("wait for bridge timeout")
			}
			record, err := b.RetrieveHistoryRecords(sender, tx)
			if err != nil && !errors.Is(err, error_types.ErrNotFound) {
				logrus.Errorf("wait for bridge timeout, err: %s", err.Error())
				continue
			}
			if record.Result == 0 || record.ResponseTxHash == "" {
				logrus.Debugf("wait for %s bridge, requests tx is %s,  result: %d", sender.Hex(), tx.Hex(), record.Result)
				continue
			}
			if record.Result == 3 {
				return record, nil
			}
			return record, errors.Errorf("bridge failed, result: %d", record.Result)
		}
	}
}

func (b *Bridge) RetrieveHistoryRecords(sender common.Address, tx common.Hash) (HistoryRecord, error) {
	record, err := GetHistoryRecords(sender, tx)
	if err != nil {
		return HistoryRecord{}, err
	}
	return record, nil
}

//func (b *Bridge)QueryLiquidity

func (b *Bridge) IsValidRecord(record HistoryRecord) (string, bool) {
	if record.ResponseTxHash == "" {
		return "waiting for bridge", false
	}
	if record.Result == 0 {
		return "", false
	}
	if record.Result != 3 {
		return fmt.Sprintf("bridge failed, result: %d", record.Result), false
	}
	return "", true
}

// GetValidChains checks the balance on each chain in the given list and filters out chains that do not meet
// the required balance. It retrieves the chain configuration and token information for each chain, and then
// creates an EVM client to interact with the chain. It calls the GetTokenBalance method to get the token balance
// for the specified wallet address. If the balance is greater than or equal to the required amount, the chain
// is considered valid and its chain ID is added to the validChains list. The method returns the list of valid
// chain IDs.
//
// Parameters:
// - ctx: The context for the operation
// - chains: The list of chain names to check
// - tokenName: The name of the token
// - wallet: The wallet address
// - amount: The required amount
//
// Returns:
// - validChains: The list of valid chain IDs
//
// Note: Chains with a chain ID of 0 or a token contract address of "" are considered invalid and are skipped.
func (b *Bridge) GetValidChains(ctx context.Context, targetChainName string, sourceChains []string, tokenName,
	wallet string, amount decimal.Decimal) (validChains []int64) {

	var (
		w sync.WaitGroup
		m sync.Mutex
	)
	for _, chain := range sourceChains {
		w.Add(1)
		go func(chainName string) {
			defer utils.Recover()
			defer w.Done()
			chain := b.config.GetChainConfig(chainName)
			token := b.config.GetTokenInfoOnChain(tokenName, chainName)
			if chain.Id == 0 || token.ContractAddress == "" {
				return
			}
			client, err := chains.NewTryClient(ctx, chain.RpcEndpoints)
			if err != nil {
				logrus.Warnf("get %s chain client error: %s", chainName, err)
				return
			}
			defer client.Close()

			balance, err := chains.GetTokenBalance(ctx, client, token.ContractAddress, wallet, token.Decimals)
			if err != nil {
				logrus.Warnf("get %s chain %s token balance error: %s", chainName, token.Name, err)
				return
			}
			if !balance.GreaterThanOrEqual(amount) {
				return
			}

			maxLimit, err := GetMaxTransfer(
				chainName,
				targetChainName,
				amount,
				common.HexToAddress(token.ContractAddress),
			)
			if err != nil {
				logrus.Warnf("get %s chain %s token max transfer error: %s", chainName, token.Name, err)
				return
			}
			if maxLimit.GreaterThanOrEqual(amount) && !maxLimit.Equal(decimal.Zero) {
				return
			}
			m.Lock()
			defer m.Unlock()
			validChains = append(validChains, int64(constant.GetChainId(chainName)))
		}(chain)
	}
	w.Wait()
	return validChains
}

// FindSourceChain selects the source chains that have a supported token and sufficient balance
// for swapping to the targetChain. It retrieves the supported chains for the given token,
// filters the chains based on the targetChain, and then checks the token balance on each
// remaining chain. If the balance is greater than or equal to the required amount,
// the chain is considered a valid source chain. The method returns a list of valid source chain IDs.
// It returns an error if there are no supported chains or if there is an error retrieving the supported chains.
// Parameters:
// - ctx: The context for the operation
// - targetChainName: The name of the target chain
// - TokenName: The name of the token
// - wallet: The wallet address
// - amount: The required amount
// Returns:
// - sourceChainIds: The list of valid source chain IDs
// - err: The error, or nil if there is no error
func (b *Bridge) FindSourceChain(ctx context.Context, targetChainName, TokenName, wallet string,
	amount decimal.Decimal, sourceChainNames ...string) (sourceChainIds []int64, err error) {

	supportedChains, err := GetTokenSupportedChains(TokenName)
	if err != nil {
		return nil, errors.Wrap(err, "get token supported chains")
	}
	var chains []string
	for _, v := range supportedChains {
		if !utils.InArrayFold(targetChainName, v.ToChains) {
			continue
		}
		if len(sourceChainNames) > 0 && !utils.InArrayFold(v.FromChain, sourceChainNames) {
			continue
		}
		chains = append(chains, v.FromChain)
	}

	// No supported chain
	if len(chains) == 0 {
		return nil, error_types.ErrUnsupportedTokenAndChain
	}

	// Check other chains for the token balance
	return b.GetValidChains(ctx, targetChainName, chains, TokenName, wallet, amount), nil
}

// GetSourceChain retrieves the source chains that have a supported token and sufficient balance for swapping to the targetChain.
// It calls the FindSourceChain method to filter the chains that meet the requirements.
// If there are no supported chains or if there is an error retrieving the supported chains, it returns an error.
// Otherwise, it checks if there are valid source chain IDs. If there are no valid source chain IDs, it returns an error.
// Otherwise, it selects a source chain randomly and returns the source chain ID.
//
// Parameters:
// - ctx: The context for the operation
// - targetChainName: The name of the target chain
// - TokenName: The name of the token
// - wallet: The wallet address
// - amount: The required amount
//
// Returns:
// - sourceChainId: The selected source chain ID
// - err: The error, or nil if there is no error
func (b *Bridge) GetSourceChain(ctx context.Context, targetChainName, TokenName, wallet string,
	amount decimal.Decimal, sourceChainNames ...string) (sourceChainId int64, err error) {

	sourceChainIds, err := b.FindSourceChain(ctx, targetChainName, TokenName, wallet, amount, sourceChainNames...)
	if err != nil {
		return 0, err
	}
	if len(sourceChainIds) == 0 {
		return 0, error_types.ErrUnsupportedTokenAndChain
	}
	return utils.Choose(sourceChainIds), nil
}

type SortedLnBridgeRelayInfosResult struct {
	Data struct {
		SortedLnBridgeRelayInfos struct {
			TransferLimit string `json:"transferLimit"`
			Records       []struct {
				SendToken        common.Address  `json:"sendToken"`
				Relayer          common.Address  `json:"relayer"`
				Margin           decimal.Decimal `json:"margin"`
				BaseFee          decimal.Decimal `json:"baseFee"`
				ProtocolFee      decimal.Decimal `json:"protocolFee"`
				LiquidityFeeRate decimal.Decimal `json:"liquidityFeeRate"`
				LastTransferId   common.Hash     `json:"lastTransferId"`
				WithdrawNonce    decimal.Decimal `json:"withdrawNonce"`
				Bridge           string          `json:"bridge"`
				Typename         string          `json:"__typename"`
			} `json:"records"`
			Typename string `json:"__typename"`
		} `json:"sortedLnBridgeRelayInfos"`
	} `json:"data"`
}

type SortedLnBridgeRelayInfoParams struct {
	OperationName string    `json:"operationName"`
	Variables     Variables `json:"variables"`
	Query         string    `json:"query"`
}

type Variables struct {
	Amount    string `json:"amount"`
	Decimals  int32  `json:"decimals"`
	Token     string `json:"token"`
	FromChain string `json:"fromChain"`
	ToChain   string `json:"toChain"`
}

// GetSortedLnBridgeRelayInfos amount must be wei
func GetSortedLnBridgeRelayInfos(amount decimal.Decimal, decimals int32, fromChain, toChain string,
	token common.Address) (*SortedLnBridgeRelayInfosResult, error) {

	var query = "query sortedLnBridgeRelayInfos($amount: String, $decimals: Int, $bridge: String, $token: String, $fromChain: String, $toChain: String) {\n  sortedLnBridgeRelayInfos(\n    amount: $amount\n    decimals: $decimals\n    bridge: $bridge\n    token: $token\n    fromChain: $fromChain\n    toChain: $toChain\n  ) {\n    transferLimit\n    records {\n      sendToken\n      relayer\n      margin\n      baseFee\n      protocolFee\n      liquidityFeeRate\n      lastTransferId\n      withdrawNonce\n      bridge\n      __typename\n    }\n    __typename\n  }\n}"
	queryParams := SortedLnBridgeRelayInfoParams{
		OperationName: "sortedLnBridgeRelayInfos",
		Variables: Variables{
			Amount:    amount.String(),
			Decimals:  decimals,
			Token:     token.Hex(),
			FromChain: fromChain,
			ToChain:   toChain,
		},
		Query: query,
	}
	var body = bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(queryParams); err != nil {
		return nil, err
	}
	resp, err := http.Post("https://apollo.helixbridge.app/graphql", "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result = new(SortedLnBridgeRelayInfosResult)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

type HistoryRecordsParams struct {
	OperationName string                  `json:"operationName,omitempty"`
	Variables     HistoryRecordsVariables `json:"variables,omitempty"`
	Query         string                  `json:"query,omitempty"`
}

type HistoryRecordsVariables struct {
	Bridges []BridgeType `json:"bridges"`
	Sender  string       `json:"sender"`
	Row     int          `json:"row"`
	Page    int          `json:"page"`
}

func GetHistoryRecords(sender common.Address, tx common.Hash) (HistoryRecord, error) {
	query := "query GetHistory($bridges: [String], $sender: String, $page: Int, $row: Int) {\n  historyRecords(bridges: $bridges, sender: $sender, page: $page, row: $row) {\n    total\n    records {\n      requestTxHash\n      responseTxHash\n      fromChain\n      toChain\n      startTime\n      sendToken\n      sendAmount\n      confirmedBlocks\n      result\n      id\n      __typename\n    }\n    __typename\n  }\n}"
	hrp := &HistoryRecordsParams{
		OperationName: "GetHistory",
		Variables: HistoryRecordsVariables{
			Bridges: []BridgeType{
				LnV2DefaultType,
				LnV2OppositeType,
				LnV3Type,
			},
			Sender: sender.Hex(),
			Row:    0,
			Page:   10,
		},
		Query: query,
	}
	var body = bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(hrp); err != nil {
		return HistoryRecord{}, err
	}
	resp, err := http.Post("https://apollo.helixbridge.app/graphql", "application/json", body)
	if err != nil {
		return HistoryRecord{}, err
	}
	defer resp.Body.Close()
	var result = new(HistoryRecordsResult)
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return HistoryRecord{}, err
	}
	for _, v := range result.Data.HistoryRecords.Records {
		if !strings.EqualFold(v.RequestTxHash, tx.Hex()) {
			continue
		}
		return v, nil
	}
	return HistoryRecord{}, error_types.ErrNotFound
}

type GetSupportChainsResult struct {
	Data struct {
		QueryLnBridgeSupportChains []QueryLnBridgeSupportChains `json:"queryLnBridgeSupportChains"`
	} `json:"data"`
}

type QueryLnBridgeSupportChains struct {
	FromChain string   `json:"fromChain"`
	ToChains  []string `json:"toChains"`
	Typename  string   `json:"__typename"`
}

func GetTokenSupportedChains(tokenName string) ([]QueryLnBridgeSupportChains, error) {
	var query = "query GetSupportChains($token: String!) {\n  queryLnBridgeSupportChains(tokenKey: $token) {\n    fromChain\n    toChains\n    __typename\n  }\n}"
	var queryBody = map[string]interface{}{
		"operationName": "GetSupportChains",
		"query":         query,
		"variables": map[string]string{
			"token": strings.ToUpper(tokenName),
		},
	}
	var body = bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(queryBody); err != nil {
		return nil, err
	}
	resp, err := http.Post("https://apollo.helixbridge.app/graphql", "application/json", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result GetSupportChainsResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Data.QueryLnBridgeSupportChains, nil
}

func GetMaxTransfer(fromChain, toChain string, balance decimal.Decimal,
	tokenAddress common.Address) (decimal.Decimal, error) {

	var query = "query GetMaxTransfer($token: String, $balance: String, $fromChain: String, $toChain: String) {\n  queryMaxTransfer(\n    token: $token\n    balance: $balance\n    fromChain: $fromChain\n    toChain: $toChain\n  )\n}"
	var queryBody = map[string]interface{}{
		"operationName": "GetMaxTransfer",
		"query":         query,
		"variables": map[string]string{
			"token":     tokenAddress.Hex(),
			"balance":   balance.String(),
			"fromChain": fromChain,
			"toChain":   toChain,
		},
	}
	var body = bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(queryBody); err != nil {
		return decimal.Zero, err
	}
	resp, err := http.Post("https://apollo.helixbridge.app/graphql", "application/json", body)
	if err != nil {
		return decimal.Zero, err
	}
	defer resp.Body.Close()
	var result = new(struct {
		Data struct {
			QueryMaxTransfer string `json:"queryMaxTransfer"`
		} `json:"data"`
	})
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return decimal.Zero, err
	}
	return decimal.NewFromString(result.Data.QueryMaxTransfer)
}
