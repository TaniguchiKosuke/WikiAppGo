package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
}

type Document struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Title   string
	Content string
	Author  string
}
