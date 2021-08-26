package service

import (
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
	ITasks
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", 60*1000, true, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) {
	exchanges := exchange.GetExchangesForTickersUpdate()
	tickersChan := make(chan interface{}, 5)
	cancelChan := make(chan error)
	tickersChannels := types.ChannelsPair{
		DataChannel:   tickersChan,
		CancelChannel: cancelChan,
	}

	for _, ex := range exchanges {
		go ex.FetchTickers(tickersChannels)
	}

	for {
		select {
		case err := <-cancelChan:
			t.log.Error(err)
		case result := <-tickersChan:
			saveChannels := types.ChannelsPair{
				CancelChannel: make(chan error),
			}
			tickers := result.(entities.ExchangeTickers)
			go t.storage.Save(tickers.Exchange, tickers.Timestamp, tickers.Tickers, saveChannels)
		}
	}
}

func NewTasksService(s *Scheduler, l Logger, st *storage.Storage) *Tasks {
	t := Tasks{
		scheduler: s,
		log:       l,
		storage:   st,
	}
	return &t
}
