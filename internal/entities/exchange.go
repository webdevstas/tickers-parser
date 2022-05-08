package entities

import (
	"fmt"
	"gorm.io/gorm"
	"tickers-parser/internal/types"
)

type IExchange interface {
	FetchTickers() ([]types.ExchangeRawTicker, error)
}

type Exchange struct {
	gorm.Model
	Key       string `json:"key" db:"key"`
	Name      string `json:"name,omitempty" db:"name"`
	Enabled   bool   `json:"enabled,omitempty" db:"enabled"`
	IExchange `gorm:"-"`
}

type ExchangeTickers struct {
	Exchange  string
	Timestamp int64
	Tickers   []types.ExchangeRawTicker
}

func (e *Exchange) FetchTickers() ([]types.ExchangeRawTicker, error) {
	if e.IExchange == nil {
		return nil, fmt.Errorf("not implemented fetch tickers method for exchange %v", e.Key)
	}
	return e.IExchange.FetchTickers()
}
