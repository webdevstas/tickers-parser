package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"io"
	"log"
	"net/http"
	"os"
	"tickers-parser/exchange"
)

func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.Print("Executing NewLogger.")
	return logger
}

func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler.")
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		logger.Print("[" + req.Method + "]" + req.RequestURI + "\n")
	}), nil
}

func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
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

func Register(mux *http.ServeMux, h http.Handler, logger *log.Logger) {
	mux.Handle("/", h)
	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "<h1>Hello Go!</h1>")
	})
	mux.HandleFunc("/new-exchange", func(writer http.ResponseWriter, request *http.Request) {
		body := request.Body
		dec := json.NewDecoder(body)
		var ex exchange.Exchange
		err := dec.Decode(&ex)
		if err != nil {
			logger.Print(err)
		}
		fmt.Printf("%v\n", ex)
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
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
