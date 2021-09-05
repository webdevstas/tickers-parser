package exchange

import (
	"tickers-parser/internal/entities"
)

func GetExchangesForTickersUpdate() []entities.IExchange {
	return []entities.IExchange{Allbit, Exmo, Ascendex}
}
