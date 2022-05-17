package updater

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
)

type TickersStore struct {
	repo repository.IRepository
}

func NewTickersStoreService(r *repository.Repository) TickersStore {
	return TickersStore{
		repo: r,
	}
}

func (ts *TickersStore) SaveTickersForExchange(exchangeId uint, tickers []entities.ExchangeRawTicker) (bool, error) {
	return ts.repo.SaveTickersForExchange(exchangeId, tickers)
}
