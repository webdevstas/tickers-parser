package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	"tickers-parser/pkg/utils"
)

type bitforex struct {
	entities.Exchange
}

func GetBitforex() bitforex {
	return bitforex{}
}

type bitforexTicker struct {
	high24hr      string
	percentChange string
	last          string
	highestBid    string
	quoteVolume   string
	baseVolume    string
	lowestAsk     string
	open          string
	low24hr       string
}

func (b bitforex) FetchTickers() ([]entities.ExchangeRawTicker, error) {
	tickersUrl := "https://www.bitforex.com/server/market.act?cmd=getAllMatchTypeTicker"
	response := make(map[string]bitforexTicker)
	tickers := make([]entities.ExchangeRawTicker, 0)

	utils.FetchJson(tickersUrl, &response)

	for key, rawTicker := range response {
		symbols := strings.Split(key, "_")
		high, _ := strconv.ParseFloat(rawTicker.high24hr, 64)
		low, _ := strconv.ParseFloat(rawTicker.low24hr, 64)
		bid, _ := strconv.ParseFloat(rawTicker.highestBid, 64)
		ask, _ := strconv.ParseFloat(rawTicker.lowestAsk, 64)
		open, _ := strconv.ParseFloat(rawTicker.open, 64)
		last, _ := strconv.ParseFloat(rawTicker.last, 64)
		baseVolume, _ := strconv.ParseFloat(rawTicker.baseVolume, 64)

		if len(symbols) != 2 {
			continue
		}

		resultTicker := entities.ExchangeRawTicker{
			BaseSymbol:  symbols[0],
			QuoteSymbol: symbols[1],
			High:        high,
			Low:         low,
			Bid:         bid,
			Ask:         ask,
			Open:        open,
			Last:        last,
			Volume:      baseVolume,
		}

		tickers = append(tickers, resultTicker)
	}
	return tickers, nil
}
