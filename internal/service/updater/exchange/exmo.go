package exchange

import (
	"io/ioutil"
	"log"
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

func FetchTickers() []entities.Ticker {
	apiUrl := "https://api.exmo.com/v1/ticker"
	resp, _ := http.Get(apiUrl)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
	return nil
}
