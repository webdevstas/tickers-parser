package entities

import (
	"gorm.io/gorm"
	"tickers-parser/internal/utils"
	"time"
)

type IExchange interface {
	FetchRawTickers() (ExchangeRawTickers, error)
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
	Tickers   []Ticker
}
type ExchangeRawTickers struct {
	Exchange   string
	Timestamp  int64
	RawTickers map[string]interface{}
}

func (e *Exchange) FetchRawTickers() (ExchangeRawTickers, error) {
	rawTickers := make(map[string]interface{})
	err := utils.FetchJson(e.TickersUrl, &rawTickers)
	if err != nil {
		return ExchangeRawTickers{}, err
	}
	return ExchangeRawTickers{
		Exchange:   e.Key,
		Timestamp:  time.Now().Unix(),
		RawTickers: rawTickers,
	}, nil
}
