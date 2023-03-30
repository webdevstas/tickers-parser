package repository

import (
	"tickers-parser/internal/entities"
	"time"

	"gorm.io/gorm/clause"
)

type ICoinRepository interface {
	GetEnabledCoins() []entities.Coin
	UpdateCoin(coin *entities.Coin)
	InsertCoins([]entities.Coin)
}

func (r *Repository) GetEnabledCoins() []entities.Coin {
	var result []entities.Coin
	r.Coin(true).Where("enabled = true").Find(&result)
	return result
}

func (r *Repository) UpdateCoin(coin *entities.Coin) {
	coin.UpdatedAt = time.Now()
	r.Coin(true).Where("id = ?", coin.ID).UpdateColumns(coin)
}

func (r *Repository) InsertCoins(coins []entities.Coin) {
	r.Coin(true).Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "symbol"}}, DoNothing: true}).Create(&coins)
}
