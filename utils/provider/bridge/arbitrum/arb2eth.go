package arbitrum

import (
	"context"
	"math/big"
	"strings"

	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type tokenConfig struct {
	l1Address common.Address
	l2Address common.Address
	gateway   common.Address
}

var (
	arbitrum2ethereum = map[string]tokenConfig{
		"COW": {
			l1Address: common.HexToAddress("0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB"),
			l2Address: common.HexToAddress("0xcb8b5CD20BdCaea9a010aC1F8d835824F5C87A04"),
			gateway:   common.HexToAddress("0x09e9222E96E7B4AE2a407B98d48e330053351EEe"),
		},
	}
	arbitrumChainId int64 = 42161
	router                = common.HexToAddress("0x5288c571Fd7aD117beA99bF60FE0846C4E84F933")
)

type Arbitrum2Ethereum struct {
	config configs.Config
}

func NewL2ToL1(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Arbitrum2Ethereum{}, nil
	}
	return &Arbitrum2Ethereum{config: conf}, nil
}

func buildL2ToL1Tx(ctx context.Context, args provider.SwapParams, client simulated.Client, decimals int32) (*types.DynamicFeeTx, error) {
	var (
		wallet      = args.Sender
		realWallet  = wallet.GetAddress(true)
		tokenConfig = arbitrum2ethereum[strings.ToUpper(args.SourceToken)]
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(args.Amount, decimals), 0)

	err := Approve(ctx, arbitrumChainId, tokenConfig.l2Address, tokenConfig.gateway,
		wallet, amount, client)
	if err != nil {
		return nil, errors.Wrap(err, "approve")
	}

	data, err := Withdraw(ctx, tokenConfig.l1Address, realWallet, amount)
	if err != nil {
		return nil, errors.Wrap(err, "withdraw tx request")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(arbitrumChainId),
		To:      &router,
		Value:   big.NewInt(0),
		Data:    data,
	}, nil
}

func (b *Arbitrum2Ethereum) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	_ decimal.Decimal,
) (bool, error) {
	if strings.ToLower(tokenInChainName) == constant.Arbitrum && strings.ToLower(tokenOutChainName) == constant.Ethereum {
		if strings.ToUpper(tokenName) == "COW" {
			return true, nil
		}
	}
	return false, nil
}

func (b *Arbitrum2Ethereum) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if strings.ToLower(args.TargetChain) == constant.Ethereum {
		return provider.TokenInCosts{
			provider.TokenInCost{
				TokenName:  "ETH",
				CostAmount: decimal.NewFromInt(0),
			},
		}, nil
	}
	return nil, nil
}

