package database

import (
	"belajar-go-echo/internal/core/models"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) (err error) {
	err = db.AutoMigrate(models.User{})
	return
}
