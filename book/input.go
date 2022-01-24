package book

import "encoding/json"

type BookInput struct {
	Id       json.Number `json:"id" binding:"required,number"`
	Title    string      `json:"title" binding:"required"`
	SubTitle string      `json:"sub_title" binding:"required"`
}
