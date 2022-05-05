package exchange

import (
	"strconv"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/types"
	"tickers-parser/pkg/utils"
)

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

type allbit struct {
	*entities.Exchange
}

func GetAllbit() *allbit {
	return &allbit{}
}

func (a *allbit) FetchTickers() ([]types.ExchangeRawTicker, error) {
	var rawTickers []allbitTicker
	err := utils.FetchJson("https://allbit.com/api/coin-market-cap-data/", &rawTickers)

	if err != nil {
		return nil, err
	}

	var result []types.ExchangeRawTicker

	for _, t := range rawTickers {
		last, _ := strconv.ParseFloat(t.Last, 64)
		volume, _ := strconv.ParseFloat(t.BaseVolume, 64)
		high, _ := strconv.ParseFloat(t.High24hr, 64)
		low, _ := strconv.ParseFloat(t.Low24hr, 64)
		bid, _ := strconv.ParseFloat(t.HighestBid, 64)
		ask, _ := strconv.ParseFloat(t.LowestAsk, 64)

		ticker := types.ExchangeRawTicker{
			BaseSymbol:  t.BaseAsset,
			QuoteSymbol: t.QuoteAsset,
			Volume:      volume,
			Bid:         bid,
			Ask:         ask,
			High:        high,
			Low:         low,
			Last:        last,
		}
		result = append(result, ticker)
	}
	return result, nil
}
