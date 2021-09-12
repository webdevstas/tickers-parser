package app

import (
	"os"
	"tickers-parser/internal/modules"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/repository/postgres"
	"tickers-parser/internal/services/config"
	"tickers-parser/internal/services/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Register(s *modules.SaverModule) {
	s.Tasks.RunTasks()
}

func StartSaverApp() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			config.InitConfigModule,
			modules.InitSaverModule,
			postgres.ConnectToPostgres,
			repository.GetRepositories,
		),
		fx.Invoke(Register),
		fx.WithLogger(
			func() fxevent.Logger {
				return &fxevent.ConsoleLogger{W: os.Stdout}
			},
		),
	)
	app.Run()
}
