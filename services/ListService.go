package services

import (
	"golang-boilerplate/database"
	"golang-boilerplate/database/models"
	"net/url"
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

func ListShow(id int64) models.List {
	db := database.Connect()

	var list models.List

	db.Model(&models.List{}).First(&list, id)

	return list
}

func ListUpdate(id int64, params models.List) models.List {
	db := database.Connect()

	var list models.List

	db.Model(&models.List{}).First(&list, id)

	list.Title = params.Title
	list.Description = params.Description

	db.Save(&list)
	return list
}

func ListDelete(id int64) models.List {
	db := database.Connect()

	var list models.List

	db.Model(&models.List{}).Delete(&list, id)

	return list
}
