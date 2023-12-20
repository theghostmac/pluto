package network

import "sync"

type NetworkAddress string // Ethereum uses 30303 for example. I want to be able to specify my own when I run from scratch.

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
