package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterUploadApi registers all file upload related api
func RegisterUploadApi(router *gin.Engine, db *gorm.DB) {
	uploadHandler := handler.NewUploadHandler()
	api := router.Group("/api")
	{
		// Upload
		api.POST("/upload", uploadHandler.UploadFile)
		api.POST("/upload/batch", uploadHandler.BatchUploadFiles)
		api.POST("/upload/delete", uploadHandler.DeleteFile)
	}
}
