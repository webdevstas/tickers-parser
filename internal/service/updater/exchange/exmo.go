package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/utils"
	"time"
)

type exmoTicker struct {
	Buy_price  string `json:"buy_price,omitempty"`
	Sell_price string `json:"sell_price,omitempty"`
	Last_trade string `json:"last_trade,omitempty"`
	High       string `json:"high,omitempty"`
	Low        string `json:"low,omitempty"`
	Avg        string `json:"avg,omitempty"`
	Vol        string `json:"vol,omitempty"`
	Vol_curr   string `json:"vol_curr,omitempty"`
	Updated    int    `json:"updated,omitempty"` //sec
}

func ExmoExchange() entities.Exchange {
	exmo := entities.Exchange{
		Key:          "exmo",
		Name:         "Exmo",
		Enabled:      true,
		FetchTickers: fetchTickers,
	}
	return exmo
}

func fetchTickers(dataChannel chan<- map[string]entities.ExchangeTickers, cancelChannel chan<- error) {
	var tickersArr []entities.Ticker
	apiUrl := "https://api.exmo.com/v1/ticker"
	rawTickers := make(map[string]exmoTicker)
	err := utils.FetchJson(apiUrl, &rawTickers)

	if err != nil {
		cancelChannel <- err
		return
	}

	for key, val := range rawTickers {
		symbols := strings.Split(key, "_")
		volume, _ := strconv.ParseFloat(val.Vol, 64)
		bid, _ := strconv.ParseFloat(val.Sell_price, 64)
		ask, _ := strconv.ParseFloat(val.Buy_price, 64)
		high, _ := strconv.ParseFloat(val.High, 64)
		low, _ := strconv.ParseFloat(val.Vol, 64)
		ticker := entities.Ticker{
			BaseSymbol:  symbols[0],
			QuoteSymbol: symbols[1],
			Volume:      volume,
			Bid:         bid,
			Ask:         ask,
			High:        high,
			Low:         low,
			UpdatedAt:   val.Updated,
		}
		tickersArr = append(tickersArr, ticker)
	}
	exchangeTickers := make(map[string]entities.ExchangeTickers)
	exchangeTickers["exmo"] = entities.ExchangeTickers{
		Timestamp: time.Now().Unix(),
		Tickers:   tickersArr,
	}

	dataChannel <- exchangeTickers
}
