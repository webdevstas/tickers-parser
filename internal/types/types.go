package types

type ChannelsPair[T any] struct {
	DataChannel   chan T
	CancelChannel chan error
}

func (pair ChannelsPair[any]) CloseAll() {
	close(pair.CancelChannel)
	close(pair.DataChannel)
}
