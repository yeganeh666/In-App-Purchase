package handlers

import (
	"encoding/json"
	"iap/helper"
	"iap/services/apple"
	"iap/services/google"
	"iap/validators"
	"net/http"

	"github.com/gorilla/mux"
)

func Verify(w http.ResponseWriter, r *http.Request) {
	var platforms = map[string]Services{
		"apple":  &apple.Apple{},
		"google": &google.Google{},
	}
	platform, ok := platforms[mux.Vars(r)["provider"]]

	if !ok {
		helper.HttpResponse(w, http.StatusBadRequest, []byte("ERROR!"))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&platform)
	if err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	if ok, response := validators.Validate(platform); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
		return
	}
	res := platform.Verify()
	if res["error"] != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(res["error"].(string)))
		return
	}
	helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	return
}

func AcknowledgeSubscription(w http.ResponseWriter, r *http.Request) {
	platform := google.Google{}
	err := json.NewDecoder(r.Body).Decode(&platform)
	if err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	if ok, response := validators.Validate(platform); !ok {
		helper.HttpResponse(w, response.StatusCode, response.Body)
		return
	}
	if err = platform.AcknowledgeSubscription(); err != nil {
		helper.HttpResponse(w, http.StatusBadRequest, []byte(err.Error()))
		return
	}
	helper.HttpResponse(w, http.StatusOK, []byte("OK!"))
	return
}
