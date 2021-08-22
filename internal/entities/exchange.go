package entities

type Exchange struct {
	Id           int             `json:"id,omitempty" db:"id"`
	Name         string          `json:"name,omitempty" db:"name"`
	ApiUrl       string          `json:"apiUrl,omitempty" db:"apiUrl"`
	Enabled      bool            `json:"enabled,omitempty" db:"enabled"`
	FetchTickers func() []Ticker `json:"fetchTickers,omitempty" db:"fetchTickers"`
}
