package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/pkg/utils"
)

type kucoin struct {
	entities.Exchange
}

type KucoinTicker struct {
	Symbol      string
	High        string
	Low         string
	Vol         string
	ChangePrice string
	Last        string
}
type KucoinData struct {
	Ticker []KucoinTicker
}
type KucoinResponse struct {
	Data KucoinData
}

func getKucoin() kucoin {
	return kucoin{}
}

func (k kucoin) FetchTickers(http *http_client.HttpClient) ([]entities.ExchangeRawTicker, error) {
	url := "https://api.kucoin.com/api/v1/market/allTickers"
	var res KucoinResponse

	err := http.FetchJson(url, &res)

	if err != nil {
		return nil, err
	}

	return utils.Map(res.Data.Ticker, func(el KucoinTicker) entities.ExchangeRawTicker {
		symbol := strings.Split(el.Symbol, "-")
		high, _ := strconv.ParseFloat(el.High, 64)
		low, _ := strconv.ParseFloat(el.Low, 64)
		vol, _ := strconv.ParseFloat(el.Vol, 64)
		change, _ := strconv.ParseFloat(el.ChangePrice, 64)
		last, _ := strconv.ParseFloat(el.Last, 64)
		return entities.ExchangeRawTicker{
			BaseSymbol:  symbol[0],
			QuoteSymbol: symbol[1],
			High:        high,
			Low:         low,
			Volume:      vol,
			Change:      change,
			Last:        last,
		}
	}), nil
}
