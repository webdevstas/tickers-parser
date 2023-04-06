package exchange

import "tickers-parser/internal/entities"

var ExchangeMapping = map[string]entities.TickersFetchable{
	"ascendex": getAscendex(),
	"coinbase": getCoinbase(),
	"kucoin":   getKucoin(),
}
