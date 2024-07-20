package services

import (
	"golang-boilerplate/database"
	"golang-boilerplate/database/models"
	"golang-boilerplate/validators/list"
)

func ListIndex(params list.ListIndex) []models.List {
	db := database.Connect()

	var lists []models.List

	db.Model(&models.List{}).Preload("User").Limit(10).Find(&lists)

	return lists
}

func ListStore(params list.ListStore) models.List {
	db := database.Connect()

	list := models.List{
		Title:       params.Title,
		Description: params.Description,
		UserID:      params.UserID,
	}

	db.Create(&list)

	return list
}

func ListShow(id int64) models.List {
	db := database.Connect()

	var list models.List

	db.Model(&models.List{}).Preload("User").First(&list, id)

	return list
}

func ListUpdate(id int64, params list.ListUpdate) models.List {
	db := database.Connect()

	var list models.List

	db.Model(&models.List{}).Preload("User").First(&list, id)

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
