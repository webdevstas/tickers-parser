package app

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"net/http"
	"os"
	"tickers-parser/internal/config"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/repository/postgres"
	"tickers-parser/internal/service"
)

func NewHandler(logger service.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		logger.Info("[" + req.Method + "]" + req.RequestURI + "\n")
	}), nil
}

func NewMux(lc fx.Lifecycle, logger service.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go func() {
				err := server.ListenAndServe()
				if err != nil {

				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})
	return mux
}

func Register(mux *http.ServeMux, h http.Handler, s *service.Services) {
	mux.Handle("/", h)
}

func StartApp() {
	app := fx.New(
		fx.Provide(
			service.NewLogger,
			config.NewConfigModule,
			postgres.NewDbConnection,
			repository.GetRepositories,
			service.GetServices,
			NewHandler,
			NewMux,
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
