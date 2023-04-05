package scheduler

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/config"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/updater"
	"tickers-parser/pkg/utils"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler  IScheduler
	config     config.IConfigService
	log        logger.ILogger
	repository repository.IRepository
	httpClient *http_client.HttpClient
}

func (t *Tasks) RunTasks() {
	schedule := t.scheduler.ScheduleRecurrentTask
	go t.scheduler.RunTask("tickers", t.startTickersParsing)
	go schedule("prices", t.config.GetFloat64("app.pricesInterval")*60*1000, false, t.StartPriceCalculation)
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

	workersCount := exchangesCount

	defer func() {
		tickersChannels.CloseAll()
		saveChannels.CloseAll()
		t.log.Info("[scheduler/tickers] All channels closed")
	}()

	queue := updater.NewQueue(exchanges)

loop:
	for {
		for i := 0; i < workersCount; i++ {
			exchange, ok := queue.Dequeue()

			if ok {
				go updater.FetchTickersWorker(exchange, tickersChannels, t.httpClient)
			}
		}

		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Warn(err)
			queue.RestoreUnconfirmed()
			continue loop
		case result := <-tickersChannels.DataChannel:
			tickers := result

			go updater.SaveTickersWorker(tickers, saveChannels, t.repository.SaveTickersForExchange)

			select {
			case err := <-saveChannels.CancelChannel:
				t.log.Warn(err)
				queue.RestoreUnconfirmed()
				continue loop
			case <-saveChannels.DataChannel:
				t.log.Info("[scheduler/tickers] Tickers saved for " + tickers.Exchange.Key)
				queue.Confirm(tickers.Exchange)
			}
		}
	}
}

func (t *Tasks) StartPriceCalculation(args ...interface{}) (interface{}, error) {
	coins := t.repository.GetEnabledCoins()
	coinsMap := utils.GetCoinsMap(coins)

	for _, coin := range coins {
		tickers := t.repository.GetTickersForCoin(&coin)
		if coin.CalculatePrice(tickers, coinsMap) {
			t.repository.UpdateCoin(&coin)
		}
	}
	return nil, nil
}

func (t *Tasks) LinkTickersToCoins(args ...interface{}) (interface{}, error) {
	tickers := t.repository.GetUnlinkedTickers()
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
	chunkLength, done, coinsLength := 1000, 0, len(coins)

	for i := chunkLength; i <= coinsLength; i += chunkLength {
		if i >= coinsLength {
			i = coinsLength
		}
		t.repository.InsertCoins(coins[done:i])
		done += chunkLength
	}

	return nil, nil
}

func NewTasksService(l *logger.Logger, c *config.Config, r *repository.Repository, h *http_client.HttpClient, s *Scheduler) *Tasks {
	return &Tasks{
		scheduler:  s,
		log:        l,
		config:     c,
		repository: r,
		httpClient: h,
	}
}
