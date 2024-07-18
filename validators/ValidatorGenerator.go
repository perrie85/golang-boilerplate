package validators

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Assign(r *http.Request)
	Validate(v *validator.Validate) error
}
