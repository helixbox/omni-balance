package monitor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"omni-balance/internal/daemons"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
	"omni-balance/utils/wallet_monitor"
	"time"
)

func init() {
	daemons.RegisterIntervalTask(daemons.Task{
		Name:            "monitor_wallet_balance",
		Description:     "Responsible for monitoring the balance of specified tokens in the wallet.",
		TaskFunc:        Run,
		DefaultInterval: time.Minute * 3,
	})
}

func Run(ctx context.Context, conf configs.Config) error {
	existBuyTokens, err := getExistingBuyTokens()
	if err != nil {
		return errors.Wrap(err, "find buy tokens error")
	}

	ignoreTokens := createIgnoreTokens(existBuyTokens)
	result, err := wallet_monitor.NewMonitor(conf).Check(ctx, ignoreTokens...)
	if err != nil {
		return errors.Wrap(err, "check wallet error")
	}

	if err := createOrder(result, conf); err != nil {
		return errors.Wrap(err, "create buy tokens error")
	}
	return nil
}

func getExistingBuyTokens() ([]*models.Order, error) {
	var existBuyTokens []*models.Order
	err := db.DB().Where("status != ? ", provider.TxStatusSuccess).Find(&existBuyTokens).Error
	if err != nil {
		return nil, err
	}
	return existBuyTokens, nil
}

func createIgnoreTokens(existBuyTokens []*models.Order) []wallet_monitor.IgnoreToken {
	var ignoreTokens []wallet_monitor.IgnoreToken
	for _, v := range existBuyTokens {
		ignoreTokens = append(ignoreTokens, wallet_monitor.IgnoreToken{
			Name:    v.TokenOutName,
			Chain:   v.TargetChainName,
			Address: v.Wallet,
		})
	}
	return ignoreTokens
}

func createOrder(result []wallet_monitor.Result, conf configs.Config) error {
	if len(result) == 0 {
		return nil
	}
	var orders []*models.Order
	for _, r := range result {
		for _, token := range r.Tokens {
			for _, v := range token.Chains {
				threshold := conf.GetTokenThreshold(r.Wallet, token.Name, v.ChainName)
				if v.TokenBalance.Add(v.Amount).LessThanOrEqual(threshold) {
					v.Amount = threshold.Add(v.Amount).Sub(v.TokenBalance)
					logrus.WithFields(logrus.Fields{
						"wallet":        r.Wallet,
						"token":         token.Name,
						"chain":         v.ChainName,
						"threshold":     threshold,
						"token_balance": v.TokenBalance,
					}).Infof("The amount was set too small, "+
						"and another rebalance would still be required after this one. "+
						"Therefore, the amount for this rebalance is set to %s.", v.Amount)
				}
				orders = append(orders, &models.Order{
					Wallet:          r.Wallet,
					TokenOutName:    token.Name,
					TargetChainName: v.ChainName,
					CurrentBalance:  v.TokenBalance,
					Amount:          v.Amount,
					Status:          provider.TxStatusPending,
				})
			}
		}
	}
	if err := db.DB().CreateInBatches(orders, 100).Error; err != nil {
		return err
	}
	return nil
}
