package updater

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
)

func GetExchangesForTickersUpdate(r *repository.Repositories) []entities.Exchange {
	var exchanges []entities.Exchange
	r.Exchange.Where("enabled = true").Find(&exchanges)
	return exchanges
}
