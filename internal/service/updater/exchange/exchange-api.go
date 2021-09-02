package exchange

import (
	"tickers-parser/internal/entities"
)

func GetExchangesForTickersUpdate() []entities.Exchange {
	return []entities.Exchange{Allbit, Exmo, Ascendex}
}
