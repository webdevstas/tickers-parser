package entities

type Ticker struct {
	id           int
	baseSymbol   string
	quoteSymbol  string
	volume       float32
	bid          float32
	open         float32
	high         float32
	low          float32
	change       float32
	updatedAt    int
	createdAt    int
	exchangeId   int
	baseCoinId   int
	quoteCoinId  int
	baseAddress  string
	quoteAddress string
	enabled      bool
}
