package server

import (
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

	logrus.Infof("Server listening on %s", server.ListenAddr)

	// Start the server with logging
	if err := server.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("Server error: %v", err)
	}
}

// Shutdown is a Graceful shutdown method
func (server *GracefulShutdown) Shutdown() {
	logrus.Info("Shutting down server...")

	// Gracefully shutdown the server
	if server.httpServer != nil {
		if err := server.httpServer.Shutdown(nil); err != nil {
			logrus.Errorf("Error shutting down server: %+v\n", err)
		}
	}

	logrus.Info("Server shutdown complete.")
	os.Exit(0)
}
