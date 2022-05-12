package types

type ChannelsPair[T any] struct {
	DataChannel   chan T
	CancelChannel chan error
}

func (pair ChannelsPair[any]) CloseAll() {
	close(pair.CancelChannel)
	close(pair.DataChannel)
}

type ExchangeRawTicker struct {
	BaseSymbol   string
	QuoteSymbol  string
	Volume       float64
	Bid          float64
	Ask          float64
	High         float64
	Low          float64
	Change       float64
	BaseAddress  string
	QuoteAddress string
	Last         float64
	Open         float64
}

type TickersFetchable interface {
	FetchTickers() ([]ExchangeRawTicker, error)
}
