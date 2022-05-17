package repository

import (
	"tickers-parser/internal/entities"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/updater/exchange"
	"tickers-parser/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IRepository interface {
	SaveTickersForExchange(exchangeId uint, tickers []entities.ExchangeRawTicker) (bool, error)
	GetExchangesForTickersUpdate() []entities.Exchange
}

type Repository struct {
	Exchange *gorm.DB
	Ticker   *gorm.DB
}

func GetRepository(db *gorm.DB, log logger.Logger) *Repository {
	return &Repository{
		Exchange: db.Model(&entities.Exchange{}),
		Ticker:   db.Model(&entities.Ticker{}),
	}
}

func (r *Repository) SaveTickersForExchange(exchangeId uint, tickers []entities.ExchangeRawTicker) (bool, error) {
	for _, ticker := range tickers {
		resultTicker := RawTickerToEntity(exchangeId, ticker)
		r.Ticker.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "baseSymbol"}, {Name: "quoteSymbol"}, {Name: "exchangeId"}},
			UpdateAll: true,
		}).Create(&resultTicker)
	}
	return true, nil
}

func (r *Repository) GetExchangesForTickersUpdate() []entities.Exchange {

	var ExchangeMapping = map[string]entities.TickersFetchable{
		"ascendex": exchange.GetAscendex(),
	}

	var exchanges []entities.Exchange
	r.Exchange.Where("enabled = true").Find(&exchanges)
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

// Usefull functions
func RawTickerToEntity(exchangeId uint, rawTicker entities.ExchangeRawTicker) entities.Ticker {
	return entities.Ticker{
		BaseSymbol:   rawTicker.BaseSymbol,
		QuoteSymbol:  rawTicker.QuoteSymbol,
		Volume:       rawTicker.Volume,
		Bid:          rawTicker.Bid,
		Ask:          rawTicker.Ask,
		Open:         rawTicker.Open,
		High:         rawTicker.High,
		Low:          rawTicker.Low,
		Change:       rawTicker.Change,
		ExchangeId:   exchangeId,
		BaseCoinId:   0,
		QuoteCoinId:  0,
		BaseAddress:  rawTicker.BaseAddress,
		QuoteAddress: rawTicker.QuoteAddress,
		Enabled:      false,
		Last:         rawTicker.Last,
	}
}
