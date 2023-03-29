package repository

import (
	"tickers-parser/internal/entities"
	"time"
)

func (r *Repository) GetEnabledCoins() []entities.Coin {
	var result []entities.Coin
	r.Coin(true).Where("enabled = true").Find(&result)
	return result
}

func (r Repository) UpdateCoin(coin *entities.Coin) {
	coin.UpdatedAt = time.Now()
	r.Coin(true).Where("id = ?", coin.ID).UpdateColumns(coin)
}
