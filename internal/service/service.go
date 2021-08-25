package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
)

type Services struct {
	Scheduler  *Scheduler
	Monitoring *Monitoring
	Exchanges  []entities.Exchange
	Tasks      *Tasks
	Storage    *storage.Storage
}

func GetServices(l Logger, c *viper.Viper) *Services {
	scheduler := InitScheduler(l)
	fileSaver := storage.NewFileSaver("./")
	storage := storage.NewStorageService(fileSaver)
	return &Services{
		Scheduler:  scheduler,
		Monitoring: NewMonitoringService(l, c),
		Exchanges:  exchange.GetExchangesForTickersUpdate(),
		Tasks:      NewTasksService(scheduler, l, storage),
		Storage:    storage,
	}
}
