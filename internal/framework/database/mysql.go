package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (db *gorm.DB, err error) {
	dsn := os.Getenv("MYSQL_DSN")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
