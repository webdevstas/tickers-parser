package utils

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/types"
)

func GetCoinsMap(r types.IRepository) map[string]entities.Coin {
	coins := r.GetEnabledCoins()

	coinsMap := Reduce(coins, func(acc map[string]entities.Coin, cur entities.Coin) map[string]entities.Coin {
		acc[cur.Symbol] = cur
		return acc
	}, make(map[string]entities.Coin))

	return coinsMap
}
