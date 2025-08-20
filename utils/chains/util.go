package chains

import (
	"context"
	"math/big"
	"strings"
	"time"

	"omni-balance/utils/constant"
	"omni-balance/utils/enclave"
	"omni-balance/utils/erc20"
	"omni-balance/utils/error_types"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

func getDecimals(decimals ...int32) int32 {
	var currentDecimals int32 = 18
	if len(decimals) != 0 && decimals[0] != 0 {
		currentDecimals = decimals[0]
	}
	return currentDecimals
}

func EthToWei(v decimal.Decimal, decimals ...int32) *big.Int {
	decimalStr := v.Mul(decimal.New(1, getDecimals(decimals...)))
	wei, _ := new(big.Int).SetString(decimalStr.Truncate(0).String(), 10)
	return wei
}

func WeiToEth(v *big.Int, decimals ...int32) decimal.Decimal {
	return decimal.NewFromBigInt(v, 0).Div(decimal.New(1, getDecimals(decimals...)))
}

// GetTokenBalance retrieves the token balance of a specified wallet address.
// ctx: Context to control the function's execution, including cancellation and timeouts.
// tokenAddress: Contract address of the token; if "0x0000000000000000000000000000000000000000", it indicates querying the native token (e.g., ETH) balance.
// walletAddress: Address of the wallet whose token balance is being queried.
// tokenDecimal: Optional argument specifying the number of decimal places for the token. Required and must not be zero when querying the native token balance.
// Returns the token balance in its smallest unit and an error if any occurs.
func GetTokenBalance(ctx context.Context, client simulated.Client, tokenAddress, walletAddress string,
	tokenDecimal ...interface{},
) (decimal.Decimal, error) {
	var (
		balance *big.Int
		err     error
	)
	if strings.EqualFold(tokenAddress, constant.ZeroAddress.Hex()) {
		if len(tokenDecimal) == 0 && tokenDecimal[0] != 0 {
			return decimal.Zero, errors.Errorf("token decimal is required")
		}
		tokenDecimals := cast.ToInt32(tokenDecimal[0])
		balance, err = client.BalanceAt(ctx, common.HexToAddress(walletAddress), nil)
		if err != nil {
			return decimal.Zero, errors.Wrap(err, "get balance error")
		}
		return WeiToEth(balance, tokenDecimals), nil
	}

	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get erc20 abi error")
	}
	input, err := erc20Abi.Pack("balanceOf", common.HexToAddress(walletAddress))
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "pack erc20 balanceOf error")
	}
	to := common.HexToAddress(tokenAddress)
	result, err := client.CallContract(ctx, ethereum.CallMsg{To: &to, Data: input}, nil)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "call erc20 balanceOf error")
	}

	balance = new(big.Int).SetBytes(result)
	if len(tokenDecimal) == 0 {
		input, err := erc20Abi.Pack("decimals")
		if err != nil {
			return decimal.Zero, errors.Wrap(err, "pack erc20 decimals error")
		}
		tokenDecimals, err := client.CallContract(ctx, ethereum.CallMsg{To: &to, Data: input}, nil)
		if err != nil {
			return decimal.Zero, errors.Wrap(err, "call erc20 decimals error")
		}
		tokenDecimal = append(tokenDecimal, new(big.Int).SetBytes(tokenDecimals).Int64())
	}
	return WeiToEth(balance, cast.ToInt32(tokenDecimal[0])), nil
}

type SignTxType string

const (
	SignTxTypeApprove        SignTxType = "approve"
	SignTxTypeTransfer       SignTxType = "transfer"
	SignTxTypeEth2ArbBridge  SignTxType = "eth-arb-bridge"
	SignTxTypeArb2EthBridge  SignTxType = "arb-eth-bridge"
	SignTxTypeArb2EthClaim   SignTxType = "arb-eth-claim"
	SignTxTypeEth2BaseBridge SignTxType = "eth-base-bridge"
	SignTxTypeBase2EthBridge SignTxType = "base-eth-bridge"
	SignTxTypeBase2EthProve  SignTxType = "base-eth-prove"
	SignTxTypeBase2EthClaim  SignTxType = "base-eth-claim"
	SignTxTypeGnosisDeposit  SignTxType = "gnosis-deposit"
	SignTxTypeGnosisWithdraw SignTxType = "gnosis-withdraw"
	SignTxTypeGnosisClaim    SignTxType = "gnosis-claim"
)

