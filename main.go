package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.MainHandler)
	v1.POST("/book", handler.BookInputHandler)
	v1.GET("/:name", handler.ParamHandler)
	
	router.Run()
}