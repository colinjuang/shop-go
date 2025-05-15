package minio

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"github.com/colinjuang/shop-go/internal/config"
	"strings"
)

// MigrationResult contains the results of a migration operation
type MigrationResult struct {
	TotalFiles    int
	SuccessCount  int
	FailedCount   int
	FailedFiles   []string
	MigratedFiles []string
}

// MigrateLocalToMinIO migrates files from local storage to MinIO
func MigrateLocalToMinIO(ctx context.Context) (*MigrationResult, error) {
	cfg := config.GetConfig()
	client := GetClient()

	// Result
	result := &MigrationResult{
		FailedFiles:   make([]string, 0),
		MigratedFiles: make([]string, 0),
	}

	// Create uploads directory if not exists
	uploadPath := cfg.Upload.SavePath
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		return result, nil // No local files to migrate
	}

	// Walk through all files in the uploads directory
	err := filepath.Walk(uploadPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process image files
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return nil
		}

		// Get relative path
		relPath, err := filepath.Rel(uploadPath, path)
		if err != nil {
			return err
		}

		// Increment total
		result.TotalFiles++

		// Check if file already exists in MinIO
		objectName := fmt.Sprintf("uploads/%s", relPath)
		exists, err := client.FileExists(ctx, objectName)
		if err != nil {
			result.FailedCount++
			result.FailedFiles = append(result.FailedFiles, path)
			return nil // Continue with next file
		}

		if exists {
			// File already exists in MinIO, skip
			return nil
		}

		// Upload file to MinIO
		file, err := os.Open(path)
		if err != nil {
			result.FailedCount++
			result.FailedFiles = append(result.FailedFiles, path)
			return nil // Continue with next file
		}
		defer file.Close()

		// Get content type
		contentType := "application/octet-stream"
		switch ext {
		case ".jpg", ".jpeg":
			contentType = "image/jpeg"
		case ".png":
			contentType = "image/png"
		case ".gif":
			contentType = "image/gif"
		}

		// Upload to MinIO
		err = client.UploadFile(ctx, objectName, file, contentType)
		if err != nil {
			result.FailedCount++
			result.FailedFiles = append(result.FailedFiles, path)
			return nil // Continue with next file
		}

		// Success
		result.SuccessCount++
		result.MigratedFiles = append(result.MigratedFiles, objectName)

		return nil
	})

	if err != nil {
		return result, err
	}

	return result, nil
}

// MigrateFile migrates a single file from local storage to MinIO
func MigrateFile(ctx context.Context, localPath string) (string, error) {
	cfg := config.GetConfig()
	client := GetClient()

	// Check if file exists
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", localPath)
	}

	// Get file info
	file, err := os.Open(localPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Determine object name
	var objectName string
	if strings.HasPrefix(localPath, cfg.Upload.SavePath) {
		// If file is in upload directory, preserve relative path
		relPath, err := filepath.Rel(cfg.Upload.SavePath, localPath)
		if err != nil {
			return "", err
		}
		objectName = fmt.Sprintf("uploads/%s", relPath)
	} else {
		// Otherwise use base filename
		objectName = fmt.Sprintf("uploads/%s", filepath.Base(localPath))
	}

	// Get content type
	ext := strings.ToLower(filepath.Ext(localPath))
	contentType := "application/octet-stream"
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	}

	// Upload to MinIO
	err = client.UploadFile(ctx, objectName, file, contentType)
	if err != nil {
		return "", err
	}

	return client.GetFileURL(objectName), nil
}
