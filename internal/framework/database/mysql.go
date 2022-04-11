package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (db *gorm.DB, err error) {
		dsn, avail := os.LookupEnv("MYSQL_DSN")
		if !avail {
				dsn = "root:animelovers@tcp(localhost:3306)/crud_go?charset=utf8&parseTime=True&loc=Local"
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		return
}
