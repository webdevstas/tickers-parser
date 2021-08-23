package exchange

import (
	"tickers-parser/internal/entities"
)

func GetExchangesForTickersUpdate() []entities.Exchange {
	var exchanges []entities.Exchange
	exchanges = append(exchanges, ExmoExchange())
	return exchanges
}
