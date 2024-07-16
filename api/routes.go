package api

import (
	"bookstore/config"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func StartServer() {
	r := gin.Default()
	RegisterRoutes(r)
	Router = r
	r.Run(":" + config.Port)
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/books", GetAllBooks)
	r.GET("/books/:id", GetBook)
	r.POST("/books", CreateBook)
	r.PUT("/books/:id", UpdateBook)
	r.DELETE("/books/:id", DeleteBook)
}
