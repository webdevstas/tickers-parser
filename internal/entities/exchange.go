package entities

type Exchange struct {
	Id           int                                                                  `json:"id,omitempty" db:"id"`
	Key          string                                                               `json:"key" db:"key"`
	Name         string                                                               `json:"name,omitempty" db:"name"`
	Enabled      bool                                                                 `json:"enabled,omitempty" db:"enabled"`
	FetchTickers func(dataChannel chan<- ExchangeTickers, cancelChannel chan<- error) `json:"fetchTickers,omitempty" db:"fetchTickers"`
}
