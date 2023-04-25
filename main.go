package main

import (
	"e-commerce-golang/database"
	"e-commerce-golang/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	rtr := gin.Default()

	userRouter := rtr.Group("")
	itemsRouter := rtr.Group("/items")

	userRouter.POST("/registration", handlers.Registration)
	userRouter.POST("/authorisation", handlers.Authorisation)

	itemsRouter.GET("", handlers.GetItems)
	itemsRouter.GET("/:id", handlers.GetItem)
	itemsRouter.GET("/sort", handlers.Filter)
	rtr.Run("localhost:8088")

}

//done
//reg auth
//filter price rating ошибка с float number 3.5 > 5

//to do
// give rating
// commenting
// publishing
//purchasing
