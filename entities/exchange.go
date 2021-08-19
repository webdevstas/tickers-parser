package entities

type Exchange struct {
	id           int
	name         string
	apiUrl       string
	enabled      bool
	fetchTickers func() []Ticker
}
