package server

import "os"

type RunServer struct {
	ListenAddr string
	Server     *GracefulShutdown
}

func (runner *RunServer) Run() error {
	server := &GracefulShutdown{
		ListenAddr: runner.ListenAddr,
	}
	runner.Server = server
	server.Start()
	return nil
}

func (runner *RunServer) Shutdown() {
	runner.Server.Shutdown()
	os.Exit(0)
}
