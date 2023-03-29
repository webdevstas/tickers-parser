package entities

import "time"

type Coin struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Key       string
	Name      string
	Symbol    string
	Enabled   bool
	Price     float64
	Volume    float64
	Max       float64
	Min       float64
}

func (c *Coin) CalculatePrice(tickers []Ticker, usdtCoins map[string]Coin) {
	if len(tickers) == 0 {
		return
	}

	var price, volume float64

	for _, ticker := range tickers {
		if ticker.QuoteSymbol == "USDT" || ticker.QuoteSymbol == "USD" {
			price += (ticker.Last * ticker.Volume)
			volume += ticker.Volume
		} else if usdtCoin := usdtCoins[ticker.QuoteSymbol]; usdtCoin.Price > 0 {
			price += ((ticker.Last * ticker.Volume) * usdtCoin.Price)
			volume += ticker.Volume
		}
	}

	if price > 0 && volume > 0 {
		c.Price = price / volume
		c.Volume = volume
	}
}
