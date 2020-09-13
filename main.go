package main

import (
	"iap/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/iap/apple", handlers.AppleService).Methods("POST")
	route.HandleFunc("/iap/apple/verify", handlers.AppleVerify).Methods("POST")
	route.HandleFunc("/iap/google", handlers.GoogleService).Methods("POST")
	route.HandleFunc("/iap/google/verify", handlers.GoogleVerify).Methods("POST")
	http.ListenAndServe(":3000", route)

}
