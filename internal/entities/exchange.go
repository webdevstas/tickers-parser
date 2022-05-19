package entities

import (
	"fmt"
	"time"
)

type TickersFetchable interface {
	FetchTickers() ([]ExchangeRawTicker, error)
}

type Exchange struct {
	ID               uint      `gorm:"primarykey"`
	CreatedAt        time.Time `gorm:"column:createdAt"`
	UpdatedAt        time.Time `gorm:"column:updatedAt"`
	Key              string    `json:"key" db:"key"`
	Name             string    `json:"name,omitempty"`
	Enabled          bool      `json:"enabled,omitempty"`
	TickersSavedAt   time.Time `gorm:"column:tickersSavedAt"`
	TickersFetchable `gorm:"-"`
}

type ExchangeTickers struct {
	Exchange Exchange
	Tickers  []ExchangeRawTicker
}

type ExchangeRawTicker struct {
	BaseSymbol   string
	QuoteSymbol  string
	Volume       float64
	Bid          float64
	Ask          float64
	High         float64
	Low          float64
	Change       float64
	BaseAddress  string
	QuoteAddress string
	Last         float64
	Open         float64
}

func (e *Exchange) FetchTickers() ([]ExchangeRawTicker, error) {
	if e.TickersFetchable == nil {
		return nil, fmt.Errorf("not implemented fetch tickers method for exchange %v", e.Key)
	}
	return e.TickersFetchable.FetchTickers()
}
