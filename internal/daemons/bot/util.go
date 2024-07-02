package bot

import (
	uuid "github.com/satori/go.uuid"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/bot"
	"omni-balance/utils/provider"
)

func getExistingBuyTokens() ([]*models.Order, error) {
	var existBuyTokens []*models.Order
	err := db.DB().Where("status != ? ", provider.TxStatusSuccess).Find(&existBuyTokens).Error
	if err != nil {
		return nil, err
	}
	return existBuyTokens, nil
}

func createIgnoreTokens(existBuyTokens []*models.Order) IgnoreTokens {
	var ignoreTokens []IgnoreToken
	for _, v := range existBuyTokens {
		ignoreTokens = append(ignoreTokens, IgnoreToken{
			Name:    v.TokenOutName,
			Chain:   v.TargetChainName,
			Address: v.Wallet,
		})
	}
	return ignoreTokens
}

func createOrder(tasks []bot.Task, processType bot.ProcessType) (orders []*models.Order, taskId string, err error) {
	if len(tasks) == 0 {
		return
	}
	var (
		txn = db.DB().Begin()
	)

	taskId = uuid.NewV4().String()

	for _, v := range tasks {
		o := &models.Order{
			Wallet:           v.Wallet,
			TokenInName:      v.TokenInName,
			TokenOutName:     v.TokenOutName,
			SourceChainName:  v.TokenInChainName,
			TargetChainName:  v.TokenOutChainName,
			CurrentChainName: v.CurrentChainName,
			Amount:           v.Amount,
			Status:           v.Status,
			ProviderType:     v.ProviderType,
			ProviderName:     v.ProviderName,
			Order:            utils.Object2JsonRawMessage(v.Order),
			TaskId:           taskId,
			ProcessType:      string(processType),
		}

		if err = txn.Create(o).Error; err != nil {
			txn.Rollback()
			return
		}
		orders = append(orders, o)
	}
	if err = txn.Commit().Error; err != nil {
		txn.Rollback()
		return
	}
	return
}
