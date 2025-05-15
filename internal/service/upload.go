package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"strings"
	"time"
)

// UploadService handles business logic for file uploads
type UploadService struct {
	minioClient *minio.Client
}

// NewUploadService creates a new upload service
func NewUploadService() *UploadService {
	return &UploadService{
		minioClient: minio.GetClient(),
	}
}

// UploadFile uploads a file using MinIO if available, falls back to local storage
func (s *UploadService) UploadFile(file *multipart.FileHeader) (string, error) {
	cfg := config.GetConfig()

	// Check file size
	if file.Size > cfg.Upload.MaxSize {
		return "", fmt.Errorf("file size exceeds the limit: %d bytes", cfg.Upload.MaxSize)
	}

	// Get file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
		return "", errors.New("only image files are allowed")
	}

	// Generate unique filename
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// Open file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload to MinIO
	contentType := getMimeType(ext)
	objectName := fmt.Sprintf("uploads/%s", fileName)

	ctx := context.Background()
	err = s.minioClient.UploadFile(ctx, objectName, src, contentType)
	if err != nil {
		// Fall back to local storage if MinIO fails
		return s.uploadToLocalStorage(file, fileName, cfg.Upload.SavePath)
	}

	// Return MinIO file URL
	return s.minioClient.GetFileURL(objectName), nil
}

// uploadToLocalStorage uploads a file to local storage (fallback method)
func (s *UploadService) uploadToLocalStorage(file *multipart.FileHeader, fileName, savePath string) (string, error) {
	// Create save path if not exists
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Save file
	dst := filepath.Join(savePath, fileName)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	// Return file URL
	return "/uploads/" + fileName, nil
}

// UploadFileFromPath uploads a file from local path to MinIO
func (s *UploadService) UploadFileFromPath(localPath, objectName string) (string, error) {
	contentType := getMimeType(filepath.Ext(localPath))

	// Open file
	file, err := os.Open(localPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Upload to MinIO
	ctx := context.Background()
	err = s.minioClient.UploadFile(ctx, objectName, file, contentType)
	if err != nil {
		return "", err
	}

	// Return MinIO file URL
	return s.minioClient.GetFileURL(objectName), nil
}

// DeleteFile deletes a file from storage
func (s *UploadService) DeleteFile(fileURL string) error {
	// Check if it's a MinIO URL
	if strings.Contains(fileURL, s.minioClient.GetFileURL("")) {
		// Extract object name from URL
		parts := strings.Split(fileURL, "/")
		objectName := parts[len(parts)-1]

		// Delete from MinIO
		ctx := context.Background()
		return s.minioClient.DeleteFile(ctx, "uploads/"+objectName)
	}

	// If not MinIO URL, treat as local file
	if strings.HasPrefix(fileURL, "/uploads/") {
		cfg := config.GetConfig()
		localPath := filepath.Join(cfg.Upload.SavePath, filepath.Base(fileURL))

		// Delete local file
		return os.Remove(localPath)
	}

	return errors.New("invalid file URL format")
}

// getMimeType returns the MIME type for a file extension
func getMimeType(ext string) string {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}
