package controllers

import (
	"fmt"
	"net/http"
	"todo/services"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	processed := services.HelloWorld(query)

	fmt.Fprintf(w, "Hello, you've requested: %s\n", processed)
}
