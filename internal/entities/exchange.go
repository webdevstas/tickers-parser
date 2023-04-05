package entities

import (
	"fmt"
	http_client "tickers-parser/internal/services/http-client"
	"time"
)

type TickersFetchable interface {
	FetchTickers(http *http_client.HttpClient) ([]ExchangeRawTicker, error)
}

type Exchange struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Key              string `json:"key" db:"key"`
	Name             string `json:"name,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	TickersFetchable `gorm:"-"`
	FetchTimeout     uint `json:"-"`
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

func (e *Exchange) FetchTickers(http *http_client.HttpClient) ([]ExchangeRawTicker, error) {
	if e.TickersFetchable == nil {
		return nil, fmt.Errorf("not implemented fetch tickers method for exchange %v", e.Key)
	}
	return e.TickersFetchable.FetchTickers(http)
}
