package list

import (
	"golang-boilerplate/validators"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ListIndex struct {
	Title       string   `validate:"required" json:"title"`
	Description string   `validate:"required" json:"description"`
	UserID      int64    `validate:"required" json:"user_id"`
	CreatedAt   string   `validate:"datetime" json:"created_at"`
	OrderBy     *orderBy `json:"order_by"`
}

type orderBy struct {
	By        string
	Direction string
}

func (rules ListIndex) Assign(r *http.Request) ListIndex {
	queryVals := r.URL.Query()

	rules.Title = queryVals.Get("title")
	rules.Description = queryVals.Get("description")
	rules.UserID, _ = strconv.ParseInt(queryVals.Get("user_id"), 10, 64)
	rules.CreatedAt = queryVals.Get("created_at")
	orderBy := orderBy{
		By:        queryVals.Get("order_by[by]"),
		Direction: queryVals.Get("order_by[direction]"),
	}
	rules.OrderBy = &orderBy

	return rules
}

func (rules ListIndex) Validate(w http.ResponseWriter, v *validator.Validate) error {
	err := v.Struct(&rules)
	if err != nil {
		validators.ValidationErrorHandler(w, err)
	}
	return err
}
