package helix_liquidity_claim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider/bridge/helix"
	"omni-balance/utils/provider/bridge/helix_liquidity_claim/abi/lnv3Bridge"
	"strings"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
)

// BuildTx The client must is item.ToChain
func (c Claim) BuildTx(ctx context.Context, client simulated.Client, Address common.Address, item NeedWithdrawRecords) (*types.LegacyTx, error) {
	fromChannel := messagerAddress[item.FromChain][item.Channel]
	toChannel := messagerAddress[item.ToChain][item.Channel]
	if fromChannel.Cmp(constant.ZeroAddress) == 0 ||
		toChannel.Cmp(constant.ZeroAddress) == 0 {
		return nil, errors.New("channel not found")
	}
	if item.TotalAmount.IsZero() {
		return nil, errors.New("amount is zero")
	}
	messager, ok := messagerInstance[item.Channel]
	if !ok {
		return nil, errors.New("messager not found")
	}

	var (
		fromChainId = constant.GetChainId(item.FromChain)
		toChainId   = constant.GetChainId(item.ToChain)
	)

	appPayload, err := c.encodeWithdrawLiquidity(item.TransferIds,
		big.NewInt(int64(toChainId)), Address)
	if err != nil {
		return nil, errors.Wrap(err, "encode withdraw liquidity")
	}
	payload, err := messager.EncodePayload(ctx, EncodePayloadArgs{
		FromChainId:      toChainId,
		LocalAppAddress:  lnv3BridgeContractAddress[item.ToChain],
		RemoteAppAddress: lnv3BridgeContractAddress[item.FromChain],
		Message:          appPayload,
	})
	if err != nil {
		return nil, errors.Wrap(err, "encode payload")
	}

	params, err := messager.Params(ctx, MessagerArgs{
		ContractAddress: toChannel,
		FromChainId:     toChainId,
		ToChainId:       fromChainId,
		RemoteMessager:  fromChannel,
		Payload:         payload,
		Refunder:        Address,
		Client:          client,
	})
	if err != nil {
		return nil, errors.Wrap(err, "get messager params")
	}
	txData, err := c.WithdrawLiquidity(ctx, fromChainId, item.TransferIds, Address, params.ExtParams)
	if err != nil {
		return nil, errors.Wrap(err, "withdraw liquidity")
	}
	to := lnv3BridgeContractAddress[item.FromChain]
	return &types.LegacyTx{
		To:    &to,
		Value: params.Fee.BigInt(),
		Data:  txData,
	}, nil
}

func (c Claim) WithdrawLiquidity(ctx context.Context, remoteChainId int, transferIds []string, provider common.Address,
	extParams []byte) ([]byte, error) {
	abi, err := lnv3Bridge.Lnv3BridgeMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}
	var _transferIds [][32]byte
	for _, transferId := range transferIds {
		var id [32]byte
		copy(id[:], common.Hex2Bytes(strings.TrimLeft(transferId, "0x")))
		_transferIds = append(_transferIds, id)
	}
	return abi.Pack("requestWithdrawLiquidity", big.NewInt(int64(remoteChainId)), _transferIds, provider, extParams)
}

func (c Claim) encodeWithdrawLiquidity(transferIds []string, toChainId *big.Int, provider common.Address) ([]byte, error) {
	abi, err := lnv3Bridge.Lnv3BridgeMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}
	var _transferIds [][32]byte
	for _, transferId := range transferIds {
		var id [32]byte
		copy(id[:], common.Hex2Bytes(strings.TrimLeft(transferId, "0x")))
		_transferIds = append(_transferIds, id)
	}
	return abi.Pack("withdrawLiquidity", _transferIds, toChainId, provider)
}

