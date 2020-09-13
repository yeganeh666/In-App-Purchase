package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//AppleService create new apple service
func AppleService(w http.ResponseWriter, r *http.Request) {
	X = &Apple{}
	X.NewService()
	fmt.Fprintf(w, "Done!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//AppleVerify handler
func AppleVerify(w http.ResponseWriter, r *http.Request) {
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&jsonMap)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonMap)
	//var x Services = &Apple{}
	//x.NewService()
	fmt.Println(X.Verify(jsonMap))
	fmt.Fprintf(w, "req map: %+v", jsonMap)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

//GoogleService create new google service
func GoogleService(w http.ResponseWriter, r *http.Request) {
	X = &Google{}
	X.NewService()
	fmt.Fprintf(w, "Done!")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//GoogleVerify handler
func GoogleVerify(w http.ResponseWriter, r *http.Request) {
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&jsonMap)
	if err != nil {
		log.Fatal(err)
	}
	// jsonMap {
	// 	"package":        "package",
	// 	"subscriptionID": "subscriptionID",
	// 	"purchaseToken":  "purchaseToken",
	// }
	fmt.Println(jsonMap)
	// var x2 Services = &Google{}
	// x2.NewService()
	fmt.Println(X.Verify(jsonMap))
	fmt.Fprintf(w, "req map: %+v", jsonMap)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
