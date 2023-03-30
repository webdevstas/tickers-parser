package utils

import (
	"tickers-parser/internal/entities"
)

type IGetEnabledCoins interface {
	GetEnabledCoins() []entities.Coin
}

func GetCoinsMap(r IGetEnabledCoins) map[string]entities.Coin {
	coins := r.GetEnabledCoins()

	coinsMap := Reduce(coins, func(acc map[string]entities.Coin, cur entities.Coin) map[string]entities.Coin {
		acc[cur.Symbol] = cur
		return acc
	}, make(map[string]entities.Coin))

	return coinsMap
}
