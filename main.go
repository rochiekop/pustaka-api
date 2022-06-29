package main

import (
	"fmt"
	"pustaka-api/handler"
	"pustaka-api/repository"
	"pustaka-api/config"
	"pustaka-api/service"
	_ "time"

	"github.com/gin-gonic/gin"
)

func main() {

	db, _ := config.ConnectToDb()
	bookRepository := repository.NewRepository(db)
	// bookFileRepository := repository.NewBookFileRepository()
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	fmt.Println("Connected Database Success")
	route := gin.Default()

	v1 := route.Group("v1")
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBookById)
	v1.PUT("/books/:id", bookHandler.PutBooksHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBooksHandler)
	route.Run(":8888")
}

// Layer
// 1. Main
// 2. Handler
// 3. Service
// 4. Repository
// 5. Models
// 6. Database
