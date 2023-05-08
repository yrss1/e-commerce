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
	//reg auth
	userRouter.POST("/registration", handlers.Registration)
	userRouter.POST("/authorisation", handlers.Authorisation)
	//items
	itemsRouter.GET("", handlers.GetItems)
	itemsRouter.GET("/:id", handlers.GetItem)
	itemsRouter.GET("/sort", handlers.Filter)
	//itemsRouter.POST("/:id/rate", handlers.GiveRating)
	itemsRouter.POST("/create_item", handlers.AddItem)
	itemsRouter.DELETE("/:id/delete", handlers.DeleteItem)
	itemsRouter.GET("/:id/comments", handlers.GetCommentaries)
	itemsRouter.POST("/create_comment", handlers.SetComment)

	//get all comments
	itemsRouter.GET("/comments", handlers.GetComments)
	//get item with comments
	itemsRouter.GET("/:id/with_comments", handlers.GetItemsWithComments)
	rtr.Run("localhost:8088")

}

//rating
//comment
//4) Giving rating for items -> sql queries
//5) Commenting items -> sql queries
// 7) Purchasing item -> sql queries
