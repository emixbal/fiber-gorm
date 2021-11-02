package database

import (
	"fiber-gorm/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Book{})
}
