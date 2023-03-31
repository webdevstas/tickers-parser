package config

import (
	"log"

	"github.com/spf13/viper"
)

type IConfigService interface {
	GetBool(string) bool
	GetFloat64(string) float64
	GetInt32(string) int32
	GetInt64(string) int64
	GetString(string) string
}

type Config struct {
	*viper.Viper
}

func InitConfigModule() *Config {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Config error: %v", err)
	}
	return &Config{
		viper.GetViper(),
	}
}
