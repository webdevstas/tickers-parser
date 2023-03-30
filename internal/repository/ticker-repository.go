package repository

import (
	"context"
	"tickers-parser/internal/entities"
	"time"

	"gorm.io/gorm/clause"
)

func (r *Repository) SaveTickersForExchange(exchangeId uint, tickers []entities.Ticker) (bool, error) {
	r.Ticker(true).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "base_symbol"}, {Name: "quote_symbol"}, {Name: "exchange_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"volume", "bid", "ask", "open", "high", "low", "last", "created_at", "updated_at"}),
	}).Create(&tickers)

	return true, nil
}

func (r *Repository) GetTickersForCoin(coin *entities.Coin) []entities.Ticker {
	var result []entities.Ticker
	r.Ticker(true).WithContext(context.Background()).Where("enabled = true").Where(`"base_coin_id"=?`, coin.ID).Find(&result)
	return result
}

func (r *Repository) GetAllTickers() []entities.Ticker {
	var tickers []entities.Ticker
	r.Ticker(true).Find(&tickers)
	return tickers
}

func (r *Repository) UpdateTicker(ticker *entities.Ticker) {
	ticker.UpdatedAt = time.Now()
	r.Ticker(true).Where("id=?", ticker.ID).Save(ticker)
}
