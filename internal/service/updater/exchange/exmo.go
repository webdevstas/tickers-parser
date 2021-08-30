package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/utils"
	"time"
)

var exmoKey = "exmo"

var Exmo = entities.Exchange{
	Key:          exmoKey,
	Name:         "Exmo",
	Enabled:      true,
	FetchTickers: exmoFetchTickers,
}

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

func exmoFetchTickers() (entities.ExchangeTickers, error) {
	var tickersArr []entities.Ticker
	apiUrl := "https://api.exmo.com/v1/ticker"
	rawTickers := make(map[string]exmoTicker)
	err := utils.FetchJson(apiUrl, &rawTickers)

	if err != nil {
		return entities.ExchangeTickers{}, err
	}

	for pair, data := range rawTickers {
		symbols := strings.Split(pair, "_")
		volume, _ := strconv.ParseFloat(data.Vol, 64)
		bid, _ := strconv.ParseFloat(data.SellPrice, 64)
		ask, _ := strconv.ParseFloat(data.BuyPrice, 64)
		high, _ := strconv.ParseFloat(data.High, 64)
		low, _ := strconv.ParseFloat(data.Vol, 64)
		ticker := entities.Ticker{
			BaseSymbol:  symbols[0],
			QuoteSymbol: symbols[1],
			Volume:      volume,
			Bid:         bid,
			Ask:         ask,
			High:        high,
			Low:         low,
			UpdatedAt:   data.Updated,
		}
		tickersArr = append(tickersArr, ticker)
	}
	res := entities.ExchangeTickers{
		Exchange:  exmoKey,
		Timestamp: time.Now().Unix(),
		Tickers:   tickersArr,
	}

	return res, nil
}
