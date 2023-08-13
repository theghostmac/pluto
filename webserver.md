runner.go:
```go
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

```

server.go:
```go
package server

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type GracefulShutdown struct {
	ListenAddr  string
	BaseHandler http.Handler
	httpServer  *http.Server
}

func (server *GracefulShutdown) getRouter() *mux.Router {
	router := mux.NewRouter()
	router.SkipClean(true)
	router.Handle("/", server.BaseHandler)
	return router
}

func (server *GracefulShutdown) Start() {
	router := server.getRouter()
	server.httpServer = &http.Server{
		Addr:    server.ListenAddr,
		Handler: router,
	}
	logrus.Infof("Server listening on %s ", server.ListenAddr)
	logrus.Info("Pluto is active 🚀")

	if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logrus.Fatalf("Server error: %v", err)
	}
}

func (server *GracefulShutdown) Shutdown() {
	logrus.Info("Shutting down server...")

	if err := server.httpServer.Shutdown(nil); err != nil {
		logrus.Errorf("Error during server shutdown: %v", err)
	}
	logrus.Info("Server shutdown complete.")
	os.Exit(0)
}

```

main.go:
```go
func main() {
    listenAddr := flag.String("listenAddress", ":8080", "listening on the default port")
    flag.Parse()

    startServer := server.RunServer{
        ListenAddr: *listenAddr,
    }
    go func () {
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
```