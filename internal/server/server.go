package server

import (
	"fmt"
	"github.com/theghostmac/pluto/internal/network"
	"time"
)

type ServerOperations struct {
	Transports []network.Transporter
}

type Server struct {
	ServerOperations
	RPCChannel  chan network.RPC
	QuitChannel chan struct{}
}

func NewServer(operations ServerOperations) *Server {
	return &Server{
		ServerOperations: operations,
		RPCChannel:       make(chan network.RPC),
		QuitChannel:      make(chan struct{}, 1),
	}
}

func (s *Server) StartServer() {
	s.InitializeTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.RPCChannel:
			fmt.Printf("%+v", rpc)
		case <-s.QuitChannel:
			break free
		case <-ticker.C:
			fmt.Println("Interacting with blocks every 500 milli seconds.")
		}
	}
	fmt.Println("Server shutdown...")
}

func (s *Server) InitializeTransports() {
	for _, trans := range s.Transports {
		go func(trans network.Transporter) {
			for rpc := range trans.Consume() {
				s.RPCChannel <- rpc
			}
		}(trans)
	}
}
