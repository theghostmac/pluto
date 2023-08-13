package network

type NetworkAddress string

type RPC struct {
	Source  string
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(transport Transport) error
	SendAMessage(NetworkAddress, []byte) error
	Address() NetworkAddress
}
