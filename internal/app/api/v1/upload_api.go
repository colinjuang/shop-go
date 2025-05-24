package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUploadApi registers all file upload related api
func RegisterUploadApi(router *gin.Engine) {
	uploadHandler := handler.NewUploadHandler()
	api := router.Group("/api")
	{
		// Upload
		api.POST("/upload", uploadHandler.UploadFile)
		api.POST("/upload/batch", uploadHandler.BatchUploadFiles)
		api.POST("/upload/delete", uploadHandler.DeleteFile)
	}
}
