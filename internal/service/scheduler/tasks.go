package scheduler

import (
	"github.com/spf13/viper"
	"runtime"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/logger"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
	"tickers-parser/internal/types"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler *Scheduler
	log       logger.Logger
	storage   *storage.Storage
	config    *viper.Viper
	ITasks
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", t.config.GetInt("app.tickersInterval")*60*1000, false, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) (interface{}, error) {
	exchanges := exchange.GetExchangesForTickersUpdate()
	exchangesCount := len(exchanges)
	tickersChannels := types.ChannelsPair{
		DataChannel:   make(chan interface{}, exchangesCount),
		CancelChannel: make(chan error, exchangesCount),
	}
	saveChannels := types.ChannelsPair{
		CancelChannel: make(chan error, exchangesCount),
		DataChannel:   make(chan interface{}, exchangesCount),
	}

	for _, ex := range exchanges {
		go func(exchange entities.IExchange, channels types.ChannelsPair) {
			res, err := exchange.FetchRawTickers()
			if err != nil {
				channels.CancelChannel <- err
				return
			}
			channels.DataChannel <- res
		}(ex, tickersChannels)
	}

	for i := 0; i < exchangesCount; i++ {
		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Warn(err)
			continue
		case result := <-tickersChannels.DataChannel:
			tickers := result.(entities.ExchangeRawTickers)
			go func(channels types.ChannelsPair) {
				res, err := t.storage.Save(tickers.Exchange, tickers.Timestamp, tickers.RawTickers)
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
				t.log.Info("[scheduler/tickers] Tickers saved for " + tickers.Exchange)
				runtime.Gosched()
			}
		}
	}
	tickersChannels.CloseAll()
	saveChannels.CloseAll()
	t.log.Info("[scheduler/tickers] All channels closed")
	return nil, nil
}

func NewTasksService(l logger.Logger, st *storage.Storage, c *viper.Viper) *Tasks {
	t := Tasks{
		scheduler: InitScheduler(l),
		log:       l,
		storage:   st,
		config:    c,
	}
	return &t
}
