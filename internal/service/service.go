package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/logger"
	"tickers-parser/internal/service/scheduler"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
)

type Services struct {
	Monitoring *Monitoring
	Exchanges  []entities.Exchange
	Tasks      *scheduler.Tasks
	Storage    *storage.Storage
}

func GetServices(l logger.Logger, c *viper.Viper) *Services {
	fileSaver := storage.NewFileSaver(c.GetString("app.dataRoot"))
	fileStorageService := storage.NewStorageService(fileSaver)
	return &Services{
		Monitoring: NewMonitoringService(l, c),
		Exchanges:  exchange.GetExchangesForTickersUpdate(),
		Tasks:      scheduler.NewTasksService(l, fileStorageService, c),
		Storage:    fileStorageService,
	}
}
