package storage

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/types"
)

type TickersStore struct {
	repo *repository.Repositories
}

func NewTickersStoreService(r *repository.Repositories) TickersStore {
	return TickersStore{
		repo: r,
	}
}

func (s *TickersStore) SaveTickersForExchange(exchangeKey string, tickers []types.ExchangeRawTicker) (bool, error) {
	for _, ticker := range tickers {
		go func(ticker types.ExchangeRawTicker) {
			resultTicker := RawTickerToEntity(ticker)
			s.repo.Ticker.Save(resultTicker)
		}(ticker)
	}
	return true, nil
}

func RawTickerToEntity(rawTicker types.ExchangeRawTicker) entities.Ticker {
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
		ExchangeId:   0,
		BaseCoinId:   0,
		QuoteCoinId:  0,
		BaseAddress:  rawTicker.BaseAddress,
		QuoteAddress: rawTicker.QuoteAddress,
		Enabled:      false,
		Last:         rawTicker.Last,
	}
}
