package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open the database")
	}

	return db
}

func init() {
	db := dbConnect()
	db.AutoMigrate(&Document{})
}
