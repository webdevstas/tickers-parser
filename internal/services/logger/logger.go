package logger

import (
	"os"
	"tickers-parser/internal/services/config"
	"time"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Print(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Errorf(format string, args ...interface{})
	Printf(format string, args ...interface{})
}

type Logger struct {
	ILogger
}

func NewLogger(c *config.Config) *Logger {
	logr := logrus.New()
	logr.Level = logrus.Level(c.GetInt32("logger.appLoggerLogLevel"))
	logr.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           time.RFC822,
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	})
	logr.SetOutput(os.Stdout)
	return &Logger{
		logr,
	}
}
