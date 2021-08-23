package service

import (
	"fmt"
	"tickers-parser/internal/service/updater/exchange"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	s *Scheduler
	ITasks
}

func (t *Tasks) RunTasks() {
	t.s.ScheduleRecurrentTask("tickers", 60*1000, true, startTickersParsing)
}

func startTickersParsing(args ...interface{}) error {
	exchanges := exchange.GetExchangesForTickersUpdate()
	tickersChan := make(chan interface{}, 5)

	for _, ex := range exchanges {
		go ex.FetchTickers(tickersChan)
	}

	for ticker := range tickersChan {
		fmt.Print(ticker)
	}

	return nil
}

func TasksService(s *Scheduler) *Tasks {
	t := Tasks{
		s: s,
	}
	return &t
}