func (c Claim) ListNeedWithdrawRecords(ctx context.Context, relayer common.Address, toChain, tokenName string) ([]NeedWithdrawRecords, error) {
	tokenAddress := c.conf.GetTokenAddress(tokenName, toChain)
	claimRecords, err := c.ListClaim(ctx, relayer, common.HexToAddress(tokenAddress), toChain)
	if err != nil {
		return nil, errors.Wrap(err, "list claim")
	}

	if len(claimRecords) == 0 {
		return nil, nil
	}
	var (
		result    []NeedWithdrawRecords
		token     = c.conf.GetTokenInfoOnChainByAddress(tokenAddress, toChain)
		threshold = c.conf.GetTokenThreshold(relayer.Hex(), token.Name, toChain)
	)

	for fromChain, items := range claimRecords {
		for channel, item := range items {
			if item.TotalAmount.IsZero() || item.TotalAmount.LessThan(threshold) {
				log.Debugf("amount is less than threshold, token: %s, chain: %s, amount: %s, threshold: %s",
					token.Name, toChain, item.TotalAmount.String(), threshold.String(),
				)
				continue
			}
			log.Debugf("amount is greater than threshold, token: %s, chain: %s, amount: %s, threshold: %s",
				token.Name, toChain, item.TotalAmount.String(), threshold.String(),
			)
			result = append(result, NeedWithdrawRecords{
				ClaimInfo: *claimRecords[fromChain][channel],
				TokenName: token.Name,
				FromChain: fromChain,
				ToChain:   toChain,
			})
		}
	}
	return result, nil
}

func (c Claim) ListClaim(ctx context.Context, relayer, tokenAddress common.Address, toChain string) (map[string]map[string]*ClaimInfo, error) {
	query := `{historyRecords(row: 100,bridges: ["lnv3"],needWithdrawLiquidity: true,relayer: "%s",toChains: ["%s"],recvTokenAddress: "%s") {total, records { id sender relayer sendTokenAddress lastRequestWithdraw sendAmount recvTokenAddress toChain fromChain }}}`
	query = fmt.Sprintf(query,
		strings.ToLower(relayer.Hex()),
		toChain,
		strings.ToLower(tokenAddress.Hex()),
	)

	var body = bytes.NewBuffer(nil)
	_ = json.NewEncoder(body).Encode(&helix.HistoryRecordsParams{Query: query})
	var records ClaimRecords
	err := utils.Request(ctx, "POST", "https://apollo.helixbridge.app/graphql", body, &records)
	if err != nil {
		return nil, errors.Wrap(err, "list claim")
	}
	if len(records.Data.HistoryRecords.Records) == 0 {
		return nil, nil
	}

	token := c.conf.GetTokenInfoOnChainByAddress(tokenAddress.Hex(), toChain)

	channels, err := c.ListMessageChannel(ctx, relayer, toChain)
	if err != nil {
		return nil, errors.Wrap(err, "list message channel")
	}

	if len(channels) == 0 {
		return nil, nil
	}
	var claims = make(map[string]map[string]*ClaimInfo)
	for _, v := range records.Data.HistoryRecords.Records {
		if _, ok := channels[v.FromChain]; !ok {
			continue
		}
		channel := channels[v.FromChain]
		if _, ok := claims[v.FromChain]; !ok {
			claims[v.FromChain] = make(map[string]*ClaimInfo)
		}
		if _, ok := claims[v.FromChain][channel]; !ok {
			claims[v.FromChain][channel] = &ClaimInfo{}
		}

		ids := strings.Split(v.Id, "-")
		claims[v.FromChain][channel].TransferIds = append(claims[v.FromChain][channel].TransferIds, ids[len(ids)-1])
		claims[v.FromChain][channel].Channel = channel
		claims[v.FromChain][channel].TotalAmount = claims[v.FromChain][channel].TotalAmount.Add(chains.WeiToEth(v.SendAmount.BigInt(), token.Decimals))
	}
	return claims, nil
}

func (c Claim) ListMessageChannel(ctx context.Context, relayer common.Address, toChain string) (map[string]string, error) {
	var result = make(map[string]string)

	query := `{
		queryLnBridgeRelayInfos(
		  relayer: "%s"
		  row: 500
		  toChain: "%s"
		) {
		  records {
			messageChannel
			fromChain
		  }
		}
	  }`
	query = fmt.Sprintf(query, strings.ToLower(relayer.Hex()), strings.ToLower(toChain))
	var body = bytes.NewBuffer(nil)
	_ = json.NewEncoder(body).Encode(&helix.HistoryRecordsParams{Query: query})
	var infos QueryLnBridgeRelayInfos
	err := utils.Request(ctx, "POST", "https://apollo.helixbridge.app/graphql", body, &infos)
	if err != nil {
		return nil, errors.Wrap(err, "list relay info")
	}

	if len(infos.Data.QueryLnBridgeRelayInfos.Records) == 0 {
		return nil, nil
	}
	for _, v := range infos.Data.QueryLnBridgeRelayInfos.Records {
		result[v.FromChain] = v.MessageChannel
	}
	return result, nil
}
