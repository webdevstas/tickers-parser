package exchange

import "tickers-parser/internal/entities"

type Exchanges struct {
	Exmo entities.Exchange
}

func GetExchangesForTickersUpdate() *Exchanges {
	return &Exchanges{
		Exmo: ExmoExchange(),
	}
}
