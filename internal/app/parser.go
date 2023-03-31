package app

import (
	"os"
	"tickers-parser/internal/postgres"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/config"
	http_client "tickers-parser/internal/services/http-client"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/scheduler"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
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
			repository.GetRepository,
			http_client.GetHttpClient,
			scheduler.NewTasksService,
			scheduler.InitScheduler,
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
