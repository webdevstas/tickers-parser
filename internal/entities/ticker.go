package entities

type Ticker struct {
	Id           int     `json:"id" db:"id"`
	BaseSymbol   string  `json:"baseSymbol" db:"baseSymbol"`
	QuoteSymbol  string  `json:"quoteSymbol" db:"quoteSymbol"`
	Volume       float32 `json:"volume" db:"volume"`
	Bid          float32 `json:"bid" db:"bid"`
	Open         float32 `json:"open" db:"open"`
	High         float32 `json:"high" db:"high"`
	Low          float32 `json:"low" db:"low"`
	Change       float32 `json:"change" db:"change"`
	UpdatedAt    int     `json:"updatedAt" db:"updatedAt"`
	CreatedAt    int     `json:"createdAt" db:"createdAt"`
	ExchangeId   int     `json:"exchangeId" db:"exchangeId"`
	BaseCoinId   int     `json:"baseCoinId" db:"baseCoinId"`
	QuoteCoinId  int     `json:"quoteCoinId" db:"quoteCoinId"`
	BaseAddress  string  `json:"baseAddress" db:"baseAddress"`
	QuoteAddress string  `json:"quoteAddress" db:"quoteAddress"`
	Enabled      bool    `json:"enabled" db:"enabled"`
}
