package exchange

type Exchange struct {
	name         string
	url          string
	foundationOn int
}

func newExchange(name string, url string, year int) Exchange {
	return Exchange{name, url, year}
}
