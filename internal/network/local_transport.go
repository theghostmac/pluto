package network

import (
	"fmt"
	"sync"
)

func NewLocalTransport(address NetworkAddress) Transporter {
	return &LocalTransport{
		address:         address,
		consumerChannel: make(chan RPC, 1024),
		lock:            sync.RWMutex{},
		Peers:           make(map[NetworkAddress]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() <-chan RPC {
	return lt.consumerChannel
}

func (lt *LocalTransport) Connect(tr Transporter) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	if _, exists := lt.Peers[tr.Address()]; exists {
		return fmt.Errorf("peer with address %v already exists", tr.Address())
	}

	lt.Peers[tr.Address()] = tr.(*LocalTransport)

	return nil
}

func (lt *LocalTransport) Address() NetworkAddress {
	return lt.address
}

func (lt *LocalTransport) SendAMessage(destination NetworkAddress, payload []byte) error {
	lt.lock.Lock()
	defer lt.lock.Unlock()

	peer, ok := lt.Peers[destination]
	if !ok {
		return fmt.Errorf("%s: could not send a messag to %s", lt.address, destination)
	}

	peer.consumerChannel <- RPC{
		Source:  string(lt.address),
		Payload: payload,
	}
	return nil
}
