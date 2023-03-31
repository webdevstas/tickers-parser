package repository

import (
	"tickers-parser/internal/entities"
	"time"

	"gorm.io/gorm/clause"
)

type ITickerRepository interface {
	GetTickersForCoin(coin *entities.Coin) []entities.Ticker
	GetAllTickers() []entities.Ticker
	GetUnlinkedTickers() []entities.Ticker
	SaveTickersForExchange(exchangeId uint, tickers []entities.Ticker) (bool, error)
	UpdateTicker(ticker *entities.Ticker)
}

func (r *Repository) SaveTickersForExchange(exchangeId uint, tickers []entities.Ticker) (bool, error) {
	r.Ticker(true).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "base_symbol"}, {Name: "quote_symbol"}, {Name: "exchange_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"volume", "bid", "ask", "open", "high", "low", "last", "created_at", "updated_at"}),
	}).Create(&tickers)

	return true, nil
}

func (r *Repository) GetTickersForCoin(coin *entities.Coin) []entities.Ticker {
	var result []entities.Ticker
	r.Ticker(true).Where("enabled = true").Where(`"base_coin_id"=?`, coin.ID).Find(&result)
	return result
}

func (r *Repository) GetAllTickers() []entities.Ticker {
	var tickers []entities.Ticker
	r.Ticker(true).Find(&tickers)
	return tickers
}

func (r *Repository) GetUnlinkedTickers() []entities.Ticker {
	var tickers []entities.Ticker
	r.Ticker(true).Where("base_coin_id = 0 OR quote_coin_id = 0").Find(&tickers)
	return tickers
}

func (r *Repository) UpdateTicker(ticker *entities.Ticker) {
	ticker.UpdatedAt = time.Now()
	r.Ticker(true).Where("id=?", ticker.ID).Save(ticker)
}
