package services

import (
	"net/url"
	"todo/database"
	"todo/database/models"
)

func ListIndex(params url.Values) []models.List {
	db := database.Connect()

	var lists []models.List

	db.Model(&models.List{}).Limit(10).Find(&lists)

	return lists
}

func ListStore(params models.List) models.List {
	db := database.Connect()

	list := models.List{
		Title:       params.Title,
		Description: params.Description,
	}

	db.Create(&list)

	return list
}
