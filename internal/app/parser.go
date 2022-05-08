package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
	"tickers-parser/internal/postgres"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/config"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/scheduler"
)

func Register(s *scheduler.Tasks) {
	s.RunTasks()
}

func StartParserApp() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			config.InitConfigModule,
			postgres.ConnectToPostgres,
			repository.GetRepositories,
			scheduler.NewTasksService,
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
