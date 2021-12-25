package database

import (
	"labqid/app/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Category{},
		&models.Product{},
	)
}
