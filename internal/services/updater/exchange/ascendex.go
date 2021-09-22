package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/utils"
	"time"
)

var Ascendex = entities.Exchange{
	Name:         "Ascendex",
	Key:          "ascendex",
	Enabled:      true,
	FetchTickers: fetchTickers,
}

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

type ascendexResponse struct {
	Code int
	Data []ascendexTicker
}

func fetchTickers() (entities.ExchangeTickers, error) {
	var tickersArr []entities.Ticker
	apiUrl := "https://ascendex.com/api/pro/v1/ticker"
	res := ascendexResponse{}
	err := utils.FetchJson(apiUrl, &res)
	if err != nil {
		return entities.ExchangeTickers{}, err
	}
	rawTickers := res.Data
	for _, rawTicker := range rawTickers {
		symbols := strings.Split(rawTicker.Symbol, "/")
		if len(symbols) < 2 {
			symbols = strings.Split(rawTicker.Symbol, "-")
		}
		open, _ := strconv.ParseFloat(rawTicker.Open, 64)
		last, _ := strconv.ParseFloat(rawTicker.Close, 64)
		high, _ := strconv.ParseFloat(rawTicker.High, 64)
		low, _ := strconv.ParseFloat(rawTicker.Low, 64)
		vol, _ := strconv.ParseFloat(rawTicker.Volume, 64)
		bid, _ := strconv.ParseFloat(rawTicker.Bid[0], 64)
		ask, _ := strconv.ParseFloat(rawTicker.Ask[0], 64)
		ticker := entities.Ticker{
			From:   symbols[0],
			To:     symbols[1],
			Open:   open,
			Last:   last,
			High:   high,
			Low:    low,
			Volume: vol,
			Bid:    bid,
			Ask:    ask,
		}
		tickersArr = append(tickersArr, ticker)
	}
	result := entities.ExchangeTickers{
		Exchange:  "ascendex",
		Timestamp: time.Now().Unix(),
		Tickers:   tickersArr,
	}
	return result, nil
}
