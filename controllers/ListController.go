package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(services.ListShow(id))
}

func ListUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	var bodyDecoded models.List

	json.NewDecoder(r.Body).Decode(&bodyDecoded)

	json.NewEncoder(w).Encode(services.ListUpdate(id, bodyDecoded))
}

func ListDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(services.ListDelete(id))
}