func (b *Arbitrum2Ethereum) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
	var (
		history  = args.LastHistory
		recordFn = func(s provider.SwapHistory, errs ...error) {
			s.ProviderType = string(b.Type())
			s.ProviderName = b.Name()
			s.Amount = args.Amount
			if args.RecordFn == nil {
				return
			}
			args.RecordFn(s, errs...)
		}
	)

	if history.Actions == targetChainReceivedAction && history.Status == string(provider.TxStatusSuccess) {
		log.Debugf("target chain received, order id: %s", history.Tx)
		return provider.SwapResult{
			ProviderType: b.Type(),
			ProviderName: b.Name(),
			OrderId:      history.Tx,
			Status:       provider.TxStatusSuccess,
			CurrentChain: args.TargetChain,
			Tx:           history.Tx,
		}, nil
	}

	args.SourceChain = constant.Arbitrum
	args.TargetChain = constant.Ethereum

	if args.SourceChain == args.TargetChain && history.Status == string(provider.TxStatusSuccess) {
		log.Debugf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
		return provider.SwapResult{}, errors.Errorf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
	}

	actionNumber := Action2Int(history.Actions)
	sourceChainConf := b.config.GetChainConfig(args.SourceChain)
	targetChainConf := b.config.GetChainConfig(args.TargetChain)
	sourceToken := b.config.GetTokenInfoOnChain(args.SourceToken, args.SourceChain)
	decimals := sourceToken.Decimals

	sr := new(provider.SwapResult).
		SetTokenInName(args.SourceToken).
		SetTokenInChainName(args.SourceChain).
		SetProviderName(b.Name()).
		SetProviderType(b.Type()).
		SetCurrentChain(args.SourceChain).
		SetTx(args.LastHistory.Tx)

	sh := &provider.SwapHistory{
		ProviderName: b.Name(),
		ProviderType: string(b.Type()),
		Amount:       args.Amount,
		CurrentChain: args.SourceChain,
		Tx:           history.Tx,
	}
	isActionSuccess := history.Status == string(provider.TxStatusSuccess)
	ctx = context.WithValue(ctx, constant.ChainNameKeyInCtx, args.SourceChain)

	wallet := args.Sender
	arbClient, err := chains.NewTryClient(ctx, sourceChainConf.RpcEndpoints)
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "arb dial rpc")
	}
	ethClient, err := chains.NewTryClient(ctx, targetChainConf.RpcEndpoints)
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "eth dial rpc")
	}
	defer arbClient.Close()
	defer ethClient.Close()

	log.Debugf("start transfer %s from %s to %s, amount: %s", args.SourceToken, args.SourceChain, args.TargetChain, args.Amount.String())

	if actionNumber <= 1 && !isActionSuccess {
		recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		sr = sr.SetReciever(args.Receiver)
		ctx = provider.WithNotify(ctx, provider.WithNotifyParams{
			OrderId:         args.OrderId,
			Receiver:        common.HexToAddress(args.Receiver),
			TokenIn:         args.SourceToken,
			TokenOut:        args.TargetToken,
			TokenInChain:    args.SourceChain,
			TokenOutChain:   args.TargetChain,
			ProviderName:    b.Name(),
			TokenInAmount:   args.Amount,
			TokenOutAmount:  args.Amount,
			TransactionType: provider.TransferTransactionAction,
		})
		tx, err := buildL2ToL1Tx(ctx, args, arbClient, decimals)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "build tx")
		}

		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeArb2EthBridge)
		txHash, err := wallet.SendTransaction(ctx, tx, arbClient)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
	}

	if actionNumber <= 2 && !isActionSuccess {
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sr.Tx), arbClient)
		if err != nil {
			recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusPending).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for tx")
		}
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 3 && !isActionSuccess {
		recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		tx, err := wallet.GetRealHash(ctx, common.HexToHash(sr.Tx), arbClient)
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "get real tx hash")
		}

		log.Debugf("waiting for claim")
		claimTx, err := b.BuildClaimTx(ctx, tx.Hex())
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait claim tx")
		}
		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeArb2EthClaim)
		txHash, err := wallet.SendTransaction(ctx, claimTx, ethClient)
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out())
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}

		recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
	}

	if actionNumber <= 4 && !isActionSuccess {
		recordFn(sh.SetActions(targetChainReceivedAction).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sh.Tx), ethClient)
		if err != nil {
			recordFn(sh.SetActions(targetChainReceivedAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for claim tx")
		}
		recordFn(sh.SetActions(targetChainReceivedAction).SetStatus(provider.TxStatusSuccess).Out())
		sr.SetOrder(sh.Tx)
	}
	return sr.SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetChain).Out(), nil
}

func (b *Arbitrum2Ethereum) BuildClaimTx(ctx context.Context, txHash string) (*types.DynamicFeeTx, error) {
	txRequest, err := WaitForClaim(ctx, txHash)
	if err != nil {
		return nil, errors.Wrap(err, "claim tx request")
	}

	toAddr := common.HexToAddress(txRequest.To)
	value, ok := new(big.Int).SetString(txRequest.Value, 10)
	if !ok {
		return nil, errors.New("invalid value string")
	}
	return &types.DynamicFeeTx{
		ChainID: big.NewInt(EthereumChianId),
		To:      &toAddr,
		Value:   value,
		Data:    common.Hex2Bytes(txRequest.Data),
	}, nil
}

func (b *Arbitrum2Ethereum) Help() []string {
	return []string{"See https://docs.arbitrum.io/build-decentralized-apps/token-bridging/bridge-tokens-programmatically/how-to-bridge-tokens-standard"}
}

func (b *Arbitrum2Ethereum) Name() string {
	return "arbitrum-ethereum"
}

func (b *Arbitrum2Ethereum) Type() configs.ProviderType {
	return configs.Bridge
}
