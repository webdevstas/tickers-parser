package modules

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/scheduler"
	"tickers-parser/internal/services/storage"
	"tickers-parser/internal/services/updater"
)

type Services struct {
	Monitoring *service.Monitoring
	Exchanges  []entities.IExchange
	Tasks      *scheduler.Tasks
	Storage    *storage.Storage
}

func InitSaverModule(l logger.Logger, c *viper.Viper) *Services {
	fileSaver := storage.NewFileSaver(c.GetString("app.dataRoot"))
	fileStorageService := storage.NewStorageService(fileSaver)
	return &Services{
		Monitoring: service.NewMonitoringService(l, c),
		Exchanges:  updater.GetExchangesForTickersUpdate(),
		Tasks:      scheduler.NewTasksService(l, fileStorageService, c),
		Storage:    fileStorageService,
	}
}
