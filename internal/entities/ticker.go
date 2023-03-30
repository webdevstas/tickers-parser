package entities

import (
	"time"
)

type Ticker struct {
	ID           uint `gorm:"primaryKey"`
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
	ExchangeID   uint     `json:"exchangeId" gorm:"index:ticker_idx,unique;"`
	BaseCoinID   uint     `json:"baseCoinId" gorm:"index:base_coin_idx;"`
	QuoteCoinID  uint     `json:"quoteCoinId" gorm:"index:quote_coin_idx;"`
	Exchange     Exchange `gorm:"foreignKey:ExchangeID"`
	BaseAddress  string   `json:"baseAddress"`
	QuoteAddress string   `json:"quoteAddress"`
	Enabled      bool
	Last         float64
}

func (t *Ticker) LinkTickerToCoins(coins map[string]Coin) bool {
	baseCoin, foundBase := coins[t.BaseSymbol]
	quoteCoin, foundQuote := coins[t.QuoteSymbol]

	if foundBase {
		t.BaseCoinID = baseCoin.ID
	}

	if foundQuote {
		t.QuoteCoinID = quoteCoin.ID
	}

	if foundBase || foundQuote {
		return true
	}

	return false
}
