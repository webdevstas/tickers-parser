package entities

import (
	"fmt"
	"tickers-parser/internal/types"
	"time"
)

type Exchange struct {
	ID                     uint      `gorm:"primarykey"`
	CreatedAt              time.Time `gorm:"column:createdAt"`
	UpdatedAt              time.Time `gorm:"column:updatedAt"`
	Key                    string    `json:"key" db:"key"`
	Name                   string    `json:"name,omitempty"`
	Enabled                bool      `json:"enabled,omitempty"`
	types.TickersFetchable `gorm:"-"`
}

type ExchangeTickers struct {
	Exchange Exchange
	Tickers  []types.ExchangeRawTicker
}

func (e *Exchange) FetchTickers() ([]types.ExchangeRawTicker, error) {
	if e.TickersFetchable == nil {
		return nil, fmt.Errorf("not implemented fetch tickers method for exchange %v", e.Key)
	}
	return e.TickersFetchable.FetchTickers()
}
