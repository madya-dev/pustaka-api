package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handler.MainHandler)
	router.GET("/:name", handler.ParamHandler)
	router.Run()
}