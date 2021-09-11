package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
	"tickers-parser/internal/config"
	"tickers-parser/internal/modules"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/repository/postgres"
	"tickers-parser/internal/services/logger"
)

func Register(s *modules.Services) {
	s.Tasks.RunTasks()
}

func StartSaverApp() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfigModule,
			postgres.ConnectToPostgres,
			repository.GetRepositories,
			modules.InitSaverModule,
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
