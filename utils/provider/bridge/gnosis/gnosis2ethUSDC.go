package gnosis

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
	l1Address    = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	l2Address    = common.HexToAddress("0x2a22f9c3b484c3629090FeED35F17Ff8F88f76F0")
	midAddress   = common.HexToAddress("0xDDAfbb505ad214D7b80b1f830fcCc89B60fb7A83")
	l2Transmuter = common.HexToAddress("0x0392A2F5Ac47388945D8c84212469F545fAE52B2")
)

type Gnosis2EthereumUSDC struct {
	config configs.Config
}

func NewL2ToL1USDC(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Gnosis2EthereumUSDC{}, nil
	}
	return &Gnosis2EthereumUSDC{config: conf}, nil
}

func (b *Gnosis2EthereumUSDC) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	_ decimal.Decimal,
) (bool, error) {
	if strings.ToLower(tokenInChainName) == constant.Gnosis && strings.ToLower(tokenOutChainName) == constant.Ethereum {
		if strings.ToUpper(tokenName) != "USDC" {
			return true, nil
		}
	}
	return false, nil
}

func (b *Gnosis2EthereumUSDC) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if strings.ToUpper(args.SourceToken) != "USDC" {
		return nil, errors.New("only support USDC")
	}

	chain := constant.Gnosis
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

func buildUsdcTransmuterTx(ctx context.Context, args provider.SwapParams, client simulated.Client, decimals int32) (*types.DynamicFeeTx, error) {
	amount := decimal.NewFromBigInt(chains.EthToWei(args.Amount, decimals), 0)

	err := Approve(ctx, gnosisChainId, l2Address, l2Transmuter, args.Sender, amount, client)
	if err != nil {
		return nil, errors.Wrap(err, "approve")
	}

	data, err := TransmuterWithdraw(ctx, amount)
	if err != nil {
		return nil, errors.Wrap(err, "transmuter tx request")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(gnosisChainId),
		To:      &l2Transmuter,
		Value:   big.NewInt(0),
		Data:    data,
	}, nil
}

func buildL2ToL1TxUSDC(ctx context.Context, args provider.SwapParams, client simulated.Client, decimals int32) (*types.DynamicFeeTx, error) {
	var (
		wallet     = args.Sender
		realWallet = wallet.GetAddress(true)
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(args.Amount, decimals), 0)

	data, err := Withdraw(ctx, l2Router, amount, realWallet.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "withdraw tx request")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(gnosisChainId),
		To:      &midAddress,
		Value:   big.NewInt(0),
		Data:    data,
	}, nil
}

func (b *Gnosis2EthereumUSDC) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
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

	args.SourceChain = constant.Gnosis
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
		SetTx(args.Tx)

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
		tx, err := buildUsdcTransmuterTx(ctx, args, baseClient, decimals)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "build tx")
		}

		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeGnosisTransmuter)
		txHash, err := wallet.SendTransaction(ctx, tx, baseClient)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(state1).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
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
		proveTx, err := buildL2ToL1TxUSDC(ctx, args, baseClient, decimals)
		if err != nil {
			recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait claim tx")
		}
		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeGnosisWithdraw)
		txHash, err := wallet.SendTransaction(ctx, proveTx, baseClient)
		if err != nil {
			recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusFailed).Out())
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(state3).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
	}
	//
	if actionNumber <= 4 && !isActionSuccess {
		recordFn(sh.SetActions(state4).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sh.Tx), baseClient)
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
		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeGnosisClaim)
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

func (b *Gnosis2EthereumUSDC) BuildClaimTx(ctx context.Context, txHash, trader string) (*types.DynamicFeeTx, error) {
	claimData, err := WaitForClaim(ctx, txHash, trader)
	if err != nil {
		return nil, errors.Wrap(err, "claim tx data")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(EthereumChianId),
		To:      &l1Claimer,
		Value:   big.NewInt(0),
		Data:    claimData,
	}, nil
}

func (b *Gnosis2EthereumUSDC) Help() []string {
	return []string{"https://docs.optimism.io/app-developers/tutorials/bridging/cross-dom-bridge-erc20"}
}

func (b *Gnosis2EthereumUSDC) Name() string {
	return "gnosis-ethereum-usdc"
}

func (b *Gnosis2EthereumUSDC) Type() configs.ProviderType {
	return configs.Bridge
}
