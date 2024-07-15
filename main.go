package main

import (
	"net/http"
	"todo/database"
	"todo/database/models"
	"todo/routes"
)

func main() {

	db := database.Connect()
	db.AutoMigrate(&models.User{}, &models.List{})

	apiRoutes := routes.Api()
	http.ListenAndServe(":80", &apiRoutes)
}