func SignTx(tx *types.Transaction, privateKey string, chainId int64, signTxType SignTxType) (*types.Transaction, error) {
	if !strings.HasPrefix(privateKey, "http") {
		return nil, errors.New("only support enclave version")
	}

	client := enclave.NewClient(privateKey)

	log.Debugf("sign tx type: %s, chainId: %d", signTxType, chainId)

	switch signTxType {
	case SignTxTypeApprove:
		return client.SignErc20Approve(tx, chainId)
	case SignTxTypeTransfer:
		if tx.Value().Cmp(big.NewInt(0)) > 0 {
			return nil, errors.Wrap(error_types.ErrEnclaveNotSupportNativeToken, "enclave native token not support")
		}
		return client.SignErc20Transfer(tx, chainId)
	case SignTxTypeEth2ArbBridge:
		return client.SignArbitrumDeposit(tx, chainId)
	case SignTxTypeArb2EthBridge:
		return client.SignArbitrumWithdraw(tx, chainId)
	case SignTxTypeArb2EthClaim:
		return client.SignArbitrumClaim(tx, chainId)
	case SignTxTypeEth2BaseBridge:
		return client.SignBaseDeposit(tx, chainId)
	case SignTxTypeBase2EthBridge:
		return client.SignBaseWithdraw(tx, chainId)
	case SignTxTypeBase2EthProve:
		return client.SignBaseProve(tx, chainId)
	case SignTxTypeBase2EthClaim:
		return client.SignBaseClaim(tx, chainId)
	case SignTxTypeGnosisDeposit:
		return client.SignGnosisDeposit(tx, chainId)
	case SignTxTypeGnosisWithdraw:
		return client.SignGnosisWithdraw(tx, chainId)
	case SignTxTypeGnosisClaim:
		return client.SignGnosisClaim(tx, chainId)
	default:
		return nil, errors.New("sign tx type not support")
	}
	return nil, errors.New("sign tx type not support")
}

func SignMsg(msg []byte, privateKey string) (sig []byte, err error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return crypto.Sign(msg, key)
}

func WaitForTx(ctx context.Context, client simulated.Client, txHash common.Hash) error {
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return errors.New("context cancelled")
		case <-t.C:
			tx, err := client.TransactionReceipt(ctx, txHash)
			if errors.Is(err, context.Canceled) ||
				errors.Is(err, context.DeadlineExceeded) {
				return err
			}
			if errors.Is(err, ethereum.NotFound) {
				log.Debugf("tx not found, txHash: %s, try again later", txHash.Hex())
				continue
			}
			if err != nil {
				log.Debugf("get tx receipt error: %s, try again later", err)
				continue
			}
			if tx.Status == 0 {
				return errors.New("tx failed")
			}
			log.Debugf("tx success, txHash: %s", txHash.Hex())
			return nil
		}
	}
}

type TokenApproveParams struct {
	ChainId         int64
	TokenAddress    common.Address
	Owner           common.Address
	SendTransaction func(ctx context.Context, tx *types.DynamicFeeTx, client simulated.Client) (common.Hash, error)
	WaitTransaction func(ctx context.Context, txHash common.Hash, client simulated.Client) error
	Spender         common.Address
	AmountWei       decimal.Decimal
	IsNotWaitTx     bool
	Client          simulated.Client
}

func TokenApprove(ctx context.Context, args TokenApproveParams) error {
	contract, err := erc20.NewTokenCaller(args.TokenAddress, args.Client)
	if err != nil {
		return errors.Wrap(err, "erc20 caller")
	}
	allowanceWei, err := contract.Allowance(nil, args.Owner, args.Spender)
	if err != nil {
		return errors.Wrap(err, "erc20 allowance")
	}
	if decimal.NewFromBigInt(allowanceWei, 0).GreaterThanOrEqual(args.AmountWei) {
		return nil
	}
	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return errors.Wrap(err, "get abi")
	}
	input, err := erc20Abi.Pack("approve", args.Spender, args.AmountWei.BigInt())
	if err != nil {
		return errors.Wrap(err, "abi pack")
	}

	ctx = context.WithValue(ctx, constant.SignTxKeyInCtx, SignTxTypeApprove)
	txHash, err := args.SendTransaction(ctx, &types.DynamicFeeTx{
		To:   &args.TokenAddress,
		Data: input,
	}, args.Client)
	if err != nil {
		return errors.Wrap(err, "erc20 approve")
	}

	if args.IsNotWaitTx {
		return nil
	}
	if args.WaitTransaction != nil {
		return args.WaitTransaction(ctx, txHash, args.Client)
	}
	return WaitForTx(ctx, args.Client, txHash)
}

func MostNewAbiType(t string, internalType string, components []abi.ArgumentMarshaling) abi.Type {
	newType, err := abi.NewType(t, internalType, components)
	if err != nil {
		panic(err)
	}
	return newType
}

