package exchange

import (
	"tickers-parser/internal/entities"
)

type exmo struct {
	*entities.Exchange
}

func getExmo() *exmo {
	ex := &entities.Exchange{
		Key:        "exmo",
		Name:       "Exmo",
		Enabled:    true,
		TickersUrl: "https://api.exmo.com/v1/ticker",
	}
	return &exmo{
		ex,
	}
}

var Exmo = getExmo()

type exmoTicker struct {
	BuyPrice  string `json:"buy_price,omitempty"`
	SellPrice string `json:"sell_price,omitempty"`
	LastTrade string `json:"last_trade,omitempty"`
	High      string `json:"high,omitempty"`
	Low       string `json:"low,omitempty"`
	Avg       string `json:"avg,omitempty"`
	Vol       string `json:"vol,omitempty"`
	VolCurr   string `json:"vol_curr,omitempty"`
	Updated   int64  `json:"updated,omitempty"` //sec
}
