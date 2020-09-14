package helper

import (
	"net/http"
)

func HttpResponse(w http.ResponseWriter, responseStatus int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(response)
}
