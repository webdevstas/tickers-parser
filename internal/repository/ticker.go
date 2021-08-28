package repository

import (
	"github.com/jmoiron/sqlx"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/logger"
)

type TickerRepository struct {
	db     *sqlx.DB
	logger logger.Logger
}

func (r *TickerRepository) GetTickers(cond string) ([]entities.Ticker, error) {
	var tickers []entities.Ticker
	rows, err := r.db.Queryx(cond)
	for rows.Next() {
		t := entities.Ticker{}
		err = rows.StructScan(&t)
		if err != nil {
			r.logger.Error(err)
		}
		tickers = append(tickers, t)
	}

	if err != nil {
		r.logger.Errorf("%v", err)
		return tickers, err
	}
	return tickers, nil
}

func GetTickerRepository(db *sqlx.DB, logger logger.Logger) *TickerRepository {
	return &TickerRepository{db: db, logger: logger}
}
