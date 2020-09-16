package main

import (
	"iap/handlers"
	"iap/validators"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	validators.Init()
	route := mux.NewRouter()
	route.HandleFunc("/iap/{provider}/verify", handlers.Verify).Methods("POST")
	route.HandleFunc("/iap/google/acknowledgeSubscription", handlers.AcknowledgeSubscription).Methods("POST")
	http.ListenAndServe(":3000", route)

}
