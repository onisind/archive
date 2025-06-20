package repositories

import (
	"archive-system/models"

	"gorm.io/gorm"
)

// Сохраняет метаданные документа в Postgres
func CreateDocument(db *gorm.DB, doc *models.Document) error {
	return db.Create(doc).Error
}

// Получает все документы из Postgres
func GetAllDocuments(db *gorm.DB) ([]models.Document, error) {
	var docs []models.Document
	err := db.Find(&docs).Error
	return docs, err
}
