package gnosis

import (
	"context"
	"math/big"
	"strings"
	"time"

	"omni-balance/utils/chains"
	"omni-balance/utils/configs"
	"omni-balance/utils/constant"
	"omni-balance/utils/provider"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	ethereum2gnosis = map[string]tokenConfig{
		"COW": {
			l1Address: common.HexToAddress("0xDEf1CA1fb7FBcDC777520aa7f396b4E015F497aB"),
			l2Address: common.HexToAddress("0x177127622c4A00F3d409B75571e12cB3c8973d3c"),
		},
		"USDC": {
			l1Address: common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"),
			l2Address: common.HexToAddress("0x2a22f9c3b484c3629090FeED35F17Ff8F88f76F0"),
		},
	}
	EthereumChianId           int64 = 1
	l1Router                        = common.HexToAddress("0x88ad09518695c6c3712AC10a214bE5109a655671")
	L2_GNOSIS_USDC_TRANSMUTER       = common.HexToAddress("0x0392A2F5Ac47388945D8c84212469F545fAE52B2")
)

type Ethereum2Gnosis struct {
	config configs.Config
}

func NewL1ToL2(conf configs.Config, noInit ...bool) (provider.Provider, error) {
	if len(noInit) > 0 && noInit[0] {
		return &Ethereum2Gnosis{}, nil
	}
	return &Ethereum2Gnosis{config: conf}, nil
}

func buildL1ToL2Tx(ctx context.Context, args provider.SwapParams, client simulated.Client, decimals int32) (*types.DynamicFeeTx, error) {
	var (
		wallet      = args.Sender
		realWallet  = wallet.GetAddress(true)
		tokenConfig = ethereum2gnosis[strings.ToUpper(args.SourceToken)]
	)
	amount := decimal.NewFromBigInt(chains.EthToWei(args.Amount, decimals), 0)

	err := Approve(ctx, EthereumChianId, tokenConfig.l1Address, l1Router,
		wallet, amount, client)
	if err != nil {
		return nil, errors.Wrap(err, "approve")
	}

	receiver := realWallet
	data := []byte{}
	if tokenConfig.l1Address == common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") {
		receiver = L2_GNOSIS_USDC_TRANSMUTER
		argType, err := abi.NewType("address", "", nil)
		if err != nil {
			return nil, erros.Wrap(err, "new abi type")
		}
		data = err := abi.Arguments{{Type: argType}}.Pack(realWallet)
		if err != nil {
			return nil, errors.Wrap(err, "pack data")
		}
	}

	data, err := Deposit(ctx, tokenConfig.l1Address, receiver, amount, data)
	if err != nil {
		return nil, errors.Wrap(err, "deposit tx request")
	}

	return &types.DynamicFeeTx{
		ChainID: big.NewInt(EthereumChianId),
		To:      &l1Router,
		Value:   big.NewInt(0),
		Data:    data,
	}, nil
}

func (b *Ethereum2Gnosis) CheckToken(_ context.Context, tokenName, tokenInChainName, tokenOutChainName string,
	_ decimal.Decimal,
) (bool, error) {
	if strings.ToLower(tokenInChainName) == constant.Ethereum && strings.ToLower(tokenOutChainName) == constant.Gnosis {
		if ethereum2gnosis[strings.ToUpper(tokenName)] != (tokenConfig{}) {
			return true, nil
		}
	}
	return false, nil
}

func (b *Ethereum2Gnosis) GetCost(ctx context.Context, args provider.SwapParams) (provider.TokenInCosts, error) {
	if strings.ToLower(args.TargetChain) == constant.Gnosis {
		return provider.TokenInCosts{
			provider.TokenInCost{
				TokenName:  "ETH",
				CostAmount: decimal.NewFromInt(0),
			},
		}, nil
	}
	return nil, nil
}

