package service

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/service/logger"
)

type Monitoring struct {
	url  string
	port int
}

func NewMonitoringService(logger logger.Logger, config *viper.Viper) *Monitoring {
	mon := Monitoring{url: "localhost", port: 4533}
	return &mon
}
