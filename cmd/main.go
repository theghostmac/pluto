package main

import (
	"fmt"
	"github.com/theghostmac/pluto/internal/network"
	"github.com/theghostmac/pluto/internal/server"
	"time"
)

func main() {
	// two testing nodes.
	node1 := network.NewLocalTransport("LOCAL NODE 1")

	// Creating the bootloading node with a specific address
	bootloadingNode := network.NewLocalTransport("BOOTLOADER")
	//

	// Remember to:
	// Initialize the blockchain next, and other networks
	// blockchain := InitializeBlockchain()
	// consensus := InitializeConsensusMechanism()
	// I can set any node to be the second node:
	err := bootloadingNode.Connect(node1)
	if err != nil {
		return
	}
	// TODO: delete the bootloadingNode impl.

	err = node1.Connect(bootloadingNode)
	if err != nil {
		return
	}

	go func() {
		for {
			err := node1.SendAMessage(bootloadingNode.Address(), []byte("Hey, Node1, I need some cash!"))
			if err != nil {
				fmt.Printf("Failed to send a message to node 2 due to: %v", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// ------> RPC Server <------
	ops := server.ServerOperations{
		Transports: []network.Transporter{node1},
	}

	serve := server.NewServer(ops)

	// Run server in a separate goroutine, add go in front of this:
	serve.StartServer()

	// Send a signal to the quitChannel to stop the server.
	serve.QuitChannel <- struct{}{}
}
