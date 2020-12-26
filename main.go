package main

import (
	"gin-db/controllers"
	"gin-db/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetOneBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":8080")
}
