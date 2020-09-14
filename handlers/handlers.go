package handlers

import (
	"encoding/json"
	"fmt"
	"iap/helper"
	"iap/models"
	"iap/validators"
	"net/http"

	"github.com/fatih/structs"
)

//AppleVerify handler
func AppleVerify(w http.ResponseWriter, r *http.Request) {
	data := models.ReceiptData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
	}
	if ok, response := validators.Validate(data); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
	} else {
		jsonMap := structs.Map(data)
		X = &Apple{}
		fmt.Println(X.Verify(jsonMap))
		helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	}
}

//GoogleVerify handler
func GoogleVerify(w http.ResponseWriter, r *http.Request) {
	data := models.VerifySubscription{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
	}
	if ok, response := validators.Validate(data); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
	} else {
		jsonMap := structs.Map(data)
		fmt.Println(jsonMap)
		X = &Google{}
		fmt.Println(X.Verify(jsonMap))
		helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	}
}
