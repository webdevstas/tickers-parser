package modules

import (
	"github.com/spf13/viper"
)

func NewConfigModule(logger Logger) *viper.Viper {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("Config error: %v", err)
	}
	logger.Info("Config module started")
	return viper.GetViper()
}
