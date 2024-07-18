package list

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ListIndex struct {
	Title       string   `validate:"required" json:"title"`
	Description string   `validate:"required" json:"description"`
	UserId      int64    `validate:"required" json:"user_id"`
	CreatedAt   string   `validate:"required,datetime" json:"created_at"`
	OrderBy     *orderBy `json:"order_by"`
}

type orderBy struct {
	By        string
	Direction string
}

func (rules ListIndex) Assign(r *http.Request) {
	queryVals := r.URL.Query()

	rules.Title = queryVals.Get("title")
	rules.Description = queryVals.Get("title")
	rules.UserId, _ = strconv.ParseInt(queryVals.Get("user_id"), 10, 64)
	rules.CreatedAt = queryVals.Get("created_at")
	// rules.OrderBy = queryVals.Get("order_by")
}

func (rules ListIndex) Validate(v *validator.Validate) error {
	return v.Struct(&rules)
}
