package exchange

import (
	"tickers-parser/internal/entities"
)

const ExchangesCount = 2

func GetExchangesForTickersUpdate() [ExchangesCount]entities.Exchange {
	return [ExchangesCount]entities.Exchange{Allbit, Exmo}
}
