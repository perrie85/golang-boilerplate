package list

import (
	"encoding/json"
	"golang-boilerplate/validators"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ListUpdate struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	UserId      int    `validate:"required" json:"user_id"`
}

func (rules ListUpdate) Assign(r *http.Request) ListUpdate {
	json.NewDecoder(r.Body).Decode(&rules)
	return rules
}

func (rules ListUpdate) Validate(w http.ResponseWriter, v *validator.Validate) error {
	err := v.Struct(&rules)
	if err != nil {
		validators.ValidationErrorHandler(w, err)
	}
	return err
}
