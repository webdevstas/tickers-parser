package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
	"tickers-parser/internal/config"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/repository/postgres"
	"tickers-parser/internal/service"
)

func Register(s *service.Services) {
	s.Tasks.RunTasks()
}

func StartApp() {
	app := fx.New(
		fx.Provide(
			service.NewLogger,
			config.NewConfigModule,
			postgres.ConnectToPostgres,
			repository.GetRepositories,
			service.GetServices,
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
