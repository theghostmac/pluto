package main

import (
	"flag"
	"github.com/theghostmac/pluto/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	listenAddr := flag.String("listenAddress", ":8080", "listening on the default port")
	flag.Parse()

	startServer := server.RunServer{
		ListenAddr: *listenAddr,
	}
	go func() {
		err := startServer.Run()
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	startServer.Shutdown()
}
