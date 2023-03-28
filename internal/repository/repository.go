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
	GetEnabledCoins() []entities.Coin
	GetTickersForCoin(coin *entities.Coin) []entities.Ticker
	SaveCoin(coin *entities.Coin)
}

type Repository struct {
	Exchange *gorm.DB
	Ticker   *gorm.DB
	Coin     *gorm.DB
}

func GetRepository(db *gorm.DB, log logger.Logger) *Repository {
	return &Repository{
		Exchange: db.Model(&entities.Exchange{}),
		Ticker:   db.Model(&entities.Ticker{}),
		Coin:     db.Model(&entities.Coin{}),
	}
}

func (r *Repository) SaveTickersForExchange(exchangeId uint, tickers []entities.ExchangeRawTicker) (bool, error) {
	for _, ticker := range tickers {
		resultTicker := RawTickerToEntity(exchangeId, ticker)
		r.Ticker.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "base_symbol"}, {Name: "quote_symbol"}, {Name: "exchange_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"volume", "bid", "ask", "open", "high", "low", "last", "created_at", "updated_at"}),
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

func (r *Repository) GetEnabledCoins() []entities.Coin {
	var result []entities.Coin
	r.Coin.Where("enabled = true").Find(&result)
	return result
}

func (r *Repository) GetTickersForCoin(coin *entities.Coin) []entities.Ticker {
	var result []entities.Ticker
	r.Ticker.Where("enabled = true").Where(`"base_coin_id"=?`, coin.ID).Find(&result)
	return result
}

func (r *Repository) SaveCoin(coin *entities.Coin) {
	r.Coin.Save(coin)
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
