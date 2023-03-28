package entities

import (
	"time"
)

type Ticker struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	BaseSymbol   string `json:"baseSymbol" gorm:"index:ticker_idx,unique;"`
	QuoteSymbol  string `json:"quoteSymbol" gorm:"index:ticker_idx,unique;"`
	Volume       float64
	Bid          float64
	Ask          float64
	Open         float64
	High         float64
	Low          float64
	Change       float64
	ExchangeId   uint   `json:"exchangeId" gorm:"index:ticker_idx,unique;"`
	BaseCoinId   int    `json:"baseCoinId"`
	QuoteCoinId  int    `json:"quoteCoinId"`
	BaseAddress  string `json:"baseAddress"`
	QuoteAddress string `json:"quoteAddress"`
	Enabled      bool
	Last         float64
}
