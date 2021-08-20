package config

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/service"
)

func NewConfigModule(logger service.Logger) *viper.Viper {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Config error: %v", err)
	}
	logger.Info("Config module started")
	return viper.GetViper()
}
