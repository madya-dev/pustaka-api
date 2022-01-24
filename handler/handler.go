package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func MainHandler(c *gin.Context) {
	c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
func ParamHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + name,
	})
}

func BookInputHandler(c *gin.Context){
	var bookInput book.BookInput
	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Id": bookInput.Id,
		"Title": bookInput.Title,
		"SubTitle": bookInput.SubTitle,
	})
}