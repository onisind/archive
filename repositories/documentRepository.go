package repositories

import (
	"archive-system/models"

	"gorm.io/gorm"
)

func CreateDocument(db *gorm.DB, doc *models.Document) error {
	return db.Create(doc).Error
}

func GetAllDocuments(db *gorm.DB) ([]models.Document, error) {
	var docs []models.Document
	err := db.Find(&docs).Error
	return docs, err
}

func GetDocumentByID(db *gorm.DB, id string) (*models.Document, error) {
	var doc models.Document
	if err := db.First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func UpdateDocument(db *gorm.DB, doc *models.Document) error {
	return db.Save(doc).Error
}

func DeleteDocument(db *gorm.DB, doc *models.Document) error {
	return db.Delete(doc).Error
}
