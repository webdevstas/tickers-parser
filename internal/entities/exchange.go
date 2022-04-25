package entities

import (
	"gorm.io/gorm"
	"tickers-parser/internal/types"
)

type IExchange interface {
	FetchTickers() ([]types.ExchangeRawTicker, error)
}

type Exchange struct {
	gorm.Model
	Key        string `json:"key" db:"key"`
	Name       string `json:"name,omitempty" db:"name"`
	Enabled    bool   `json:"enabled,omitempty" db:"enabled"`
	TickersUrl string `json:"tickersUrl" db:"tickersUrl"`
	IExchange  `gorm:"-"`
}

type ExchangeTickers struct {
	Exchange  string
	Timestamp int64
	Tickers   []types.ExchangeRawTicker
}
