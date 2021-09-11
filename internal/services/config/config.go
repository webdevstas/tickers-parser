package config

import (
	"github.com/spf13/viper"
	"tickers-parser/internal/services/logger"
)

func NewConfigModule(logger logger.Logger) *viper.Viper {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Config error: %v", err)
	}
	return viper.GetViper()
}
