package models

import (
	"context"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type TokenPrices []TokenPrice

type TokenPrice struct {
	gorm.Model
	TokenName string          `json:"token_name" gorm:"type:varchar(64);uniqueIndex:TokenName_Source_Chain"`
	Price     decimal.Decimal `json:"price" gorm:"type:decimal(32,16); default:0"`
	Source    string          `json:"source" gorm:"type:varchar(64);uniqueIndex:TokenName_Source_Chain"`
	Chain     string          `json:"chain" gorm:"type:varchar(64);uniqueIndex:TokenName_Source_Chain"`
}

func (t *TokenPrice) Save(db *gorm.DB) error {
	queryDb := db.Model(&TokenPrices{})
	var tp TokenPrice
	queryDb = queryDb.Where("token_name = ? and source = ? and chain = ?", t.TokenName, t.Source, t.Chain)
	queryDb.First(&tp)
	if tp.ID == 0 {
		return db.Create(&t).Error
	}
	return db.Model(&TokenPrices{}).Where("id = ?", tp.ID).Updates(t).Error
}

func (tokenPrices TokenPrices) AveragePrice() decimal.Decimal {
	if len(tokenPrices) == 0 {
		return decimal.Zero
	}
	var total decimal.Decimal
	for _, v := range tokenPrices {
		total = total.Add(v.Price)
	}
	return total.Div(decimal.NewFromInt(int64(len(tokenPrices))))
}

func (tokenPrices TokenPrices) AveragePriceOnChain(chain string) decimal.Decimal {
	if len(tokenPrices) == 0 || chain == "" {
		return decimal.Zero
	}
	var total decimal.Decimal
	for _, v := range tokenPrices {
		if v.Chain == chain {
			total = total.Add(v.Price)
		}
	}
	return total.Div(decimal.NewFromInt(int64(len(tokenPrices))))
}

func FindTokenPrice(db *gorm.DB, tokenName []string) (result map[string]decimal.Decimal, err error) {
	var (
		tokenPrices     []TokenPrice
		tokenName2Price = make(map[string]TokenPrices)
	)
	err = db.Where("token_name in (?)", tokenName).Find(&tokenPrices).Error
	if err != nil {
		return
	}
	for index, v := range tokenPrices {
		tokenName2Price[v.TokenName] = append(tokenName2Price[v.TokenName], tokenPrices[index])
	}
	result = make(map[string]decimal.Decimal)
	for _, v := range tokenName {
		if len(tokenName2Price[v]) == 0 {
			continue
		}
		result[v] = tokenName2Price[v].AveragePrice()
	}
	return result, nil
}

func SaveTokenPrice(_ context.Context, txDb *gorm.DB, price []TokenPrice) error {
	for _, v := range price {
		if err := v.Save(txDb); err != nil {
			txDb.Rollback()
			return err
		}
	}
	return txDb.Commit().Error
}
