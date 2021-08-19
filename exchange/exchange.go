package exchange

import "fmt"

type Exchange struct {
	Name         string
	Url          string
	FoundationOn int
}

func (e Exchange) PrintName() {
	fmt.Printf("name: %v\n", e.Name)
}

func newExchange(name string, url string, year int) Exchange {
	return Exchange{name, url, year}
}
