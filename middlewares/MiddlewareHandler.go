package middlewares

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddlewares(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
