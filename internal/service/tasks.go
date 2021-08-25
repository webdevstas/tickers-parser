package service

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
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
	tickersChan := make(chan entities.ExchangeTickers, 5)
	cancelChan := make(chan error)

	for _, ex := range exchanges {
		go ex.FetchTickers(tickersChan, cancelChan)
	}

	for {
		select {
		case err := <-cancelChan:
			t.log.Error(err)
		case tickers := <-tickersChan:
			err := t.storage.Save(tickers.Exchange, tickers.Timestamp, tickers.Tickers)
			if err != nil {
				t.log.Error(err)
			}
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
