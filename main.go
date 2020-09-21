package main

import (
	"context"
	"iap/env"
	"iap/services/google"
	"iap/validators"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

var (
	bindAddress       = env.String("BIND_ADDRESS", false, ":3002", "Bind address for the SSV verify server")
	logLevel          = env.String("LOG_LEVEL", false, "debug", "Log output level for the server [debug, info, trace]")
	googlePubSubToken = env.String("PUBSUB_VERIFICATION_TOKEN", true, "", "google pub/sub token")
)

func init() {
	validators.Init()
}

func main() {
	l := hclog.New(
		&hclog.LoggerOptions{
			Name:  "iap",
			Level: hclog.LevelFromString(*logLevel),
		},
	)

	err := env.Parse()
	if err != nil {
		l.Error("Error parsing env", "error", err)
		os.Exit(1)
	}

	router := mux.NewRouter()

	pubsubService := google.NewPubSubNotificationService(l, *googlePubSubToken)
	pubsubReq := router.Methods(http.MethodPost).Subrouter()
	pubsubReq.HandleFunc("/", pubsubService.MessageHandler)
	pubsubReq.Use(pubsubService.VerifyRequest)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	s := http.Server{
		Addr:         *bindAddress,                                                      // configure the bind address
		Handler:      ch(router),                                                        // set the default handler
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{InferLevels: true}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                                   // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                                  // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                                 // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("Starting server on", "port", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		l.Error("Server Shutdown with", "error", err)
	} else {
		l.Info("Server Shutdown gracefully")
	}
}
