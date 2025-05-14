package routes

import (
	"shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUploadRoutes registers all file upload related routes
func RegisterUploadRoutes(api *gin.RouterGroup, uploadHandler *handler.UploadHandler) {
	// Upload
	api.POST("/upload", uploadHandler.UploadFile)
	api.POST("/upload/batch", uploadHandler.BatchUploadFiles)
	api.POST("/upload/delete", uploadHandler.DeleteFile)
}
