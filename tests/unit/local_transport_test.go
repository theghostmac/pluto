package network

import (
	"github.com/theghostmac/pluto/internal/network"
	"testing"
)

func TestLocalTransport_Connect(t *testing.T) {
	lt1 := network.NewLocalTransport("node1")
	lt2 := network.NewLocalTransport("node2")

	err := lt1.Connect(lt2)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	err = lt1.Connect(lt2) // Connecting again should result in an error
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestLocalTransport_SendAMessage(t *testing.T) {
	lt1 := network.NewLocalTransport("node1")
	lt2 := network.NewLocalTransport("node2")

	err := lt1.Connect(lt2)
	if err != nil {
		t.Errorf("Error connecting nodes: %v", err)
	}

	payload := []byte("test message")
	err = lt1.SendAMessage("node2", payload)
	if err != nil {
		t.Errorf("Error sending message: %v", err)
	}
}

func TestLocalTransport_SendAMessage_InvalidDestination(t *testing.T) {
	lt1 := network.NewLocalTransport("node1")
	lt2 := network.NewLocalTransport("node2")

	err := lt1.Connect(lt2)
	if err != nil {
		t.Errorf("Error connecting nodes: %v", err)
	}

	payload := []byte("test message")
	err = lt1.SendAMessage("node3", payload) // Attempting to send to an invalid destination
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
