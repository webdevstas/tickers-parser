package service

import (
	"github.com/spf13/viper"
)

type Monitoring struct {
	url  string
	port int
}

func NewMonitoringService(logger Logger, config *viper.Viper) *Monitoring {
	mon := Monitoring{url: "localhost", port: 4533}
	return &mon
}