type SendTokenParams struct {
	Client        simulated.Client
	Sender        common.Address
	TokenAddress  common.Address
	TokenDecimals int32
	ToAddress     common.Address
	AmountWei     decimal.Decimal
}

func BuildSendToken(ctx context.Context, args SendTokenParams) (*types.DynamicFeeTx, error) {
	isNativeToken := strings.EqualFold(args.TokenAddress.Hex(), constant.ZeroAddress.Hex())
	balance, err := GetTokenBalance(ctx, args.Client, args.TokenAddress.Hex(), args.Sender.Hex(), args.TokenDecimals)
	if err != nil {
		return nil, errors.Wrap(err, "get balance")
	}
	balanceWei := decimal.NewFromBigInt(EthToWei(balance, args.TokenDecimals), 0)
	if balanceWei.IsZero() || balanceWei.IsNegative() || balanceWei.LessThan(args.AmountWei) {
		return nil, error_types.ErrInsufficientBalance
	}

	if !isNativeToken {
		nativeTokenBalanceWei, err := args.Client.BalanceAt(ctx, args.Sender, nil)
		if err != nil {
			return nil, error_types.ErrNativeTokenInsufficient
		}
		if nativeTokenBalanceWei.Cmp(big.NewInt(0)) < 0 {
			return nil, error_types.ErrNativeTokenInsufficient
		}
	}

	var tx *types.DynamicFeeTx
	if isNativeToken {
		tx = &types.DynamicFeeTx{
			To:    &args.ToAddress,
			Value: args.AmountWei.BigInt(),
		}
	}

	if !isNativeToken {
		erc20Abi, err := erc20.TokenMetaData.GetAbi()
		if err != nil {
			return nil, errors.Wrap(err, "get abi")
		}
		input, err := erc20Abi.Pack("transfer", args.ToAddress, args.AmountWei.BigInt())
		if err != nil {
			return nil, errors.Wrap(err, "abi pack")
		}
		tx = &types.DynamicFeeTx{
			To:   &args.TokenAddress,
			Data: input,
		}
	}
	return tx, nil
}

func SendTransaction(ctx context.Context, client simulated.Client, tx *types.DynamicFeeTx,
	sender common.Address, privateKey string,
) (common.Hash, error) {
	if tx.Nonce == 0 {
		nonce, err := client.NonceAt(ctx, sender, nil)
		if err != nil {
			return common.Hash{}, errors.Wrap(err, "get nonce")
		}
		tx.Nonce = nonce
	}
	if tx.GasFeeCap == nil {
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			return common.Hash{}, errors.Wrap(err, "suggest gas price")
		}
		tx.GasFeeCap = gasPrice
	}

	if tx.GasTipCap == nil {
		tip, err := client.SuggestGasTipCap(ctx)
		if err != nil {
			return common.Hash{}, errors.Wrap(err, "suggest gas tip")
		}
		tx.GasTipCap = tip
	}

	if tx.Gas == 0 {
		gas, err := client.EstimateGas(ctx, ethereum.CallMsg{
			From:  sender,
			To:    tx.To,
			Value: tx.Value,
			Data:  tx.Data,
		})
		if err != nil {
			return common.Hash{}, errors.Wrap(err, "estimate gas")
		}
		tx.Gas = gas
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "get chain id")
	}

	var signTxType SignTxType = SignTxTypeTransfer
	if ctxKey := ctx.Value(constant.SignTxKeyInCtx); ctxKey != nil {
		signTxType = ctxKey.(SignTxType)
	}
	transaction, err := SignTx(types.NewTx(tx), privateKey, chainId.Int64(), signTxType)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "sign tx")
	}
	return transaction.Hash(), client.SendTransaction(ctx, transaction)
}

// GetTransferInfo input ex: a9059cbb0000000000000000000000000350101f2cb6aa65caab7954246a56f906a3f57d0000000000000000000000000000000000000000033b2e3c9fd0803ce8000000
func GetTransferInfo(input []byte) (to common.Address, amount *big.Int, err error) {
	if len(input) < 4 {
		return common.Address{}, nil, errors.New("invalid input")
	}
	erc20Abi, err := erc20.TokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "get abi")
	}
	args, err := erc20Abi.Methods["transfer"].Inputs.Unpack(input[4:])
	if err != nil {
		return common.Address{}, nil, errors.Wrap(err, "unpack")
	}
	if len(args) != 2 {
		return common.Address{}, nil, errors.New("invalid args")
	}
	return args[0].(common.Address), args[1].(*big.Int), nil
}
