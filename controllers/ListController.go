package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/database/models"
	"todo/services"

	"github.com/gorilla/mux"
)

func ListIndex(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	services.ListIndex(query)

	json.NewEncoder(w).Encode(services.ListIndex(query))
}

func ListStore(w http.ResponseWriter, r *http.Request) {
	var list models.List

	json.NewDecoder(r.Body).Decode(&list)

	json.NewEncoder(w).Encode(services.ListStore(list))
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
