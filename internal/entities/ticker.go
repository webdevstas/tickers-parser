package entities

type Ticker struct {
	From         string  `json:"from" db:"from"`
	To           string  `json:"to" db:"to"`
	Volume       float64 `json:"volume" db:"volume"`
	Bid          float64 `json:"bid" db:"bid"`
	Ask          float64 `json:"ask" db:"ask"`
	Open         float64 `json:"open" db:"open"`
	High         float64 `json:"high" db:"high"`
	Low          float64 `json:"low" db:"low"`
	Change       float64 `json:"change" db:"change"`
	UpdatedAt    int64   `json:"updatedAt" db:"updatedAt"`
	CreatedAt    int     `json:"createdAt" db:"createdAt"`
	ExchangeId   int     `json:"exchangeId" db:"exchangeId"`
	BaseCoinId   int     `json:"baseCoinId" db:"baseCoinId"`
	QuoteCoinId  int     `json:"quoteCoinId" db:"quoteCoinId"`
	BaseAddress  string  `json:"baseAddress" db:"baseAddress"`
	QuoteAddress string  `json:"quoteAddress" db:"quoteAddress"`
	Enabled      bool    `json:"enabled" db:"enabled"`
	Last         float64 `json:"last"`
}
