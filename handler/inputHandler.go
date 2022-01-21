package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookInput struct {
	Id       json.Number `json:"id" binding:"required,number"`
	Title    string `json:"title" binding:"required"`
	SubTitle string `json:"sub_title" binding:"required"`
}

func BookInputHandler(c *gin.Context){
	var bookInput BookInput
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