package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/models"
	"pustaka-api/response"
	"pustaka-api/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService service.Service
}

func NewBookHandler(service service.Service) *bookHandler {
	return &bookHandler{service}
}

func (handler *bookHandler) GetBooks(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	var booksResponse []response.BookResponse

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	for _, book := range books {
		bookResponse := convertResponse(book)

		booksResponse = append(booksResponse, bookResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "success",
		"data":   booksResponse,
		"status": http.StatusOK,
	})
}


func (handler *bookHandler) GetBookById(c *gin.Context){
	var book models.Book

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	book, err := handler.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	bookResponse := convertResponse(book)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   bookResponse,
	})
}



func (handler *bookHandler) PostBooksHandler(c *gin.Context) {
	var book service.InputBook
	err := c.ShouldBindJSON(&book)
	if err != nil {
		errorMsgs := []string{}
		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", v.Field(), v.ActualTag())
			errorMsgs = append(errorMsgs, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	newBook, err := handler.bookService.Create(book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   convertResponse(newBook),
	})
}


func (handler *bookHandler) PutBooksHandler(c *gin.Context) {
	var book service.InputBook
	err := c.ShouldBindJSON(&book)
	if err != nil {
		errorMsgs := []string{}
		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", v.Field(), v.ActualTag())
			errorMsgs = append(errorMsgs, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMsgs,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	newBook, err := handler.bookService.Update(id,book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   convertResponse(newBook),
	})
}



func (handler *bookHandler) DeleteBooksHandler(c *gin.Context) {

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	newBook, err := handler.bookService.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   convertResponse(newBook),
	})
}


func convertResponse(book models.Book) response.BookResponse{
	bookResponse := response.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Rating:      book.Rating,
	}
	return bookResponse
}