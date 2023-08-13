package network

import "sync"

type NetworkAddress string

type RPC struct {
	Source  string
	Payload []byte
}

type Transporter interface {
	Consume() <-chan RPC
	Connect(transport Transporter) error
	SendAMessage(NetworkAddress, []byte) error
	Address() NetworkAddress
}

type LocalTransport struct {
	address         NetworkAddress
	consumerChannel chan RPC
	lock            sync.RWMutex
	Peers           map[NetworkAddress]*LocalTransport
}
