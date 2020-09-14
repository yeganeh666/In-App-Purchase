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
		return
	}
	if ok, response := validators.Validate(data); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
		return
	}
	jsonMap := structs.Map(data)
	X = &Apple{}
	res := X.Verify(jsonMap)
	if res["error"] != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(res["error"].(string)))
		return
	}
	helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	return
}

//GoogleVerify handler
func GoogleVerify(w http.ResponseWriter, r *http.Request) {
	data := models.VerifySubscription{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	if ok, response := validators.Validate(data); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
		return
	}
	jsonMap := structs.Map(data)
	fmt.Println(jsonMap)
	X = &Google{}
	res := X.Verify(jsonMap)
	if res["error"] != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(res["error"].(string)))
		return
	}
	helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	return

}
