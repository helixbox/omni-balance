package base

import (
	"context"
	"fmt"
	"math/big"
	"time"

	base_portal "omni-balance/utils/enclave/router/base/portal"

	"github.com/ethereum-optimism/optimism/op-node/bindings"
	bindingspreview "github.com/ethereum-optimism/optimism/op-node/bindings/preview"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

func Withdrawer(withdrawal common.Hash) (*FPWithdrawer, error) {
	ctx := context.Background()
	l1Client, err := ethclient.DialContext(ctx, l1RPC)
	if err != nil {
		return nil, fmt.Errorf("Error dialing L1 client: %w", err)
	}
	l2Client, err := rpc.DialContext(ctx, l2RPC)
	if err != nil {
		return nil, fmt.Errorf("Error dialing L2 client: %w", err)
	}

	portal, err := bindingspreview.NewOptimismPortal2(portal, l1Client)
	if err != nil {
		return nil, fmt.Errorf("Error binding OptimismPortal2 contract: %w", err)
	}

	dgf, err := bindings.NewDisputeGameFactory(disputeGameFactory, l1Client)
	if err != nil {
		return nil, fmt.Errorf("Error binding DisputeGameFactory contract: %w", err)
	}

	return &FPWithdrawer{
		Ctx:      ctx,
		L1Client: l1Client,
		L2Client: l2Client,
		L2TxHash: withdrawal,
		Portal:   portal,
		Factory:  dgf,
		Opts:     &bind.TransactOpts{},
	}, nil
}

type FPWithdrawer struct {
	Ctx      context.Context
	L1Client *ethclient.Client
	L2Client *rpc.Client
	L2TxHash common.Hash
	Portal   *bindingspreview.OptimismPortal2
	Factory  *bindings.DisputeGameFactory
	Opts     *bind.TransactOpts
}

func (w *FPWithdrawer) CheckIfProvable() error {
	l2WithdrawalBlock, err := txBlock(w.Ctx, w.L2Client, w.L2TxHash)
	if err != nil {
		return fmt.Errorf("error querying withdrawal tx block: %w", err)
	}

	latestGame, err := withdrawals.FindLatestGame(w.Ctx, &w.Factory.DisputeGameFactoryCaller, &w.Portal.OptimismPortal2Caller)
	if err != nil {
		return fmt.Errorf("failed to find latest game: %w", err)
	}
	l2BlockNumber := new(big.Int).SetBytes(latestGame.ExtraData[0:32])

	if l2BlockNumber.Uint64() < l2WithdrawalBlock.Uint64() {
		return fmt.Errorf("the latest L2 block proposed in the DisputeGameFactory is %d and is not past L2 block %d that includes the withdrawal - the withdrawal cannot be proven yet",
			l2BlockNumber.Uint64(), l2WithdrawalBlock.Uint64())
	}
	return nil
}

func (w *FPWithdrawer) getWithdrawalHash() (common.Hash, error) {
	l2 := ethclient.NewClient(w.L2Client)
	receipt, err := l2.TransactionReceipt(w.Ctx, w.L2TxHash)
	if err != nil {
		return common.HexToHash(""), err
	}

	ev, err := withdrawals.ParseMessagePassed(receipt)
	if err != nil {
		return common.HexToHash(""), err
	}

	hash, err := withdrawals.WithdrawalHash(ev)
	if err != nil {
		return common.HexToHash(""), err
	}

	return hash, nil
}

func (w *FPWithdrawer) GetProvenWithdrawalTime() (uint64, error) {
	hash, err := w.getWithdrawalHash()
	if err != nil {
		return 0, err
	}

	// the proven withdrawal structure now contains an additional mapping, as withdrawal proofs are now stored per submitter address
	provenWithdrawal, err := w.Portal.ProvenWithdrawals(&bind.CallOpts{}, hash, w.Opts.From)
	if err != nil {
		return 0, err
	}

	return provenWithdrawal.Timestamp, nil
}

func (w *FPWithdrawer) ProveWithdrawal() ([]byte, error) {
	l2 := ethclient.NewClient(w.L2Client)
	l2g := gethclient.New(w.L2Client)

	params, err := withdrawals.ProveWithdrawalParametersFaultProofs(w.Ctx, l2g, l2, l2, w.L2TxHash, &w.Factory.DisputeGameFactoryCaller, &w.Portal.OptimismPortal2Caller)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get prove withdrawal parameters")
	}
	fmt.Printf("params: %+v\n", params)

	portalAbi, err := base_portal.BasePortalMetaData.ParseABI()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx := base_portal.TypesWithdrawalTransaction{
		Nonce:    params.Nonce,
		Sender:   params.Sender,
		Target:   params.Target,
		Value:    params.Value,
		GasLimit: params.GasLimit,
		Data:     params.Data,
	}

	packedData, err := portalAbi.Pack("proveWithdrawalTransaction", tx, params.L2OutputIndex, params.OutputRootProof, params.WithdrawalProof)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack prove withdrawal transaction")
	}

	return packedData, nil
}

