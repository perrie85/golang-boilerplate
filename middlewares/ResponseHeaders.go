package middlewares

import (
	"net/http"
)

type Headers struct {
	Name  string
	Value string
}

func ResponseHeaders(headers []Headers) Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			for _, header := range headers {
				w.Header().Set(header.Name, header.Value)
			}
			f(w, r)
		}
	}
}
