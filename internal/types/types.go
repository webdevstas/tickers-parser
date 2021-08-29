package types

import "fmt"

type ChannelsPair struct {
	DataChannel   chan interface{}
	CancelChannel chan error
}

func (pair ChannelsPair) CloseAll() {
	close(pair.CancelChannel)
	close(pair.DataChannel)
	fmt.Println("Channels closed")
}
