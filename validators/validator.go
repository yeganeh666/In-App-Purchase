package validators

import (
	"encoding/json"
	"iap/models"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

// Init initialize the validator.
func Init() {
	validate = validator.New()
}
func Validate(model interface{}) (bool, models.Response) {
	err := validate.Struct(model)
	targets := []models.Target{}
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			targets = append(targets, models.Target{
				Name:        err.Field(),
				Description: err.Tag(),
			})
		}
		errResponse, err := json.Marshal(targets)
		if err != nil {
			return false, models.ServerDownResponse
		}
		return false, models.Response{StatusCode: http.StatusBadRequest, Body: errResponse}
	}
	return true, models.Response{}
}
