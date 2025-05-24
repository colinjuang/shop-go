package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
)

// UploadHandler handles file upload API endpoints
type UploadHandler struct {
	uploadService *service.UploadService
}

// NewUploadHandler creates a new upload handler
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{
		uploadService: service.NewUploadService(),
	}
}

// UploadFile handles file upload
func (h *UploadHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Missing file"))
		return
	}

	fileURL, err := h.uploadService.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"url": fileURL,
	}))
}

// BatchUploadFiles handles multiple file uploads
func (h *UploadHandler) BatchUploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid form data"))
		return
	}

	files := form.File["files"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "No files provided"))
		return
	}

	urls := make([]string, 0, len(files))
	for _, file := range files {
		fileURL, err := h.uploadService.UploadFile(file)
		if err != nil {
			// Log error but continue with other files
			logger.Errorf("Failed to upload file %s: %v", file.Filename, err)
			continue
		}
		urls = append(urls, fileURL)
	}

	logger.Infof("Successfully uploaded %d files", len(urls))
	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"urls": urls,
	}))
}

// DeleteFile handles file deletion
func (h *UploadHandler) DeleteFile(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid request"))
		return
	}

	err := h.uploadService.DeleteFile(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}
