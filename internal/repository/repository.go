package repository

import (
	"github.com/jmoiron/sqlx"
	"tickers-parser/internal/services/logger"
)

type Repositories struct {
	Temp *sqlx.DB
}

func GetRepositories(db *sqlx.DB, log logger.Logger) *Repositories {
	return &Repositories{
		Temp: db,
	}
}
