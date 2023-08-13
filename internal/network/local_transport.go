package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	address         NetworkAddress
	consumerChannel chan RPC
	lock            sync.RWMutex
	peers           map[NetworkAddress]*LocalTransport
}

func NewLocalTransport(address NetworkAddress) *LocalTransport {
	return &LocalTransport{
		address:         address,
		consumerChannel: make(chan RPC, 1024),
		lock:            sync.RWMutex{},
		peers:           make(map[NetworkAddress]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	return lt.consumerChannel
}

func (lt *LocalTransport) Connect(tr *LocalTransport) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	if _, exists := lt.peers[tr.Address()]; exists {
		return fmt.Errorf("peer with address %v already exists", tr.Address())
	}

	lt.peers[tr.Address()] = tr

	return nil
}

func (lt *LocalTransport) Address() NetworkAddress {
	return lt.address
}

func (lt *LocalTransport) SendAMessage(NetworkAddress, []byte) error {
	return nil
}
