package updater

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/updater/exchange"
)

func GetExchangesForTickersUpdate() []entities.IExchange {
	return []entities.IExchange{exchange.Allbit, exchange.Exmo, exchange.Ascendex}
}
