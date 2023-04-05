package repository

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/updater/exchange"
	"tickers-parser/pkg/utils"
)

type IExchangeRepository interface {
	GetExchangesForTickersUpdate() []entities.Exchange
}

func (r Repository) GetExchangesForTickersUpdate() []entities.Exchange {
	var exchanges []entities.Exchange
	r.Exchange(true).Where("enabled = true").Find(&exchanges)

	return utils.Map(exchanges, func(el entities.Exchange) entities.Exchange {
		api := exchange.ExchangeMapping[el.Key]
		return entities.Exchange{
			ID:               el.ID,
			Key:              el.Key,
			Name:             el.Name,
			FetchTimeout:     el.FetchTimeout,
			TickersFetchable: api,
		}
	})
}
