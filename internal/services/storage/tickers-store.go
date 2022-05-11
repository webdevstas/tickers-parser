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

func (s *TickersStore) SaveTickersForExchange(exchangeId uint, tickers []types.ExchangeRawTicker) (bool, error) {
	for _, ticker := range tickers {
		resultTicker := RawTickerToEntity(exchangeId, ticker)
		s.repo.Ticker.Save(resultTicker)
	}
	return true, nil
}

func RawTickerToEntity(exchangeId uint, rawTicker types.ExchangeRawTicker) entities.Ticker {
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
		ExchangeId:   exchangeId,
		BaseCoinId:   0,
		QuoteCoinId:  0,
		BaseAddress:  rawTicker.BaseAddress,
		QuoteAddress: rawTicker.QuoteAddress,
		Enabled:      false,
		Last:         rawTicker.Last,
	}
}
