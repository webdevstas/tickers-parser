package exchange

import (
	"strconv"
	"strings"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/utils"
	"time"
)

var allBitKey = "allbit"

var Allbit = entities.Exchange{
	Key:          allBitKey,
	Name:         "Allbit",
	Enabled:      true,
	FetchTickers: allbitFetchTickers,
}

type allbitTicker struct {
	QuoteAssetName string `json:"quoteAssetName,omitempty"`
	Last           string `json:"last,omitempty"`
	QuoteVolume    string `json:"quoteVolume,omitempty"`
	High24hr       string `json:"high24hr,omitempty"`
	QuoteAsset     string `json:"quoteAsset,omitempty"`
	HighestBid     string `json:"highestBid,omitempty"`
	BaseAsset      string `json:"baseAsset,omitempty"`
	IsFrozen       int    `json:"isFrozen,omitempty"`
	PercentChange  string `json:"percentChange,omitempty"`
	Id             int    `json:"id,omitempty"`
	Low24hr        string `json:"low24hr,omitempty"`
	LowestAsk      string `json:"lowestAsk,omitempty"`
	BaseAssetName  string `json:"baseAssetName,omitempty"`
	BaseVolume     string `json:"baseVolume,omitempty"`
}

func allbitFetchTickers() (entities.ExchangeTickers, error) {
	var tickersArr []entities.Ticker
	apiUrl := "https://allbit.com/api/coin-market-cap-data/"
	rawTickers := make(map[string]allbitTicker)
	err := utils.FetchJson(apiUrl, &rawTickers)
	if err != nil {
		return entities.ExchangeTickers{}, err
	}

	for pair, data := range rawTickers {
		symbols := strings.Split(pair, "_")
		volume, _ := strconv.ParseFloat(data.BaseVolume, 64)
		bid, _ := strconv.ParseFloat(data.HighestBid, 64)
		ask, _ := strconv.ParseFloat(data.LowestAsk, 64)
		high, _ := strconv.ParseFloat(data.High24hr, 64)
		low, _ := strconv.ParseFloat(data.Low24hr, 64)
		last, _ := strconv.ParseFloat(data.Last, 64)
		ticker := entities.Ticker{
			BaseSymbol:  symbols[1],
			QuoteSymbol: symbols[0],
			Volume:      volume,
			Bid:         bid,
			Ask:         ask,
			High:        high,
			Low:         low,
			UpdatedAt:   time.Now().Unix(),
			Last:        last,
		}
		tickersArr = append(tickersArr, ticker)
	}
	res := entities.ExchangeTickers{
		Exchange:  allBitKey,
		Timestamp: time.Now().Unix(),
		Tickers:   tickersArr,
	}
	return res, nil
}
