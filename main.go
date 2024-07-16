package main

import (
	"log"
	"net/http"
	"todo/database"
	"todo/database/models"
	"todo/routes"
)

func main() {
	db := database.Connect()
	db.AutoMigrate(&models.User{}, &models.List{})

	apiRoutes := routes.Api()
	log.Println("Listening On Port 80")
	http.ListenAndServe(":80", &apiRoutes)
}
