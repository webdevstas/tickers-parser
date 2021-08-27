package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
	"tickers-parser/internal/types"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler *Scheduler
	log       Logger
	storage   *storage.Storage
	config    *viper.Viper
	ITasks
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", t.config.GetInt("app.tickersInterval")*60*1000, true, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) {
	exchanges := exchange.GetExchangesForTickersUpdate()
	tickersChannels := types.ChannelsPair{
		DataChannel:   make(chan interface{}, 5),
		CancelChannel: make(chan error),
	}

	for _, ex := range exchanges {
		go ex.FetchTickers(tickersChannels)
	}

	for {
		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Error(err)
		case result := <-tickersChannels.DataChannel:
			saveChannels := types.ChannelsPair{
				CancelChannel: make(chan error),
			}
			tickers := result.(entities.ExchangeTickers)
			go t.storage.Save(tickers.Exchange, tickers.Timestamp, tickers.Tickers, saveChannels)
			select {
			case err := <-saveChannels.CancelChannel:
				t.log.Error(err)
			}
		}
	}
}

func NewTasksService(s *Scheduler, l Logger, st *storage.Storage, c *viper.Viper) *Tasks {
	t := Tasks{
		scheduler: s,
		log:       l,
		storage:   st,
		config:    c,
	}
	return &t
}
