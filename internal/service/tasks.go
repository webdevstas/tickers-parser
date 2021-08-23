package service

import (
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

	for _, ex := range exchanges {
		ex.FetchTickers()
	}

	return nil
}

func TasksService(s *Scheduler) *Tasks {
	t := Tasks{
		s: s,
	}
	return &t
}
