package repository

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/updater/exchange"
	"tickers-parser/pkg/utils"

	"gorm.io/gorm"
)

func (r Repository) GetExchangesForTickersUpdate() []entities.Exchange {

	var ExchangeMapping = map[string]entities.TickersFetchable{
		"ascendex": exchange.GetAscendex(),
	}

	var exchanges []entities.Exchange
	r.Exchange(true).Where("enabled = true").Session(&gorm.Session{NewDB: true}).Find(&exchanges)
	return utils.Map(exchanges, func(exchange entities.Exchange) entities.Exchange {
		api := ExchangeMapping[exchange.Key]
		return entities.Exchange{
			ID:               exchange.ID,
			Key:              exchange.Key,
			Name:             exchange.Name,
			TickersFetchable: api,
		}
	})
}
