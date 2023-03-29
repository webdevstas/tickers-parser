package updater

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/types"
)

type TickersStore struct {
	repo types.IRepository
}

func NewTickersStoreService(r *repository.Repository) TickersStore {
	return TickersStore{
		repo: r,
	}
}

func (ts *TickersStore) SaveTickersForExchange(exchangeId uint, tickers []entities.ExchangeRawTicker) (bool, error) {
	return ts.repo.SaveTickersForExchange(exchangeId, tickers)
}
