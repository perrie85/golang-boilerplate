package middlewares

import (
	"log"
	"net/http"
)

func RequestLogging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method + " " + r.URL.Path)
			f(w, r)
		}
	}
}
