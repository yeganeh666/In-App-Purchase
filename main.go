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
	route.HandleFunc("/iap/apple/verify", handlers.AppleVerify).Methods("POST")
	route.HandleFunc("/iap/google/verify", handlers.GoogleVerify).Methods("POST")
	http.ListenAndServe(":3000", route)

}
