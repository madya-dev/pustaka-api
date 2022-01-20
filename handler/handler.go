package handler

import "github.com/gin-gonic/gin"

func MainHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func ParamHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}