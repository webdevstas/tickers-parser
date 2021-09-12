package modules

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/scheduler"
	"tickers-parser/internal/services/storage"
)

type SaverModule struct {
	Monitoring *service.Monitoring
	Tasks      *scheduler.Tasks
	Storage    *storage.Storage
}

func InitSaverModule(l logger.Logger, c *viper.Viper, r *repository.Repositories) *SaverModule {
	fileSaver := storage.NewFileSaver(c.GetString("app.dataRoot"))
	fileStorageService := storage.NewStorageService(fileSaver)
	return &SaverModule{
		Monitoring: service.NewMonitoringService(l, c),
		Tasks:      scheduler.NewTasksService(l, fileStorageService, c, r),
		Storage:    fileStorageService,
	}
}
