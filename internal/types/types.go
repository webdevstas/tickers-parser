package types

type ChannelsPair struct {
	DataChannel   chan interface{}
	CancelChannel chan error
}
