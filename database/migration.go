package database

import (
	"fiber-gorm/models"

	"gorm.io/gorm"
)

var (
	err error
	db  *gorm.DB
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{})
}
