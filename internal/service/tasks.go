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
	cancelChan := make(chan struct{})
	var curExchange entities.Exchange

	for _, ex := range exchanges {
		curExchange = ex
		go ex.FetchTickers(tickersChan, cancelChan)
	}

	for {
		select {
		case <-cancelChan:
			t.log.Error("Error to parse tickers for exchange: " + curExchange.Name) //TODO: Сделать возврат ошибки и имени биржи из канала cancelChan
			return nil
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
