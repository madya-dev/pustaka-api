package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("DB Connection Error")
	}
	log.Println("DB Connection Success")
	db.AutoMigrate(&book.Book{})

	// ----CREATE DATA-----
	// book := book.Book{Title: "The Amazing Spiderman 2", Description: "Next Chapter", Price: 15000, Rating: 4}
	// result := db.Create(&book)
	// fmt.Println(result)

	// -----QUERY DATA-------
	// book := book.Book{};
	// result := db.Find(&book);
	// fmt.Println(result)
	// fmt.Println(book)
	// fmt.Println("===========================")
	// log.Println("Title : ",book.Title)
	// fmt.Println("===========================")

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/", handler.MainHandler)
	v1.POST("/book", handler.BookInputHandler)
	v1.GET("/:name", handler.ParamHandler)
	
	router.Run()
}