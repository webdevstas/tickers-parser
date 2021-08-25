package entities

type Ticker struct {
	Id           int     `json:"id" db:"id"`
	BaseSymbol   string  `json:"baseSymbol" db:"baseSymbol"`
	QuoteSymbol  string  `json:"quoteSymbol" db:"quoteSymbol"`
	Volume       float64 `json:"volume" db:"volume"`
	Bid          float64 `json:"bid" db:"bid"`
	Ask          float64 `json:"ask" db:"ask"`
	Open         float64 `json:"open" db:"open"`
	High         float64 `json:"high" db:"high"`
	Low          float64 `json:"low" db:"low"`
	Change       float64 `json:"change" db:"change"`
	UpdatedAt    int     `json:"updatedAt" db:"updatedAt"`
	CreatedAt    int     `json:"createdAt" db:"createdAt"`
	ExchangeId   int     `json:"exchangeId" db:"exchangeId"`
	BaseCoinId   int     `json:"baseCoinId" db:"baseCoinId"`
	QuoteCoinId  int     `json:"quoteCoinId" db:"quoteCoinId"`
	BaseAddress  string  `json:"baseAddress" db:"baseAddress"`
	QuoteAddress string  `json:"quoteAddress" db:"quoteAddress"`
	Enabled      bool    `json:"enabled" db:"enabled"`
}

type ExchangeTickers struct {
	Exchange  string
	Timestamp int64
	Tickers   []Ticker
}
