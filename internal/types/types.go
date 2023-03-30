package types

import "tickers-parser/internal/entities"

type ChannelsPair[T any] struct {
	DataChannel   chan T
	CancelChannel chan error
}

func (pair ChannelsPair[any]) CloseAll() {
	close(pair.CancelChannel)
	close(pair.DataChannel)
}

type IRepository interface {
	IExchangeRepository
	ITickerRepository
	ICoinRepository
}

type ITickerRepository interface {
	GetTickersForCoin(coin *entities.Coin) []entities.Ticker
	GetAllTickers() []entities.Ticker
	SaveTickersForExchange(exchangeId uint, tickers []entities.Ticker) (bool, error)
	UpdateTicker(ticker *entities.Ticker)
}

type ICoinRepository interface {
	GetEnabledCoins() []entities.Coin
	UpdateCoin(coin *entities.Coin)
	InsertCoins([]entities.Coin)
}

type IExchangeRepository interface {
	GetExchangesForTickersUpdate() []entities.Exchange
}
