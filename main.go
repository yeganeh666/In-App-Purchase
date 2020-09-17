package main

import (
	"iap/handlers"
	"iap/services/google"
	"iap/validators"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	validators.Init()
	app := &google.App{
		DefaultHTTPClient:       http.DefaultClient,
		PubsubVerificationToken: os.Getenv("PUBSUB_VERIFICATION_TOKEN"),
	}
	route := mux.NewRouter()
	route.HandleFunc("/iap/{provider}/verify", handlers.Verify).Methods("POST")
	route.HandleFunc("/iap/google/acknowledgeSubscription", handlers.AcknowledgeSubscription).Methods("POST")
	route.HandleFunc("/pubsub/message/list", app.ListHandler)
	route.HandleFunc("/pubsub/message/receive", app.ReceiveMessagesHandler)
	http.ListenAndServe(":3000", route)

}
