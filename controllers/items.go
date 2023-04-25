package controllers

import (
	"e-commerce-golang/database"
	"e-commerce-golang/models"
)

func GetItems() []models.Item {
	var items []models.Item
	res := database.GetDB().Find(&items)
	if res.Error != nil {
		panic(res.Error)
	}
	return items
}

func GetItemById(id string) (models.Item, error) {
	var item models.Item
	res := database.GetDB().Find(&item, models.Item{Id: id})
	return item, res.Error
}

func GetSortedItemsByRating(order string) []models.Item {
	var items []models.Item
	if order == "asc" {
		database.GetDB().Order("price asc, rating").Find(&items)
	} else if order == "desc" {
		database.GetDB().Order("price desc, rating").Find(&items)
	}
	return items
}

func GetSortedItemsByPrice(order string) []models.Item {
	var items []models.Item
	if order == "asc" {
		database.GetDB().Order("price asc, price").Find(&items)
	} else if order == "desc" {
		database.GetDB().Order("price desc, price").Find(&items)
	}
	return items
}
