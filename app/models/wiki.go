package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password []byte
}

type Document struct {
	gorm.Model
	ID      string `gorm:"primaryKey"`
	Title   string
	Content string
	AuthorID  string
	Author  User
}
