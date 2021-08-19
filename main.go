package main

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"net/http"
	"os"
	"tickers-parser/modules"
)

func NewHandler(logger modules.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		logger.Print("[" + req.Method + "]" + req.RequestURI + "\n")
	}), nil
}

func NewMux(lc fx.Lifecycle, logger modules.Logger) *http.ServeMux {
	logger.Print("Executing NewMux.")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})
	return mux
}

func Register(mux *http.ServeMux, h http.Handler, logger modules.Logger) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		fx.Provide(
			modules.NewLogger,
			modules.NewMonitoring,
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
