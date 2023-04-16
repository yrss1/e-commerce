package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	rtr := gin.Default()
	booksRouter := rtr.Group("/books")
	userRouter := rtr.Group("")
	booksRouter.POST("/")

}
