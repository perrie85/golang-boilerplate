package routes

import (
	"todo/controllers"

	"github.com/gorilla/mux"
)

func Api() mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	listRouter := api.PathPrefix("/list").Subrouter()

	listRouter.HandleFunc("", controllers.ListIndex).Methods("GET")
	listRouter.HandleFunc("", controllers.ListStore).Methods("POST")
	listRouter.HandleFunc("/{id}", controllers.ListShow).Methods("GET")
	listRouter.HandleFunc("/{id}", controllers.ListUpdate).Methods("PUT")
	listRouter.HandleFunc("/{id}", controllers.ListDelete).Methods("DELETE")

	return *r
}
