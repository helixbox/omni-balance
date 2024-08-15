package helix_liquidity_claim

import (
	"context"
	"encoding/json"
	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/provider"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type Claim struct {
	conf configs.Config
}

func (c Claim) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	return nil, error_types.ErrUnsupportedTokenAndChain
}

func (c Claim) CheckToken(ctx context.Context, tokenName, tokenInChainName, tokenOutChainName string, amount decimal.Decimal) (bool, error) {
	return false, error_types.ErrUnsupportedTokenAndChain
}

func (c Claim) Swap(ctx context.Context, args provider.SwapParams) (provider.SwapResult, error) {
	var (
		result = &provider.SwapResult{
			ProviderType: c.Type(),
			ProviderName: c.Name(),
			Order:        args.Order,
		}
		sh   = new(provider.SwapHistory).SetProviderType(c.Type()).SetProviderName(c.Name())
		last = args.LastHistory
		tx   = last.Tx
	)

	if args.SourceChain == "" || args.TargetChain == "" || args.TargetToken == "" || args.SourceToken == "" {
		log.Error("sourceChain or targetChain or targetToken or sourceToken is empty")
		err := error_types.ErrUnsupportedTokenAndChain
		return result.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
	}
	if len(args.Order) == 0 {
		log.Warnf("order is nil")
		err := errors.New("order is nil")
		return result.SetError(err).Out(), err
	}

	var item NeedWithdrawRecords
	if err := json.Unmarshal([]byte(args.Order), &item); err != nil {
		log.Errorf("unmarshal order failed, err: %v", err)
		return result.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
	}
	if item.ToChain == "" {
		log.Errorf("toChain is empty")
		err := errors.Errorf("toChain is empty")
		return result.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
	}
	client, err := chains.NewTryClient(ctx, c.conf.GetChainConfig(args.TargetChain).RpcEndpoints)
	if err != nil {
		return result.SetStatus(provider.TxStatusFailed).SetError(err).Out(), err
	}
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.TargetChain)
	defer client.Close()
	if tx == "" {
		txData, err := c.BuildTx(ctx, client, args.Sender, item)
		if err != nil {
			return result.SetError(err).Out(), err
		}
		txHash, err := args.Sender.SendTransaction(ctx, txData, client)
		if err != nil {
			log.Errorf("send tx failed, err: %v", err)
			return result.SetError(err).SetStatus(provider.TxStatusFailed).Out(), err
		}
		log.Debugf("send tx success, tx: %s", txHash.Hex())
		tx = txHash.Hex()
		sh = sh.SetTx(txHash.Hex())
		result = result.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
		args.RecordFn(sh.SetStatus(provider.TxStatusSuccess).Out())
	}
	if err := args.Sender.WaitTransaction(ctx, common.HexToHash(tx), client); err != nil {
		return result.SetError(err).Out(), err
	}
	return result.SetStatus(provider.TxStatusSuccess).Out(), nil
}

func (c Claim) Help() []string {
	return []string{
		"This provider can only be called by helix-claim bot.",
		"Claim assets from helix relayer contract following the order.",
		"Process flow from https://github.com/helix-bridge/relayer/blob/main/src/relayer/relayer.service.ts#L632-L700",
	}
}

func (c Claim) Name() string {
	return "helix-claim"
}

func (c Claim) Type() configs.ProviderType {
	return configs.Bridge
}
