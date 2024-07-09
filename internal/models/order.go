package models

import (
	"context"
	"encoding/json"
	"omni-balance/utils/configs"
	"omni-balance/utils/provider"
	"sync"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	OrderStatusWaitTransferFromOperator provider.TxStatus = "wait_transfer_from_operator"
	OrderStatusWaitCrossChain           provider.TxStatus = "wait_cross_chain"
)

var (
	orderLocker sync.Map
)

type Order struct {
	gorm.Model
	// 唯一索引
	Wallet           string               `json:"wallet" gorm:"type:varchar(64)"`
	TokenInName      string               `json:"token_in_name"`
	TokenOutName     string               `json:"token_out_name" gorm:"type:varchar(64)"`
	SourceChainName  string               `json:"source_chain_name"`
	TargetChainName  string               `json:"target_chain_name"`
	CurrentChainName string               `json:"current_chain_name" gorm:"type:varchar(64)"`
	CurrentBalance   decimal.Decimal      `json:"current_balance" gorm:"type:decimal(32,16); default:0"`
	Amount           decimal.Decimal      `json:"amount" gorm:"type:decimal(32,16); default:0"`
	IsLock           bool                 `json:"is_lock" gore:"type:boolean; default:false"`
	LockTime         int64                `json:"lock_time" gorm:"default:0"`
	Status           provider.TxStatus    `json:"status" gorm:"type:varchar(32); default:'pending';index"`
	ProviderType     configs.ProviderType `json:"provider_type" gorm:"type:varchar(64)"`
	ProviderName     string               `json:"provider_name" gorm:"type:varchar(64)"`
	ProviderOrderId  string               `json:"order_id" gorm:"type:varchar(64)"`
	Tx               string               `json:"tx" gorm:"type:varchar(64)"`
	Order            datatypes.JSON       `json:"order" gorm:"default:null"`
	Error            string               `json:"error" gorm:"type:varchar(255)"`
	TaskId           string               `json:"task_id" gorm:"type:varchar(64)"`
	ProcessType      string               `json:"process_type"`
	Remark           string               `json:"remark" grom:"type:varchar(32)"`
}

type Tasks struct {
	Id           string
	ProviderType string  `json:"type"`
	Orders       []Order `json:"orders"`
}

type OrderProcess struct {
	gorm.Model
	// Order.ID 一对多关联
	OrderId          uint                 `json:"order_id" gorm:"type:int;index"`
	Error            string               `json:"error"`
	ProviderType     configs.ProviderType `json:"type"`
	ProviderName     string               `json:"provider_name"`
	CurrentChainName string               `json:"current_chain_name"`
	Status           string               `json:"status"`
	Action           string               `json:"action"`
	Amount           decimal.Decimal      `json:"amount"`
	// Tx is the transaction hash
	Tx string `json:"tx"`
}

func (o *Order) UnLock(db *gorm.DB) {
	orderLocker.Delete(o.ID)
}

func (o *Order) HasLocked() bool {
	_, ok := orderLocker.Load(o.ID)
	return ok
}

func (o *Order) Lock(db *gorm.DB) bool {
	if o.HasLocked() {
		return true
	}
	orderLocker.Store(o.ID, struct{}{})
	return false
}

func (o *Order) SaveProvider(db *gorm.DB, provider configs.ProviderType, providerName string) error {
	return db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{"provider_type": provider, "provider_name": providerName}).Error
}

func (o *Order) Success(db *gorm.DB, tx string, order interface{}, balance decimal.Decimal) error {
	data, _ := json.Marshal(order)
	return db.Model(&Order{}).Where("id = ?", o.ID).Updates(map[string]interface{}{
		"status":             provider.TxStatusSuccess,
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

func GetOrder(ctx context.Context, db *gorm.DB, orderId uint) Order {
	var result = new(Order)
	_ = db.WithContext(ctx).Where("id = ?", orderId).First(&result)
	if result.ID == 0 {
		return Order{}
	}
	return *result
}

func ListOrdersByTaskId(ctx context.Context, db *gorm.DB, taskId string) ([]Order, error) {
	var result []Order
	err := db.WithContext(ctx).Where("task_id = ?", taskId).Find(&result).Error
	if err != nil {
		return nil, errors.Wrap(err, "find buy tokens error")
	}

	return result, nil
}

func ListNotSuccessTasks(ctx context.Context, db *gorm.DB, isInclude func(order Order) bool) ([]Tasks, error) {
	var (
		tasks  = make(map[string][]Order)
		result []Tasks
		orders []Order
	)
	err := db.WithContext(ctx).Where("status != ?", provider.TxStatusSuccess).Find(&orders).Error
	if err != nil {
		return nil, errors.Wrap(err, "find buy tokens error")
	}
	for index, v := range orders {
		if isInclude != nil && !isInclude(v) {
			continue
		}
		tasks[v.TaskId] = append(tasks[v.TaskId], orders[index])
	}

	for k := range tasks {
		result = append(result, Tasks{
			Id:           k,
			ProviderType: tasks[k][0].ProcessType,
			Orders:       tasks[k],
		})
	}
	return result, nil
}

func (o *Order) GetLogs() *logrus.Entry {
	fields := map[string]interface{}{
		"orderId":         o.ID,
		"TargetChainName": o.TargetChainName,
		"TokenOutName":    o.TokenOutName,
		"TaskId":          o.TaskId,
		"Amout":           o.Amount,
	}
	if o.ProviderType != "" {
		fields["ProviderType"] = o.ProviderType
	}
	if o.ProviderName != "" {
		fields["ProviderName"] = o.ProviderName
	}
	if o.TokenInName != "" {
		fields["TokenInName"] = o.TokenInName
	}
	if o.SourceChainName != "" {
		fields["SourceChainName"] = o.SourceChainName
	}
	return logrus.WithFields(logrus.Fields{
		"order": fields,
	})
}
