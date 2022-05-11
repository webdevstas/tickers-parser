package exchange

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/types"
	"tickers-parser/pkg/utils"
)

type ascendex struct {
	*entities.Exchange
}

func GetAscendex() *ascendex {
	return &ascendex{}
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

func (a *ascendex) FetchTickers() ([]types.ExchangeRawTicker, error) {
	tickersUrl := "https://ascendex.com/api/pro/v1/ticker"
	var response ascendexResponse
	err := utils.FetchJson(tickersUrl, &response)

	if err != nil {
		return nil, err
	}

	rawTickers := response.Data
	result := make([]types.ExchangeRawTicker, 0, len(rawTickers))

	for _, ticker := range rawTickers {
		baseSymbol := strconv.
		result = append(result, types.ExchangeRawTicker{
			BaseSymbol: ,
		})
	}
}
