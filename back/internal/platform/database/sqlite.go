package database

import (
	"github.com/briangtn/codepic/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(dbFile string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domain.CodePic{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
