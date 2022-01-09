package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("cannot open the database")
	}

	return db
}

func init() {
	db := DbConnect()
	db.AutoMigrate(&Document{}, &User{})
}
