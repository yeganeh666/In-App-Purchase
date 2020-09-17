package routes

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func Serve(ctx context.Context, router *mux.Router, l hclog.Logger) error {
	s := http.Server{
		Addr:         ":3000",                                          // configure the bind address
		Handler:      router,                                           // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                  // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                 // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Info("Starting server on port 3000")

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	l.Info("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
	}()

	if err := s.Shutdown(ctxShutDown); err != nil {
		l.Error("server Shutdown Failed:", err)
		os.Exit(1)
	}
	l.Info("server exited properly")

	return nil
}
