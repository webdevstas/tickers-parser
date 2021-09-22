package exchange

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
)

func GetExchangesForTickersUpdate(r *repository.Repositories) []entities.Exchange {
	return []entities.Exchange{Allbit, Exmo, Ascendex}
}
