package exchange

import (
	"io/ioutil"
	"net/http"
	"tickers-parser/internal/entities"
)

func ExmoExchange() entities.Exchange {
	exmo := entities.Exchange{
		Key:          "exmo",
		Name:         "Exmo",
		Enabled:      true,
		FetchTickers: FetchTickers,
	}
	return exmo
}

func FetchTickers(dataChannel chan<- interface{}, cancelChannel chan struct{}) {
	apiUrl := "https://api.exmo.com/v1/ticker"
	resp, err := http.Get(apiUrl)

	if err != nil {
		cancelChannel <- struct{}{}
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		cancelChannel <- struct{}{}
	}

	dataChannel <- body
}
