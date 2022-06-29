package repository

import (
	"pustaka-api/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll()([]models.Book, error)
	FindByID(id int)(models.Book, error)
	Create(book models.Book)(models.Book, error)
	Update(book models.Book)(models.Book, error)
	Delete(book models.Book)(models.Book, error)
}

type repository struct{
	db *gorm.DB
}


func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


func (r *repository) FindAll()([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) FindByID(id int)(models.Book, error) {
	var books models.Book
	err := r.db.First(&books, id).Error
	return books, err
}

func (r *repository) Create(book models.Book)(models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book models.Book)(models.Book, error){
	err := r.db.Save(&book).Error
	return book, err
}
func (r *repository) Delete(book models.Book)(models.Book, error){
	err := r.db.Delete(&book).Error
	return book, err
}