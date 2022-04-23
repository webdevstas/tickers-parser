package types

type ChannelsPair struct {
	DataChannel   chan interface{}
	CancelChannel chan error
}

func (pair ChannelsPair) CloseAll() {
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
	BaseAddress  float64
	QuoteAddress float64
	Last         float64
}
