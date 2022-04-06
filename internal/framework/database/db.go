package database

import "gorm.io/gorm"

func InitDatabase() (db *gorm.DB) {
	var err error
	db, err = initSQLite()
	if err != nil {
		panic(err)
	}
	err = migrateDB(db)
	if err != nil {
		panic(err)
	}

	return db
}
