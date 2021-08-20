package service

import (
	"github.com/spf13/viper"
)

type Monitoring struct {
	url  string
	port int
}

func NewMonitoringModule(logger Logger, config viper.Viper) Monitoring {
	mon := Monitoring{url: "localhost", port: 4533}
	logger.Info("Config: $v", config.Get("url"))
	return mon
}
