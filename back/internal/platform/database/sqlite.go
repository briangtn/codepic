package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbFile string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate()
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
