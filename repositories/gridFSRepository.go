package repositories

import (
	"archive-system/databaseProvaider"
	"io"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Загрузка PDF в GridFS
func UploadPDF(file multipart.File, filename string) (primitive.ObjectID, error) {
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

// Загрузка PDF из GridFS по ObjectID
func DownloadPDF(fileID primitive.ObjectID) ([]byte, error) {
	downloadStream, err := databaseProvaider.MongoBucket.OpenDownloadStream(fileID)
	if err != nil {
		return nil, err
	}
	defer downloadStream.Close()

	return io.ReadAll(downloadStream)
}
