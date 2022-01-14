package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Title    string
	Content  string
	AuthorID string
	Author   User   `gorm:"foreignKey:AuthorID"`
}
