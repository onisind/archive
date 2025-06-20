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

func UpdateDocument(c *gin.Context) {
	id := c.Param("id")

	doc, err := repositories.GetDocumentByID(databaseProvaider.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Документ не найден"})
		return
	}

	if newAuthor := c.PostForm("author"); newAuthor != "" {
		doc.Author = newAuthor
	}

	if newFilename := c.PostForm("filename"); newFilename != "" {
		doc.Filename = newFilename
	}

	fileHeader, err := c.FormFile("pdf")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка открытия файла"})
			return
		}
		defer file.Close()

		fileID, err := repositories.UploadPDF(file, fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка загрузки файла в GridFS"})
			return
		}

		doc.MongoIDs = append(doc.MongoIDs, fileID.Hex())
	}

	if err := repositories.UpdateDocument(databaseProvaider.DB, doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления документа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Документ обновлён",
		"document": doc,
	})
}

func DeleteDocument(c *gin.Context) {
	id := c.Param("id")

	doc, err := repositories.GetDocumentByID(databaseProvaider.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Документ не найден"})
		return
	}

	if err := repositories.DeleteDocument(databaseProvaider.DB, doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления документа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Документ удалён"})
}

func ListDocuments(c *gin.Context) {
	docs, err := repositories.GetAllDocuments(databaseProvaider.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения документов"})
		return
	}

	c.JSON(http.StatusOK, docs)
}
