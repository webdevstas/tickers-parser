package utils

import (
	"tickers-parser/internal/entities"
)

func Map[F, T any](iterable []F, cb func(el F) T) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		res = append(res, cb(el))
	}

	return res
}

func Filter[T comparable](iterable []T, cb func(el T) bool) []T {
	var res = make([]T, 0, len(iterable))

	for _, el := range iterable {
		if cb(el) {
			res = append(res, el)
		}
	}

	return res
}

func Reduce[T, K any](iterable []T, cb func(K, T) K, initVal K) K {
	res := initVal

	for _, el := range iterable {
		res = cb(res, el)
	}

	return res
}

func Find[T comparable](iterable []T, cb func(el T) bool) (found T, index int) {
	var result T
	var foundIndex = -1
	for i, el := range iterable {
		if cb(el) {
			result = el
			foundIndex = i
		}
	}
	return result, foundIndex
}

func RawTickerToEntity(exchangeId uint, rawTicker entities.ExchangeRawTicker) entities.Ticker {
	return entities.Ticker{
		BaseSymbol:   rawTicker.BaseSymbol,
		QuoteSymbol:  rawTicker.QuoteSymbol,
		Volume:       rawTicker.Volume,
		Bid:          rawTicker.Bid,
		Ask:          rawTicker.Ask,
		Open:         rawTicker.Open,
		High:         rawTicker.High,
		Low:          rawTicker.Low,
		Change:       rawTicker.Change,
		ExchangeID:   exchangeId,
		BaseAddress:  rawTicker.BaseAddress,
		QuoteAddress: rawTicker.QuoteAddress,
		Enabled:      true,
		Last:         rawTicker.Last,
	}
}
