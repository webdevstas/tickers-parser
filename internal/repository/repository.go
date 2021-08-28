package repository

import (
	"github.com/jmoiron/sqlx"
	"tickers-parser/internal/service/logger"
)

type Repositories struct {
	Ticker *TickerRepository
}

func GetRepositories(db *sqlx.DB, log logger.Logger) *Repositories {
	return &Repositories{
		Ticker: GetTickerRepository(db, log),
	}
}
