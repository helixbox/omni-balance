package base

import (
	"context"
	"math/big"
	"strings"

	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	log "omni-balance/utils/logging"
	"omni-balance/utils/provider"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	base2ethereum = map[string]tokenConfig{
		"COW": {
			l1Address: common.HexToAddress("0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB"),
			l2Address: common.HexToAddress("0xc694a91e6b071bF030A18BD3053A7fE09B6DaE69"),
		},
		"COMP": {
			l1Address: common.HexToAddress("0xc00e94Cb662C3520282E6f5717214004A7f26888"),
			l2Address: common.HexToAddress("0x9e1028f5f1d5ede59748ffcee5532509976840e0"),
		},
	}
	baseChainId        int64 = 8453
	l2Router                 = common.HexToAddress("0x4200000000000000000000000000000000000010")
	portal                   = common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e")
	disputeGameFactory       = common.HexToAddress("0x43edB88C4B80fDD2AdFF2412A7BebF9dF42cB40e")
	l1RPC                    = "https://erpc.ringdao.com/main/evm/1"
	// l2RPC                    = "https://erpc.ringdao.com/main/evm/8453"
	// l2RPC = "https://mainnet.base.org"
	l2RPC = "https://dimensional-proportionate-daylight.base-mainnet.quiknode.pro/87ee57aa1f4e3ad70ddcaaf1ea313c88bfededa5/"
)

type Base2Ethereum struct {
	config configs.Config
}

func NewL2ToL1(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Base2Ethereum{}, nil
	}
	return &Base2Ethereum{config: conf}, nil
}

func (b *Base2Ethereum) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	_ decimal.Decimal,
) (bool, error) {
	if strings.ToLower(tokenInChainName) == constant.Base && strings.ToLower(tokenOutChainName) == constant.Ethereum {
		if base2ethereum[strings.ToUpper(tokenName)] != (tokenConfig{}) {
			return true, nil
		}
	}
	return false, nil
}

func (b *Base2Ethereum) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	chain := constant.Base
	chainConfig := b.config.GetChainConfig(chain)
	client, err := chains.NewTryClient(ctx, chainConfig.RpcEndpoints)
	if err != nil {
		return nil, nil
	}
	defer client.Close()
	token := b.config.GetTokenInfoOnChain(args.SourceToken, chain)
	balance, err := args.Sender.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, client)
	if err != nil {
		return nil, nil
	}
	log.Debugf("check in get cost for %s from %s to %s, balance: %s, amount %s", args.SourceToken, chain, args.TargetChain, balance.String(), args.Amount.String())
	if args.Amount.GreaterThan(balance) {
		return nil, nil
	}
	if strings.ToLower(args.TargetChain) == constant.Ethereum {
		return provider.TokenInCosts{
			provider.TokenInCost{
				TokenName:  "ETH",
				CostAmount: decimal.NewFromInt(2),
			},
		}, nil
	}
	return nil, nil
}

func buildL2ToL1Tx(ctx context.Context, args provider.SwapParams, client simulated.Client, decimals int32) (*types.DynamicFeeTx, error) {
	var (
		wallet      = args.Sender
		realWallet  = wallet.GetAddress(true)
		tokenConfig = base2ethereum[strings.ToUpper(args.SourceToken)]
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(args.Amount, decimals), 0)

	data, err := Withdraw(ctx, tokenConfig.l2Address, realWallet, amount)
	if err != nil {
		return nil, errors.Wrap(err, "withdraw tx request")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(baseChainId),
		To:      &l2Router,
		Value:   big.NewInt(0),
		Data:    data,
	}, nil
}

