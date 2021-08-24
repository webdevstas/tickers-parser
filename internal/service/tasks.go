package service

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/updater/exchange"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler *Scheduler
	log       Logger
	ITasks
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", 60*1000, true, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) error {
	exchanges := exchange.GetExchangesForTickersUpdate()
	tickersChan := make(chan map[string]entities.ExchangeTickers, 5)
	cancelChan := make(chan error)

	for _, ex := range exchanges {
		go ex.FetchTickers(tickersChan, cancelChan)
	}

	for {
		select {
		case err := <-cancelChan:
			return err
		case tickers := <-tickersChan:
			t.log.Info(tickers) //TODO: Реализовать сервис сохранялку
		}
	}
}

func TasksService(s *Scheduler, l Logger) *Tasks {
	t := Tasks{
		scheduler: s,
		log:       l,
	}
	return &t
}