func (b *Ethereum2Gnosis) Swap(ctx context.Context, args provider.SwapParams) (result provider.SwapResult, err error) {
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

	args.SourceChain = constant.Ethereum
	args.TargetChain = constant.Gnosis

	if args.SourceChain == args.TargetChain && history.Status == string(provider.TxStatusSuccess) {
		log.Debugf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
		return provider.SwapResult{}, errors.Errorf("source chain %s and target chain %s is same", args.SourceChain, args.TargetChain)
	}

	actionNumber := Action2Int(history.Actions)
	sourceChainConf := b.config.GetChainConfig(args.SourceChain)
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
	ethClient, err := chains.NewTryClient(ctx, sourceChainConf.RpcEndpoints)
	if err != nil {
		return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "dial rpc")
	}
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
		tx, err := buildL1ToL2Tx(ctx, args, ethClient, decimals)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "build tx")
		}

		ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, chains.SignTxTypeEth2BaseBridge)
		log.Debugf("waiting for send deposit tx.....")
		txHash, err := wallet.SendTransaction(ctx, tx, ethClient)
		if err != nil {
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "send tx")
		}
		recordFn(sh.SetActions(sourceChainSendingAction).SetStatus(provider.TxStatusSuccess).Out())
		sh = sh.SetTx(txHash.Hex())
		sr = sr.SetTx(txHash.Hex()).SetOrderId(txHash.Hex())
	}

	if actionNumber <= 2 && !isActionSuccess {
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusPending).Out())
		err = wallet.WaitTransaction(ctx, common.HexToHash(sr.Tx), ethClient)
		if err != nil {
			recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusPending).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for tx")
		}
		recordFn(sh.SetActions(sourceChainSentAction).SetStatus(provider.TxStatusSuccess).Out())
	}

	if actionNumber <= 4 && !isActionSuccess {
		recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusPending).Out())
		log.Debugf("waiting for bridge success")
		tx, err := wallet.GetRealHash(ctx, common.HexToHash(sr.Tx), ethClient)
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "get real tx hash")
		}

		childTx, err := b.WaitForBridgeSuccess(ctx, tx.Hex(), args.Sender.GetAddress(true).Hex())
		if err != nil {
			recordFn(sh.SetActions(targetChainSendingAction).SetStatus(provider.TxStatusFailed).Out(), err)
			return sr.SetStatus(provider.TxStatusFailed).SetError(err).Out(), errors.Wrap(err, "wait for bridge success")
		}
		sr.SetTx(childTx).SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetChain)

		recordFn(sh.SetActions(targetChainReceivedAction).SetStatus(provider.TxStatusSuccess).SetCurrentChain(args.TargetChain).Out())
	}
	return sr.Out(), nil
}

func (b *Ethereum2Gnosis) WaitForBridgeSuccess(ctx context.Context, txHash, trader string) (string, error) {
	const maxRetries = 10
	const retryInterval = 2 * time.Minute

	var receiveTx string
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		receiveTx, err = WaitForChildTransactionReceipt(ctx, txHash, trader)
		if err == nil && receiveTx != "" {
			return receiveTx, nil
		}
		if attempt < maxRetries-1 {
			log.Infof("wait for child transaction receiptTx: %s, trader: %s, receiveTx: %s, err: %v, retrying in %v (attempt %d/%d)", txHash, trader, receiveTx, err, retryInterval, attempt+1, maxRetries)
			timer := time.NewTimer(retryInterval)
			select {
			case <-ctx.Done():
				timer.Stop()
				return "", ctx.Err()
			case <-timer.C:
			}
		}
	}
	if err != nil {
		return "", err
	}
	if receiveTx == "" {
		return "", errors.New("no receive tx")
	}
	return receiveTx, nil
}

func (b *Ethereum2Gnosis) Help() []string {
	return []string{"https://bridge.gnosischain.com/"}
}

func (b *Ethereum2Gnosis) Name() string {
	return "ethereum-gnosis"
}

func (b *Ethereum2Gnosis) Type() configs.ProviderType {
	return configs.Bridge
}
