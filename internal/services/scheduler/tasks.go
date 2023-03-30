package scheduler

import (
	"runtime"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/updater"
	"tickers-parser/pkg/utils"

	"github.com/spf13/viper"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler  *Scheduler
	log        logger.Logger
	config     *viper.Viper
	repository repository.IRepository
	httpClient *http_client.HttpClient
}

func (t *Tasks) RunTasks() {
	schedule := t.scheduler.ScheduleRecurrentTask
	go schedule("prices", t.config.GetFloat64("app.pricesInterval")*60*1000, false, t.StartPriceCalculation)
	go schedule("tickers", t.config.GetFloat64("app.tickersInterval")*60*1000, false, t.startTickersParsing)
	go schedule("tickers-link", t.config.GetFloat64("app.tickersLinkInterval")*60*1000, false, t.LinkTickersToCoins)
	go schedule("coins-parse", t.config.GetFloat64("app.coinParseInterval")*60*1000, false, t.ParseCoins)
}

func (t *Tasks) startTickersParsing(args ...interface{}) (interface{}, error) {
	exchanges := t.repository.GetExchangesForTickersUpdate()
	exchangesCount := len(exchanges)

	tickersChannels := utils.ChannelsPair[entities.ExchangeTickers]{
		DataChannel:   make(chan entities.ExchangeTickers, exchangesCount),
		CancelChannel: make(chan error, exchangesCount),
	}

	saveChannels := utils.ChannelsPair[interface{}]{
		CancelChannel: make(chan error, exchangesCount),
		DataChannel:   make(chan interface{}, exchangesCount),
	}

	defer func() {
		tickersChannels.CloseAll()
		saveChannels.CloseAll()
		t.log.Info("[scheduler/tickers] All channels closed")
	}()

	for _, ex := range exchanges {
		go func(exchange entities.Exchange, channels utils.ChannelsPair[entities.ExchangeTickers]) {
			exchangeTickers := entities.ExchangeTickers{
				Exchange: exchange,
			}
			res, err := exchange.FetchTickers(t.httpClient)
			if err != nil {
				channels.CancelChannel <- err
				return
			}
			exchangeTickers.Tickers = res
			channels.DataChannel <- exchangeTickers
		}(ex, tickersChannels)
	}

	for i := 0; i < exchangesCount; i++ {
		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Warn(err)
			continue
		case result := <-tickersChannels.DataChannel:
			tickers := result
			go func(channels utils.ChannelsPair[interface{}]) {
				tickerEntities := utils.Map(tickers.Tickers, func(el entities.ExchangeRawTicker) entities.Ticker {
					return utils.RawTickerToEntity(tickers.Exchange.ID, el)
				})

				res, err := t.repository.SaveTickersForExchange(tickers.Exchange.ID, tickerEntities)
				if err != nil {
					channels.CancelChannel <- err
				} else {
					channels.DataChannel <- res
				}
			}(saveChannels)
			select {
			case err := <-saveChannels.CancelChannel:
				t.log.Warn(err)
				continue
			case <-saveChannels.DataChannel:
				t.log.Info("[scheduler/tickers] Tickers saved for " + tickers.Exchange.Key)
				runtime.Gosched()
			}
		}
	}
	return nil, nil
}

func (t *Tasks) StartPriceCalculation(args ...interface{}) (interface{}, error) {
	coins := t.repository.GetEnabledCoins()
	coinsMap := utils.GetCoinsMap(coins)

	for _, coin := range coins {
		tickers := t.repository.GetTickersForCoin(&coin)
		coin.CalculatePrice(tickers, coinsMap)
		t.repository.UpdateCoin(&coin)
	}
	return nil, nil
}

func (t *Tasks) LinkTickersToCoins(args ...interface{}) (interface{}, error) {
	tickers := t.repository.GetAllTickers()
	coins := t.repository.GetEnabledCoins()
	coinsMap := utils.GetCoinsMap(coins)

	for _, ticker := range tickers {
		if ticker.LinkTickerToCoins(coinsMap) {
			t.repository.UpdateTicker(&ticker)
		}
	}

	return nil, nil
}

func (t *Tasks) ParseCoins(args ...interface{}) (interface{}, error) {
	coins := updater.ParseCoinsFromCryptorank(t.httpClient)
	chunkLength := 1000
	done := 0

	for i := chunkLength; i <= len(coins); i += chunkLength {
		if i >= len(coins) {
			i = len(coins)
		}
		t.repository.InsertCoins(coins[done:i])
		done += chunkLength
	}

	return nil, nil
}

func NewTasksService(l logger.Logger, c *viper.Viper, r *repository.Repository, h *http_client.HttpClient) *Tasks {
	return &Tasks{
		scheduler:  InitScheduler(l),
		log:        l,
		config:     c,
		repository: r,
		httpClient: h,
	}
}
