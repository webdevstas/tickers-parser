package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/updater/exchange"
)

type Services struct {
	Scheduler  *Scheduler
	Monitoring *Monitoring
	Exchanges  []entities.Exchange
	Tasks      *Tasks
}

func GetServices(l Logger, c *viper.Viper) *Services {
	scheduler := InitScheduler(l)
	return &Services{
		Scheduler:  scheduler,
		Monitoring: NewMonitoring(l, c),
		Exchanges:  exchange.GetExchangesForTickersUpdate(),
		Tasks:      TasksService(scheduler),
	}
}
