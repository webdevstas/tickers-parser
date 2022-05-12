package updater

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/updater/exchange"
	"tickers-parser/internal/types"
	"tickers-parser/pkg/utils"
)

var exchangeMapping = map[string]types.TickersFetchable{
	"ascendex": exchange.GetAscendex(),
}

func GetExchangesForTickersUpdate(repo *repository.Repositories) []entities.Exchange {
	var exchanges []entities.Exchange
	repo.Exchange.Where("enabled = true").Find(&exchanges)
	return utils.Map(exchanges, func(exchange entities.Exchange) entities.Exchange {
		api := exchangeMapping[exchange.Key]
		return entities.Exchange{
			ID:               exchange.ID,
			Key:              exchange.Key,
			Name:             exchange.Name,
			TickersFetchable: api,
		}
	})
}
