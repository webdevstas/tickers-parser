package entities

import "tickers-parser/internal/types"

type Exchange struct {
	Id           int                                `json:"id,omitempty" db:"id"`
	Key          string                             `json:"key" db:"key"`
	Name         string                             `json:"name,omitempty" db:"name"`
	Enabled      bool                               `json:"enabled,omitempty" db:"enabled"`
	FetchTickers func(channels *types.ChannelsPair) `json:"fetchTickers,omitempty" db:"fetchTickers"`
}

type ExchangeTickers struct {
	Exchange  string
	Timestamp int64
	Tickers   []Ticker
}