func (b *Base2Ethereum) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
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

	if history.Actions == state6 && history.Status == string(provider.TxStatusSuccess) {
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

	args.SourceChain = constant.Base
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
	baseClient, err := chains.NewTryClient(ctx, sourceChainConf.RpcEndpoints)
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "arb dial rpc")
	}
	ethClient, err := chains.NewTryClient(ctx, targetChainConf.RpcEndpoints)
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "eth dial rpc")
	}
	defer baseClient.Close()
	defer ethClient.Close()

	log.Debugf("start transfer %s from %s to %s, amount: %s", args.SourceToken, args.SourceChain, args.TargetChain, args.Amount.String())

	if actionNumber <= 1 && !isActionSuccess {
		recordFn(sh.SetActions(state1).SetStatus(provider.TxStatusPending).Out())
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
		tx, err := buildL2ToL1Tx(ctx, args, baseClient, decimals)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "build tx")
		}

		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeBase2EthBridge)
		txHash, err := wallet.SendTransaction(ctx, tx, baseClient)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(state1).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
	}

	if actionNumber <= 2 && !isActionSuccess {
		recordFn(sh.SetActions(state2).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sh.Tx), baseClient)
		if err != nil {
			recordFn(sh.SetActions(state2).SetStatus(provider.TxStatusPending).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for tx")
		}
		recordFn(sh.SetActions(state2).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 3 && !isActionSuccess {
		recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusPending).Out())
		log.Debugf("waiting for prove, tx: %s", sr.Tx)
		proveTx, err := b.BuildProveTx(ctx, sr.Tx, args.Sender.GetAddress(true).Hex())
		if err != nil {
			recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait claim tx")
		}
		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeBase2EthProve)
		_, err = wallet.SendTransaction(ctx, proveTx, ethClient)
		if err != nil {
			recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusFailed).Out())
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 4 && !isActionSuccess {
		recordFn(sh.SetActions(state4).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sh.Tx), ethClient)
		if err != nil {
			recordFn(sh.SetActions(state4).SetStatus(provider.TxStatusPending).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for tx")
		}
		recordFn(sh.SetActions(state4).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 5 && !isActionSuccess {
		recordFn(sh.SetActions(state5).SetStatus(provider.TxStatusPending).Out())
		log.Debugf("waiting for claim, tx: %s", sr.Tx)
		claimTx, err := b.BuildClaimTx(ctx, sr.Tx, args.Sender.GetAddress(true).Hex())
		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeBase2EthClaim)
		txHash, err := wallet.SendTransaction(ctx, claimTx, ethClient)
		if err != nil {
			recordFn(sh.SetActions(state5).SetStatus(provider.TxStatusFailed).Out())
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(state5).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
	}

	if actionNumber <= 6 && !isActionSuccess {
		recordFn(sh.SetActions(state6).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sh.Tx), ethClient)
		if err != nil {
			recordFn(sh.SetActions(state6).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for claim tx")
		}
		recordFn(sh.SetActions(state6).SetStatus(provider.TxStatusSuccess).Out())
	}
	return sr.SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetChain).Out(), nil
}

func (b *Base2Ethereum) BuildProveTx(ctx context.Context, txHash, trader string) (*types.DynamicFeeTx, error) {
	proveData, err := WaitForProve(ctx, txHash, trader)
	if err != nil {
		return nil, errors.Wrap(err, "prove tx data")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(EthereumChianId),
		To:      &portal,
		Value:   big.NewInt(0),
		Data:    proveData,
	}, nil
}

func (b *Base2Ethereum) BuildClaimTx(ctx context.Context, txHash, trader string) (*types.DynamicFeeTx, error) {
	claimData, err := WaitForClaim(ctx, txHash, trader)
	if err != nil {
		return nil, errors.Wrap(err, "claim tx data")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(EthereumChianId),
		To:      &portal,
		Value:   big.NewInt(0),
		Data:    claimData,
	}, nil
}

func (b *Base2Ethereum) Help() []string {
	return []string{"https://docs.optimism.io/app-developers/tutorials/bridging/cross-dom-bridge-erc20"}
}

func (b *Base2Ethereum) Name() string {
	return "base-ethereum"
}

func (b *Base2Ethereum) Type() configs.ProviderType {
	return configs.Bridge
}
