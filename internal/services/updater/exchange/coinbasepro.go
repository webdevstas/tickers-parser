package exchange

import (
	"fmt"
	"strconv"
	"tickers-parser/internal/entities"
	http_client "tickers-parser/internal/services/http-client"
	"time"
)

type coinbase struct {
	entities.Exchange
}

type CoinbaseProduct struct {
	Id            string
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

type CoinbaseTicker struct {
	Ask     string
	Bid     string
	Volume  string
	TradeId int32 `json:"trade_id"`
	Price   string
	Size    string
	Time    time.Time
}

func GetCoinbase() coinbase {
	return coinbase{}
}

func (c coinbase) FetchTickers(http *http_client.HttpClient) ([]entities.ExchangeRawTicker, error) {
	var products []CoinbaseProduct
	var tickers []entities.ExchangeRawTicker
	productsUrl := "https://api.exchange.coinbase.com/products"
	tickersUrl := "https://api.exchange.coinbase.com/products/%v/ticker"
	err := http.FetchJson(productsUrl, &products)

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		var coinbaseTicker CoinbaseTicker
		time.Sleep(200 * time.Millisecond)
		err := http.FetchJson(fmt.Sprintf(tickersUrl, product.Id), &coinbaseTicker)
		if err != nil {
			continue
		}

		volume, _ := strconv.ParseFloat(coinbaseTicker.Volume, 64)
		bid, _ := strconv.ParseFloat(coinbaseTicker.Bid, 64)
		ask, _ := strconv.ParseFloat(coinbaseTicker.Ask, 64)
		last, _ := strconv.ParseFloat(coinbaseTicker.Price, 64)

		ticker := entities.ExchangeRawTicker{
			BaseSymbol:  product.BaseCurrency,
			QuoteSymbol: product.QuoteCurrency,
			Volume:      volume,
			Bid:         bid,
			Ask:         ask,
			Last:        last,
		}
		tickers = append(tickers, ticker)
	}

	return tickers, nil
}
