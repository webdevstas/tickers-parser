package exchange

import (
	"tickers-parser/internal/entities"
)

type ascendex struct {
	*entities.Exchange
}

func getAscendex() *ascendex {
	ex := &entities.Exchange{
		Key:        "ascendex",
		Name:       "Ascendex",
		Enabled:    true,
		TickersUrl: "https://ascendex.com/api/pro/v1/ticker",
	}
	return &ascendex{
		ex,
	}
}

var Ascendex = getAscendex()

type ascendexTicker struct {
	Symbol     string    `json:"symbol,omitempty"`
	Open       string    `json:"open,omitempty"`
	Close      string    `json:"close,omitempty"`
	High       string    `json:"high,omitempty"`
	Low        string    `json:"low,omitempty"`
	Volume     string    `json:"volume,omitempty"`
	Ask        [2]string `json:"ask,omitempty"`
	Bid        [2]string `json:"bid,omitempty"`
	TickerType string    `json:"type,omitempty"`
}
