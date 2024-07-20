package list

import (
	"encoding/json"
	"golang-boilerplate/validators"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ListStore struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	UserID      int    `validate:"required" json:"user_id"`
}

func (rules ListStore) Assign(r *http.Request) ListStore {
	json.NewDecoder(r.Body).Decode(&rules)
	return rules
}

func (rules ListStore) Validate(w http.ResponseWriter, v *validator.Validate) error {
	err := v.Struct(&rules)
	if err != nil {
		validators.ValidationErrorHandler(w, err)
	}
	return err
}
