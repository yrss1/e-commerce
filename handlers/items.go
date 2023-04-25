package handlers

import (
	"e-commerce-golang/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetItems(cnt *gin.Context) {
	items := controllers.GetItems()
	cnt.IndentedJSON(http.StatusOK, items)
}

func GetItem(cnt *gin.Context) {
	id := cnt.Param("id")
	item, err := controllers.GetItemById(id)
	if err != nil {
		cnt.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
		return
	}
	cnt.IndentedJSON(http.StatusOK, item)
}

func Filter(cnt *gin.Context) {
	based := cnt.Query("based")
	order := cnt.Query("order")
	if based == "price" {
		items := controllers.GetSortedItemsByPrice(order)
		cnt.IndentedJSON(http.StatusOK, items)
	} else if based == "rating" {
		items := controllers.GetSortedItemsByRating(order)
		cnt.IndentedJSON(http.StatusOK, items)
	} else {
		cnt.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
	}
}
