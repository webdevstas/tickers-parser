package modules

import (
	"log"
	"os"
)

type Logger interface {
	Print(v ...interface{})
}

func NewLogger(m Monitoring) Logger {
	logger := log.New(os.Stdout, "", 0)
	logger.Print("Executing NewLogger." + m.url)
	return logger
}
