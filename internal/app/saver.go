package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"os"
	"tickers-parser/internal/modules"
	"tickers-parser/internal/services/config"
	"tickers-parser/internal/services/logger"
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
