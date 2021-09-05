package exchange

import (
	"tickers-parser/internal/entities"
)

type allbit struct {
	*entities.Exchange
}

func getAllbit() *allbit {
	ex := &entities.Exchange{
		Key:        "allbit",
		Name:       "Allbit",
		Enabled:    true,
		TickersUrl: "https://allbit.com/api/coin-market-cap-data/",
	}
	return &allbit{
		ex,
	}
}

var Allbit = getAllbit()

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
