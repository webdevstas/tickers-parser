package storage

import (
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
	return true, nil
}
