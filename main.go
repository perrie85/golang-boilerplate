package main

import (
	"golang-boilerplate/database"
	"golang-boilerplate/database/models"
	"golang-boilerplate/routes"
	"log"
	"net/http"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(&models.User{}, &models.List{})

	apiRoutes := routes.Api()
	log.Println("Listening On Port 80")
	http.ListenAndServe(":80", &apiRoutes)
}
