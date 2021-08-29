package types

type ChannelsPair struct {
	DataChannel   chan interface{}
	CancelChannel chan error
}

func (pair ChannelsPair) CloseAll() {
	close(pair.CancelChannel)
	close(pair.DataChannel)
}
