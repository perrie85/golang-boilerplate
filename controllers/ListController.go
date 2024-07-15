package controllers

import (
	"fmt"
	"net/http"
	"todo/services"

	"github.com/gorilla/mux"
)

func ListIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	processed := services.HelloWorld(query)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", processed)
}

func ListStore(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	processed := services.HelloWorld(query)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", processed)
}

func ListShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	// processed := services.HelloWorld(id)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", id)
}

func ListUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	// processed := services.HelloWorld(id)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", id)
}

func ListDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	// processed := services.HelloWorld(id)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", id)
}
