package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"tickers-parser/modules"
)

type DbConf struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func NewDbConnection(config *viper.Viper, logger modules.Logger) (*sqlx.DB, error) {
	var conf = DbConf{
		host:     config.GetString("url"),
		port:     config.GetString("port"),
		user:     config.GetString("user"),
		dbname:   config.GetString("database"),
		password: config.GetString("password"),
	}

	var db *sqlx.DB

	db, error := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.host, conf.port, conf.user, conf.dbname, conf.password))

	if error != nil {
		logger.Error(error)
	}

	// force a connection and test that it worked
	err := db.Ping()

	if err != nil {
		logger.Error("Error to connect DB")
	}
	return db, nil
}
