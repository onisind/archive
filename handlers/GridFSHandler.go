package handlers

import (
	"archive-system/databaseProvaider"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// загрузка пдф в GridFS
func UploadPDFToGridFS(file multipart.File, filename string) (primitive.ObjectID, error) {
	uploadStream, err := databaseProvaider.MongoBucket.OpenUploadStream(filename)
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	_, err = io.Copy(uploadStream, file)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return uploadStream.FileID.(primitive.ObjectID), nil
}

// скачивание пдф с GridFS
func DownloadPDFFromGridFS(fileID primitive.ObjectID) ([]byte, error) {
	downloadStream, err := databaseProvaider.MongoBucket.OpenDownloadStream(fileID)
	if err != nil {
		return nil, err
	}
	defer downloadStream.Close()

	return io.ReadAll(downloadStream)
}

// скачивание пдф по MongoID
func DownloadMongoID(c *gin.Context) {
	idHex := c.Param("id")
	fileID, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ObjectID"})
		return
	}

	fileData, err := DownloadPDFFromGridFS(fileID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Файл не найден"})
		return
	}

	c.Data(http.StatusOK, "application/pdf", fileData)
}