func (w *FPWithdrawer) FinalizeWithdrawalData() ([]byte, error) {
	// get the WithdrawalTransaction info needed to finalize the withdrawal
	l2 := ethclient.NewClient(w.L2Client)

	// Transaction receipt
	receipt, err := l2.TransactionReceipt(w.Ctx, w.L2TxHash)
	if err != nil {
		return nil, err
	}
	// Parse the receipt
	ev, err := withdrawals.ParseMessagePassed(receipt)
	if err != nil {
		return nil, err
	}
	fmt.Printf("ev: %+v\n", ev)

	basePortal, err := base_portal.BasePortalMetaData.ParseABI()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return basePortal.Pack("finalizeWithdrawalTransaction", base_portal.TypesWithdrawalTransaction{
		Nonce:    ev.Nonce,
		Sender:   ev.Sender,
		Target:   ev.Target,
		Value:    ev.Value,
		GasLimit: ev.GasLimit,
		Data:     ev.Data,
	})
}

func (w *FPWithdrawer) IsProofFinalized() (bool, error) {
	return w.Portal.FinalizedWithdrawals(&bind.CallOpts{}, w.L2TxHash)
}

func (w *FPWithdrawer) FinalizeWithdrawal() error {
	// get the withdrawal hash
	hash, err := w.getWithdrawalHash()
	if err != nil {
		return err
	}

	// check if the withdrawal can be finalized using the calculated withdrawal hash
	err = w.Portal.CheckWithdrawal(&bind.CallOpts{}, hash, w.Opts.From)
	if err != nil {
		return err
	}

	// get the WithdrawalTransaction info needed to finalize the withdrawal
	l2 := ethclient.NewClient(w.L2Client)

	// Transaction receipt
	receipt, err := l2.TransactionReceipt(w.Ctx, w.L2TxHash)
	if err != nil {
		return err
	}
	// Parse the receipt
	ev, err := withdrawals.ParseMessagePassed(receipt)
	if err != nil {
		return err
	}

	// finalize the withdrawal
	tx, err := w.Portal.FinalizeWithdrawalTransaction(
		w.Opts,
		bindingspreview.TypesWithdrawalTransaction{
			Nonce:    ev.Nonce,
			Sender:   ev.Sender,
			Target:   ev.Target,
			Value:    ev.Value,
			GasLimit: ev.GasLimit,
			Data:     ev.Data,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Completed withdrawal for %s: %s\n", w.L2TxHash.String(), tx.Hash().String())

	// Wait 5 mins max for confirmation
	ctxWithTimeout, cancel := context.WithTimeout(w.Ctx, 5*time.Minute)
	defer cancel()
	return waitForConfirmation(ctxWithTimeout, w.L1Client, tx.Hash())
}

func txBlock(ctx context.Context, l2c *rpc.Client, l2TxHash common.Hash) (*big.Int, error) {
	l2 := ethclient.NewClient(l2c)
	// Figure out when our withdrawal was included
	receipt, err := l2.TransactionReceipt(ctx, l2TxHash)
	if err != nil {
		return nil, err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, errors.New("unsuccessful withdrawal receipt status")
	}
	return receipt.BlockNumber, nil
}

func waitForConfirmation(ctx context.Context, client *ethclient.Client, tx common.Hash) error {
	for {
		receipt, err := client.TransactionReceipt(ctx, tx)
		if err == ethereum.NotFound {
			fmt.Printf("waiting for tx confirmation\n")
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(5 * time.Second):
			}
		} else if err != nil {
			return err
		} else if receipt.Status != types.ReceiptStatusSuccessful {
			return errors.New("unsuccessful withdrawal receipt status")
		} else {
			break
		}
	}
	fmt.Printf("%s confirmed\n", tx.String())
	return nil
}
