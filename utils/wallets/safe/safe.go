package safe

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/constant"
	"omni-balance/utils/error_types"
	"omni-balance/utils/notice"
	"omni-balance/utils/wallets"
	"strings"
	"time"
)

type Safe struct {
	conf         wallets.WalletConfig
	operatorSafe *Safe
}

func NewSafe(conf wallets.WalletConfig) *Safe {
	if conf.Operator.Address.Cmp(constant.ZeroAddress) == 0 {
		panic("operator address is empty")
	}
	if conf.Operator.Address.Cmp(conf.Address) == 0 {
		panic("operator address is same as safe address")
	}
	s := &Safe{
		conf: conf,
	}
	if conf.Operator.MultiSignType != "" {
		operator := conf.Operator
		s.operatorSafe = &Safe{
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

func (*Safe) IsSupportEip712() bool {
	return false
}

func (s *Safe) CheckFullAccess(ctx context.Context) error {
	info, err := s.safeWalletInfo(ctx)
	if err != nil {
		return err
	}
	if !utils.InArrayFold(s.conf.Operator.Address.Hex(), info.Owners) {
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
func (s *Safe) GetAddress(isReal ...bool) common.Address {
	if len(isReal) > 0 && isReal[0] {
		return s.conf.Operator.Address
	}
	return s.conf.Address
}

func (s *Safe) IsDifferentAddress() bool {
	return true
}

func (s *Safe) SignRawMessage(_ []byte) (sig []byte, err error) {
	return nil, error_types.ErrUnsupportedWalletType
}

func (s *Safe) GetExternalBalance(ctx context.Context, tokenAddress common.Address, decimals int32,
	client simulated.Client) (decimal.Decimal, error) {
	return s.GetBalance(ctx, tokenAddress, decimals, client)
}

func (s *Safe) GetBalance(ctx context.Context, tokenAddress common.Address, decimals int32,
	client simulated.Client) (decimal.Decimal, error) {
	return chains.GetTokenBalance(ctx, client, tokenAddress.Hex(), s.GetAddress().Hex(), decimals)
}

func (s *Safe) GetNonce(ctx context.Context, client simulated.Client) (uint64, error) {
	log := utils.GetLogFromCtx(ctx)
	if s.operatorSafe != nil {
		nonce, err := s.operatorSafe.nonce(ctx)
		log.Debugf("operator %s safe nonce: %d", s.conf.Operator.Address.Hex(), nonce)
		return uint64(nonce), err
	}
	return client.NonceAt(ctx, s.GetAddress(), nil)
}

func (s *Safe) SendTransaction(ctx context.Context, tx *types.LegacyTx, client simulated.Client) (common.Hash, error) {
	return s.MultisigTransaction(ctx, tx, client)
}

func (s *Safe) WaitTransaction(ctx context.Context, txHash common.Hash, _ simulated.Client) error {
	var (
		log   = utils.GetLogFromCtx(ctx).WithFields(utils.ToMap(s))
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
			if len(tx.Confirmations) < tx.ConfirmationsRequired {
				count = 0
				log.Infof("%s transaction %s confirmations: %d, required: %d,",
					tx.Safe, txHash, len(tx.Confirmations), tx.ConfirmationsRequired)
				if err := notice.Send(ctx,
					fmt.Sprintf("wait %s safeHash %s confirmations and execute.",
						constant.GetChainName(s.GetChainIdByCtx(ctx)), txHash),
					fmt.Sprintf("Please go to %s %s safe address to confirm  and execute #%d transaction.",
						constant.GetChainName(s.GetChainIdByCtx(ctx)), tx.Safe, tx.Nonce),
					logrus.WarnLevel,
				); err != nil {
					log.Debugf("send notice error: %s", err)
				}
				continue
			}
			if !tx.IsExecuted {
				count = 0
				log.Debugf("transaction %s is not executed", txHash)
				continue
			}
			if !tx.IsSuccessful {
				log.Debugf("transaction %s is not successful", txHash)
				count++
				continue
			}
			log.Debugf("transaction %s is successful", txHash)
			return nil
		}
	}
	return errors.Errorf("wait transaction %s timeout", txHash)
}

func (s *Safe) MarshalJSON() ([]byte, error) {
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

func (s *Safe) GetRealHash(ctx context.Context, txHash common.Hash, client simulated.Client) (common.Hash, error) {
	tx, err := s.GetMultiSigTransaction(ctx, txHash)
	if err != nil {
		return [32]byte{}, err
	}
	if len(tx.Confirmations) != tx.ConfirmationsRequired {
		return [32]byte{}, errors.New("transaction not confirmed")
	}
	if tx.TransactionHash == nil {
		return [32]byte{}, errors.New("transaction hash is empty")
	}
	txHashStr := cast.ToString(tx.TransactionHash)
	if txHashStr == "" {
		return [32]byte{}, errors.Errorf("transaction hash convert to hash error, the result is %+v", tx.TransactionHash)
	}
	return common.HexToHash(txHashStr), nil
}
