package postgres

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"tickers-parser/internal/services/logger"
)

type DbConf struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func ConnectToPostgres(config *viper.Viper, logger logger.Logger) (*gorm.DB, error) {

	var conf = DbConf{
		host:     config.GetString("postgres.url"),
		port:     config.GetString("postgres.port"),
		user:     config.GetString("postgres.user"),
		dbname:   config.GetString("postgres.database"),
		password: config.GetString("postgres.password"),
	}

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow", conf.host, conf.user, conf.password, conf.dbname, conf.port)))

	if err != nil {
		logger.Error(err)
	}

	logger.Info("Connection with Postgres succeed")
	return db, nil
}
