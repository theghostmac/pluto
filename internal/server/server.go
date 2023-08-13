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
