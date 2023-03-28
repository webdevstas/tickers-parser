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

func (c *Coin) CalculatePrice(tickers []Ticker) {
	var sum, volume float64 = 0, 0
	for _, ticker := range tickers {
		sum += (ticker.Last * ticker.Volume)
		volume += ticker.Volume
	}
	c.Price = sum / volume
	c.Volume = volume
}
