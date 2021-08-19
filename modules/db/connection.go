package db

import (
	"github.com/spf13/viper"
	"tickers-parser/modules"
)

func NewDbConnection(config viper.Viper, logger modules.Logger) {
	logger.Printf("Db start on:%v", config.Get("url"))
}
