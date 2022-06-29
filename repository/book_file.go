package repository

import (
	"fmt"
	"pustaka-api/models"
)

type bookFileRepository struct{}

func NewBookFileRepository() *bookFileRepository {
	return &bookFileRepository{}
}

func (b *bookFileRepository) FindAll() ([]models.Book, error) {
	books := []models.Book{}
	fmt.Println("FindAll")
	return books, nil
}

func (b *bookFileRepository) FindByID(id int) (models.Book, error) {
	book := models.Book{}
	fmt.Println("FindByID")
	return book, nil
}

func (b *bookFileRepository) Create(book models.Book) (models.Book, error) {
	fmt.Println("Create")
	return book, nil
}
