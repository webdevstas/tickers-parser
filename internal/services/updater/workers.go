package updater

import (
	"tickers-parser/internal/entities"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/pkg/utils"
	"time"
)

type saveCb func(uint, []entities.Ticker) (bool, error)

func FetchTickersWorker(exchange entities.Exchange, channels utils.ChannelsPair[entities.ExchangeTickers], httpClient *http_client.HttpClient) {
	exchangeTickers := entities.ExchangeTickers{
		Exchange: exchange,
	}
	timeout := time.Duration(exchange.FetchTimeout * uint(time.Minute))
	time.Sleep(timeout)

	res, err := exchange.FetchTickers(httpClient)
	if err != nil {
		channels.CancelChannel <- err
		return
	}
	exchangeTickers.Tickers = res
	channels.DataChannel <- exchangeTickers
}

func SaveTickersWorker(tickers entities.ExchangeTickers, channels utils.ChannelsPair[interface{}], saveCb saveCb) {
	tickerEntities := utils.Map(tickers.Tickers, func(el entities.ExchangeRawTicker) entities.Ticker {
		return utils.RawTickerToEntity(tickers.Exchange.ID, el)
	})

	res, err := saveCb(tickers.Exchange.ID, tickerEntities)
	if err != nil {
		channels.CancelChannel <- err
	} else {
		channels.DataChannel <- res
	}
}
