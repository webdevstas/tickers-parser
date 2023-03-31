package postgres

import (
	"fmt"
	"log"
	"os"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConf struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
}

func ConnectToPostgres(config *config.Config) (*gorm.DB, error) {
	dbLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,                                                // Slow SQL threshold
			LogLevel:                  logger.LogLevel(config.GetInt("logger.sqlLoggerLogLevel")), // Log level
			IgnoreRecordNotFoundError: true,                                                       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                                       // Disable color
		})
	var conf = DbConf{
		host:     config.GetString("postgres.url"),
		port:     config.GetString("postgres.port"),
		user:     config.GetString("postgres.user"),
		dbname:   config.GetString("postgres.database"),
		password: config.GetString("postgres.password"),
	}

	db, _ := gorm.Open(postgres.Open(fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow", conf.host, conf.user, conf.password, conf.dbname, conf.port)), &gorm.Config{
		Logger: dbLogger,
	})

	db.AutoMigrate(&entities.Exchange{}, &entities.Ticker{}, &entities.Coin{}) // For development only

	return db, nil
}
