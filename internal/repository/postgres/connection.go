package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"tickers-parser/internal/service"
)

type DbConf struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func ConnectToPostgres(config *viper.Viper, logger service.Logger) (*sqlx.DB, error) {

	var conf = DbConf{
		host:     config.GetString("postgres.url"),
		port:     config.GetString("postgres.port"),
		user:     config.GetString("postgres.user"),
		dbname:   config.GetString("postgres.database"),
		password: config.GetString("postgres.password"),
	}

	var db *sqlx.DB

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.host, conf.port, conf.user, conf.dbname, conf.password))

	if err != nil {
		logger.Error(err)
	}

	// force a connection and test that it worked
	err = db.Ping()

	if err != nil {
		logger.Errorf("Error to connect DB: %v", err)
	}

	logger.Info("Connection with Postgress succeed")
	return db, nil
}
