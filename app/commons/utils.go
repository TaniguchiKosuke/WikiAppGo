package commons

import (
	"WikiAppGo/app/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetDocumentsList(user models.User, db *gorm.DB) []models.Document {
	var documents []models.Document
	db.Preload(clause.Associations).Where("author_id", user.ID).Order("created_at desc").Limit(8).Find(&documents)
	return documents
}