package repository

import (
	"github.com/jmoiron/sqlx"
	"tickers-parser/internal/service"
)

type Repositories struct {
	Ticker *TickerRepository
}

func GetRepositories(db *sqlx.DB, log service.Logger) *Repositories {
	return &Repositories{
		Ticker: GetTickerRepository(db, log),
	}
}
