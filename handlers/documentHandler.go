package handlers

import (
	"archive-system/databaseProvaider"
	"archive-system/models"
	"archive-system/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadDocumentWithMeta(c *gin.Context) {
	author := c.PostForm("author")
	if author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствует автор документа"})
		return
	}

	fileHeader, err := c.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл не загружен"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка открытия файла"})
		return
	}
	defer file.Close()

	fileID, err := repositories.UploadPDF(file, fileHeader.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки в GridFS"})
		return
	}

	doc := models.Document{
		Filename: fileHeader.Filename,
		Author:   author,
		MongoIDs: []string{fileID.Hex()},
	}

	if err := repositories.CreateDocument(databaseProvaider.DB, &doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения в PostgreSQL"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":       "Документ успешно загружен",
		"document":      doc,
		"mongo_file_id": fileID.Hex(),
	})
}

func ListDocuments(c *gin.Context) {
	docs, err := repositories.GetAllDocuments(databaseProvaider.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения документов"})
		return
	}

	c.JSON(http.StatusOK, docs)
}
