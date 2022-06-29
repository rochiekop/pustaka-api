package config

import (
	"log"
	"pustaka-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	dsn := "root@tcp(127.0.0.1:3306)/database_pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal()
	}
	db.AutoMigrate(&models.Book{})

	return db, err
}
