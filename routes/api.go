package routes

import (
	"log"
	"net/http"
	"todo/controllers"

	"github.com/gorilla/mux"
)

func Api() mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	listRouter := api.PathPrefix("/list").Subrouter()

	listRouter.HandleFunc("", requestLogging(controllers.ListIndex)).Methods("GET")
	listRouter.HandleFunc("", requestLogging(controllers.ListStore)).Methods("POST")
	listRouter.HandleFunc("/{id}", requestLogging(controllers.ListShow)).Methods("GET")
	listRouter.HandleFunc("/{id}", requestLogging(controllers.ListUpdate)).Methods("PUT")
	listRouter.HandleFunc("/{id}", requestLogging(controllers.ListDelete)).Methods("DELETE")

	return *r
}

func requestLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + " " + r.URL.Path)
		f(w, r)
	}
}
