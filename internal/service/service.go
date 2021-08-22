package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/service/updater/exchange"
)

type Services struct {
	Scheduler  *Scheduler
	Monitoring *Monitoring
	Exchanges  *exchange.Exchanges
}

func GetServices(l Logger, c *viper.Viper) *Services {
	return &Services{
		Scheduler:  NewScheduler(l),
		Monitoring: NewMonitoring(l, c),
		Exchanges:  exchange.GetExchangesForTickersUpdate(),
	}
}
