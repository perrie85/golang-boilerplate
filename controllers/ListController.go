package controllers

import (
	"encoding/json"
	"golang-boilerplate/services"
	"golang-boilerplate/validators/list"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

func ListIndex(w http.ResponseWriter, r *http.Request) {
	var params list.ListIndex
	params = params.Assign(r)

	validate = validator.New(validator.WithRequiredStructEnabled())
	err := params.Validate(w, validate)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(services.ListIndex(params))
}

func ListStore(w http.ResponseWriter, r *http.Request) {
	var params list.ListStore
	params = params.Assign(r)

	validate = validator.New(validator.WithRequiredStructEnabled())
	err := params.Validate(w, validate)
	if err != nil {
		return
	}

	json.NewEncoder(w).Encode(services.ListStore(params))
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

	var params list.ListUpdate
	params = params.Assign(r)

	validate = validator.New(validator.WithRequiredStructEnabled())
	validationErr := params.Validate(w, validate)
	if validationErr != nil {
		return
	}

	json.NewEncoder(w).Encode(services.ListUpdate(id, params))
}

func ListDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(services.ListDelete(id))
}
