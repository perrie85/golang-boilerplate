package routes

import (
	"golang-boilerplate/controllers"
	"golang-boilerplate/middlewares"

	"github.com/gorilla/mux"
)

var apiMiddlewares []middlewares.Middleware
var apiResponseHeaders []middlewares.Headers

func Api() mux.Router {
	r := mux.NewRouter()

	setHeaders()
	setMiddlewares()

	api := r.PathPrefix("/api").Subrouter()

	listRouter := api.PathPrefix("/list").Subrouter()

	listRouter.HandleFunc("", middlewares.ApplyMiddlewares(controllers.ListIndex, apiMiddlewares...)).Methods("GET")
	listRouter.HandleFunc("", middlewares.ApplyMiddlewares(controllers.ListStore, apiMiddlewares...)).Methods("POST")
	listRouter.HandleFunc("/{id}", middlewares.ApplyMiddlewares(controllers.ListShow, apiMiddlewares...)).Methods("GET")
	listRouter.HandleFunc("/{id}", middlewares.ApplyMiddlewares(controllers.ListUpdate, apiMiddlewares...)).Methods("PUT")
	listRouter.HandleFunc("/{id}", middlewares.ApplyMiddlewares(controllers.ListDelete, apiMiddlewares...)).Methods("DELETE")

	return *r
}

func setHeaders() {
	apiResponseHeaders = append(apiResponseHeaders, middlewares.Headers{Name: "Content-Type", Value: "application/json"})
}

func setMiddlewares() {
	apiMiddlewares = append(apiMiddlewares, middlewares.ResponseHeaders(apiResponseHeaders))
	apiMiddlewares = append(apiMiddlewares, middlewares.RequestLogging())
}
