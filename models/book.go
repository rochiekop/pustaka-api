package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       uint      `json:"price" binding:"required"`
	Rating      uint      `json:"rating" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
