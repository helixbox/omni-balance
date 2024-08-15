package dsafe

import (
	"context"
	"encoding/json"
	"fmt"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/notice"
	"omni-balance/utils/wallets"
	"strings"
	"time"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
)

type Dsafe struct {
	conf         wallets.WalletConfig
	operatorSafe *Dsafe
}

func NewDsafe(conf wallets.WalletConfig) *Dsafe {
	if conf.Operator.Address.Cmp(constant.ZeroAddress) == 0 {
		panic("operator address is empty")
	}
	if conf.Operator.Address.Cmp(conf.Address) == 0 {
		panic("operator address is same as safe address")
	}
	s := &Dsafe{
		conf: conf,
	}
	if conf.Operator.MultiSignType != "" {
		operator := conf.Operator
		s.operatorSafe = &Dsafe{
			conf: wallets.WalletConfig{
				PrivateKey: operator.PrivateKey,
				Address:    operator.Address,
				Operator: wallets.Operator{
					Address:    operator.Operator,
					PrivateKey: operator.PrivateKey,
				},
				MultiSignType: operator.MultiSignType,
			},
		}
	}
	return s
}

func (*Dsafe) IsSupportEip712() bool {
	return false
}

func (s *Dsafe) CheckFullAccess(ctx context.Context) error {
	info, err := s.safeWalletInfo(ctx)
	if err != nil {
		return err
	}
	var owners []string
	for _, v := range info.Owners {
		owners = append(owners, v.Value)
	}
	if !utils.InArrayFold(s.conf.Operator.Address.Hex(), owners) {
		return errors.New("operator is not in owners")
	}
	return nil
}

// GetAddress retrieves the address from the security configuration.
// If isReal is provided as true and the multisig type is set,
// it returns the operator's address; otherwise, it returns the default address.
// Parameters:
//
//	isReal - An optional boolean indicating whether to return the real address.
//
// Returns:
//
//	common.Address - The address from the configuration.
func (s *Dsafe) GetAddress(isReal ...bool) common.Address {
	if len(isReal) > 0 && isReal[0] {
		if s.conf.Operator.MultiSignType == "" {
			return s.conf.Address
		}
		return s.conf.Operator.Address
	}
	return s.conf.Address
}

func (s *Dsafe) IsDifferentAddress() bool {
	return true
}

func (s *Dsafe) SignRawMessage(_ []byte) (sig []byte, err error) {
	return nil, error_types.ErrUnsupportedWalletType
}

func (s *Dsafe) GetExternalBalance(ctx context.Context, tokenAddress common.Address, decimals int32,
	client simulated.Client) (decimal.Decimal, error) {
	return s.GetBalance(ctx, tokenAddress, decimals, client)
}

func (s *Dsafe) GetBalance(ctx context.Context, tokenAddress common.Address, decimals int32,
	client simulated.Client) (decimal.Decimal, error) {
	return chains.GetTokenBalance(ctx, client, tokenAddress.Hex(), s.GetAddress().Hex(), decimals)
}

func (s *Dsafe) GetNonce(ctx context.Context, client simulated.Client) (uint64, error) {
	if s.operatorSafe != nil {
		nonce, err := s.operatorSafe.nonce(ctx)
		log.Debugf("operator %s safe nonce: %d", s.conf.Operator.Address.Hex(), nonce)
		return uint64(nonce), err
	}
	return client.NonceAt(ctx, s.GetAddress(), nil)
}

func (s *Dsafe) SendTransaction(ctx context.Context, tx *types.LegacyTx, client simulated.Client) (common.Hash, error) {
	return s.MultisigTransaction(ctx, tx, client)
}

func (s *Dsafe) WaitTransaction(ctx context.Context, txHash common.Hash, _ simulated.Client) error {
	var (
		t     = time.NewTicker(time.Second * 10)
		count = 0
	)
	defer t.Stop()
	for count < 60 {
		select {
		case <-ctx.Done():
			return context.Canceled
		case <-t.C:
			tx, err := s.GetMultiSigTransaction(ctx, txHash)
			if err != nil && strings.Contains(err.Error(), "No MultisigTransaction matches the given query") {
				return errors.Wrap(err, "transaction not found")
			}
			if err != nil {
				log.Debugf("wait transaction %s error: %s", txHash, err)
				continue
			}
			if len(tx.DetailedExecutionInfo.Confirmations) < tx.DetailedExecutionInfo.ConfirmationsRequired {
				count = 0
				log.Infof("%s transaction %s confirmations: %d, required: %d,",
					tx.SafeAddress, txHash, len(tx.DetailedExecutionInfo.Confirmations), tx.DetailedExecutionInfo.ConfirmationsRequired)
				if err := notice.Send(ctx,
					fmt.Sprintf("wait %s safeHash %s confirmations and execute.",
						constant.GetChainName(s.GetChainIdByCtx(ctx)), txHash),
					fmt.Sprintf("Please go to https://dsafe.dcdao.box/transactions/queue?safe=darwinia:%s  to confirm  and execute #%d transaction.",
						tx.SafeAddress, tx.DetailedExecutionInfo.Nonce),
					zapcore.WarnLevel,
				); err != nil {
					log.Debugf("send notice error: %s", err)
				}
				continue
			}
			if tx.ExecutedAt == nil || *tx.ExecutedAt <= 0 {
				count = 0
				log.Debugf("transaction %s is not executed", txHash)
				continue
			}
			if !strings.EqualFold(tx.TxStatus, "SUCCESS") {
				log.Debugf("transaction %s status is %s", txHash, tx.TxStatus)
				count++
				continue
			}
			log.Debugf("transaction %s is successful", txHash)
			return nil
		}
	}
	return errors.Errorf("wait transaction %s timeout", txHash)
}

func (s *Dsafe) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"address": s.conf.Address.Hex(),
		"operator": map[string]interface{}{
			"address":         s.conf.Operator.Address.Hex(),
			"operator":        s.conf.Operator.Operator.Hex(),
			"multi_sign_type": s.conf.MultiSignType,
		},
		"multi_sign_type": s.conf.MultiSignType,
	})
}

func (s *Dsafe) GetRealHash(ctx context.Context, txHash common.Hash, client simulated.Client) (common.Hash, error) {
	tx, err := s.GetMultiSigTransaction(ctx, txHash)
	if err != nil {
		return [32]byte{}, err
	}
	if len(tx.DetailedExecutionInfo.Confirmations) != tx.DetailedExecutionInfo.ConfirmationsRequired {
		return [32]byte{}, errors.New("transaction not confirmed")
	}
	if tx.TxHash == "" {
		return [32]byte{}, errors.New("transaction hash is empty")
	}
	txHashStr := cast.ToString(tx.TxHash)
	if txHashStr == "" {
		return [32]byte{}, errors.Errorf("transaction hash convert to hash error, the result is %+v", tx.TxHash)
	}
	return common.HexToHash(txHashStr), nil
}

func (s *Dsafe) Name(_ context.Context) string {
	return "dsafe"
}
