package routes

import (
	"archive-system/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/documents/upload", handlers.UploadDocumentWithMeta)
	r.GET("/documents", handlers.ListDocuments)
}
