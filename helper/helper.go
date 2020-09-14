package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func StructToByte(model interface{}) []byte {
	responseBytes := new(bytes.Buffer)
	json.NewEncoder(responseBytes).Encode(model)
	return responseBytes.Bytes()
}
func HttpResponse(w http.ResponseWriter, responseStatus int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(response)
}
