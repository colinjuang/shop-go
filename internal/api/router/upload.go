package router

import (
	"github.com/colinjuang/shop-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUploadApi registers all file upload related api
func RegisterUploadApi(api *gin.RouterGroup) {
	uploadHandler := handler.NewUploadHandler()
	// Upload
	api.POST("/upload", uploadHandler.UploadFile)
	api.POST("/upload/batch", uploadHandler.BatchUploadFiles)
	api.POST("/upload/delete", uploadHandler.DeleteFile)
}
