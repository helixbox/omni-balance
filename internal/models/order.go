package models

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	"time"
)

type OrderStatus string

const (
	OrderStatusWait OrderStatus = "wait"
	//OrderStatusProcessing               OrderStatus = "processing"
	OrderStatusSuccess OrderStatus = "success"
	//OrderStatusFail                     OrderStatus = "fail"
	OrderStatusWaitTransferFromOperator OrderStatus = "wait_transfer_from_operator"
	OrderStatusWaitCrossChain           OrderStatus = "wait_cross_chain"
	OrderStatusUnknown                  OrderStatus = "unknown"
)

type Order struct {
	gorm.Model
	// 唯一索引
	Wallet           string                        `json:"wallet" gorm:"type:varchar(64)"`
	TokenInName      string                        `json:"token_in_name"`
	TokenOutName     string                        `json:"token_out_name" gorm:"type:varchar(64)"`
	SourceChainName  string                        `json:"source_chain_name"`
	TargetChainName  string                        `json:"target_chain_name"`
	CurrentChainName string                        `json:"current_chain_name" gorm:"type:varchar(64)"`
	CurrentBalance   decimal.Decimal               `json:"current_balance" gorm:"type:decimal(32,16); default:0"`
	Amount           decimal.Decimal               `json:"amount" gorm:"type:decimal(32,16); default:0"`
	IsLock           bool                          `json:"is_lock" gore:"type:boolean; default:false"`
	LockTime         int64                         `json:"lock_time" gorm:"default:0"`
	Status           OrderStatus                   `json:"status" gorm:"type:int; default:0;index"`
	ProviderType     configs.LiquidityProviderType `json:"provider_type" gorm:"type:varchar(64)"`
	ProviderName     string                        `json:"provider_name" gorm:"type:varchar(64)"`
	ProviderOrderId  string                        `json:"order_id" gorm:"type:varchar(64)"`
	Tx               string                        `json:"tx" gorm:"type:varchar(64)"`
	Order            *json.RawMessage              `json:"order" gorm:"type:json;default:null"`
	Error            string                        `json:"error" gorm:"type:varchar(255)"`
}

type OrderProcess struct {
	gorm.Model
	// Order.ID 一对多关联
	OrderId          uint                          `json:"order_id" gorm:"type:int;index"`
	Error            string                        `json:"error"`
	ProviderType     configs.LiquidityProviderType `json:"type"`
	ProviderName     string                        `json:"provider_name"`
	CurrentChainName string                        `json:"current_chain_name"`
	Status           string                        `json:"status"`
	Action           string                        `json:"action"`
	Amount           decimal.Decimal               `json:"amount"`
	// Tx is the transaction hash
	Tx string `json:"tx"`
}

func (o *Order) UnLock(db *gorm.DB) {
	db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{"is_lock": false, "lock_time": 0})
}

func (o *Order) Lock(db *gorm.DB) bool {
	var order Order
	db.Where("id = ?", o.ID).First(&order)
	if order.IsLock && time.Unix(order.LockTime, 0).Add(time.Hour).Unix() < time.Now().Unix() {
		return true
	}
	db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{"is_lock": 1, "lock_time": time.Now().Unix()})
	return false
}

func (o *Order) SaveProvider(db *gorm.DB, provider configs.LiquidityProviderType, providerName string) error {
	return db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{"provider_type": provider, "provider_name": providerName}).Error
}

func (o *Order) Success(db *gorm.DB, tx string, order interface{}, balance decimal.Decimal) error {
	data, _ := json.Marshal(order)
	return db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{
		"status":             OrderStatusSuccess,
		"tx":                 tx,
		"current_chain_name": o.TargetChainName,
		"order":              json.RawMessage(data),
		"current_balance":    balance,
	}).Error
}

func GetLastOrderProcess(ctx context.Context, db *gorm.DB, orderId uint) OrderProcess {
	var result OrderProcess
	_ = db.WithContext(ctx).Where("order_id = ?", orderId).Order("id desc").First(&result)
	return result
}

func (o *Order) GetLogs() *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"order": utils.ToMap(o),
	})
}
