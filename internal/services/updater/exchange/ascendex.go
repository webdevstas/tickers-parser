package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	http_client "tickers-parser/internal/services/http-client"
)

type ascendex struct {
	entities.Exchange
}

func GetAscendex() ascendex {
	return ascendex{}
}

type ascendexTicker struct {
	Symbol string    `json:"symbol,omitempty"`
	Open   string    `json:"open,omitempty"`
	Close  string    `json:"close,omitempty"`
	High   string    `json:"high,omitempty"`
	Low    string    `json:"low,omitempty"`
	Volume string    `json:"volume,omitempty"`
	Ask    [2]string `json:"ask"`
	Bid    [2]string `json:"bid"`
	Type   string    `json:"type,omitempty"`
}

type ascendexResponse struct {
	Code int              `json:"code"`
	Data []ascendexTicker `json:"data"`
}

func (a ascendex) FetchTickers(http *http_client.HttpClient) ([]entities.ExchangeRawTicker, error) {
	tickersUrl := "https://ascendex.com/api/pro/v1/ticker"
	var response ascendexResponse
	err := http.FetchJson(tickersUrl, &response)

	if err != nil {
		return nil, err
	}

	rawTickers := response.Data
	result := make([]entities.ExchangeRawTicker, 0, len(rawTickers))

	for _, ticker := range rawTickers {
		symbol := strings.Split(ticker.Symbol, "/")

		if len(symbol) != 2 {
			continue
		}

		open, _ := strconv.ParseFloat(ticker.Open, 64)
		high, _ := strconv.ParseFloat(ticker.High, 64)
		low, _ := strconv.ParseFloat(ticker.Low, 64)
		volume, _ := strconv.ParseFloat(ticker.Volume, 64)
		ask, _ := strconv.ParseFloat(ticker.Ask[0], 64)
		bid, _ := strconv.ParseFloat(ticker.Bid[0], 64)
		last, _ := strconv.ParseFloat(ticker.Close, 64)

		result = append(result, entities.ExchangeRawTicker{
			BaseSymbol:  symbol[0],
			QuoteSymbol: symbol[1],
			Open:        open,
			High:        high,
			Low:         low,
			Volume:      volume,
			Ask:         ask,
			Bid:         bid,
			Last:        last,
		})
	}
	return result, nil
}
