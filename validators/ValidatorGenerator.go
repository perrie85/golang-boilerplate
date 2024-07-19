package validators

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Assign(r *http.Request)
	Validate(v *validator.Validate) error
}

func ValidationErrorHandler(w http.ResponseWriter, err error) {
	var response []interface{}

	for _, err := range err.(validator.ValidationErrors) {

		row := []string{
			err.Namespace(),
			err.Field(),
			err.StructNamespace(),
			err.StructField(),
			err.Tag(),
			err.ActualTag(),
			// err.Kind(),
			// err.Type(),
			// err.Value(),
			// err.Param(),
		}
		response = append(response, row)
	}

	json.NewEncoder(w).Encode(response)
}
