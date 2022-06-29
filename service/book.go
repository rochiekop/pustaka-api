package service

import (
	"encoding/json"
	"pustaka-api/models"
	"pustaka-api/repository"
)

type InputBook struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`
}

type Service interface {
	FindAll() ([]models.Book, error)
	FindByID(id int) (models.Book, error)
	Create(inputBook InputBook) (models.Book, error)
	Update(id int, inputBook InputBook) (models.Book, error)
	Delete(id int) (models.Book, error)

}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]models.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(id int) (models.Book, error) {
	book, err := s.repository.FindByID(id)
	return book, err
}

func (s *service) Create(inputBook InputBook) (models.Book, error) {
	price, _ := inputBook.Price.Int64()
	rating, _ := inputBook.Rating.Int64()
	book := models.Book{
		Title:       inputBook.Title,
		Description: inputBook.Description,
		Price:       uint(price),
		Rating:      uint(rating),
	}
	book, err := s.repository.Create(book)
	return book, err
}

func (s *service) Update(id int, inputBook InputBook) (models.Book, error) {
	book, err := s.repository.FindByID(id)

	if err != nil {
		return book, err
	}

	price, _ := inputBook.Price.Int64()
	rating, _ := inputBook.Rating.Int64()
	book.Title = inputBook.Title
	book.Description = inputBook.Description
	book.Price = uint(price)
	book.Rating = uint(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}


func (s *service) Delete(id int) (models.Book, error) {
	book, err := s.repository.FindByID(id)

	if err != nil {
		return book, err
	}

	newBook, err := s.repository.Delete(book)
	return newBook, err
}
