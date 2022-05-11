package entities

import (
	"time"
)

type Ticker struct {
	ID           uint      `gorm:"primarykey"`
	CreatedAt    time.Time `gorm:"column: createdAt"`
	UpdatedAt    time.Time `gorm:"column: updatedAt"`
	BaseSymbol   string    `json:"baseSymbol" gorm:"index:ticker_idx,unique; column:baseSymbol"`
	QuoteSymbol  string    `json:"quoteSymbol" gorm:"index:ticker_idx,unique; column:quoteSymbol"`
	Volume       float64
	Bid          float64
	Ask          float64
	Open         float64
	High         float64
	Low          float64
	Change       float64
	ExchangeId   uint   `json:"exchangeId" gorm:"index:ticker_idx,unique; column:exchangeId"`
	BaseCoinId   int    `json:"baseCoinId" gorm:"column:baseCoinId"`
	QuoteCoinId  int    `json:"quoteCoinId" gorm:"column:quoteCoinId"`
	BaseAddress  string `json:"baseAddress" gorm:"column:baseAddress"`
	QuoteAddress string `json:"quoteAddress" gorm:"column:quoteAddress"`
	Enabled      bool
	Last         float64
}
