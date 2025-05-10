package minio

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// CacheOptions represents options for caching files in MinIO
type CacheOptions struct {
	// Prefix is the prefix for the object name in MinIO
	Prefix string
	// TTL is the time-to-live for the cache
	TTL time.Duration
	// ContentType is the content type of the cached file
	ContentType string
}

// DefaultCacheOptions returns default cache options
func DefaultCacheOptions() CacheOptions {
	return CacheOptions{
		Prefix:      "cache",
		TTL:         24 * time.Hour,
		ContentType: "application/octet-stream",
	}
}

// CacheFile caches a file in MinIO and returns its URL
// If the file is already cached, it returns the cached URL
func CacheFile(ctx context.Context, reader io.Reader, key string, opts CacheOptions) (string, error) {
	client := GetClient()

	// Generate object name from key
	hash := md5.Sum([]byte(key))
	objectName := fmt.Sprintf("%s/%s", opts.Prefix, hex.EncodeToString(hash[:]))

	// Check if file exists in cache
	exists, err := client.FileExists(ctx, objectName)
	if err == nil && exists {
		// File exists in cache, return URL
		return client.GetFileURL(objectName), nil
	}

	// File doesn't exist or error occurred, cache the file
	err = client.UploadFile(ctx, objectName, reader, opts.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to cache file: %w", err)
	}

	// Return file URL
	return client.GetFileURL(objectName), nil
}

// CacheFileFromPath caches a file from local path in MinIO and returns its URL
func CacheFileFromPath(ctx context.Context, path string, key string, opts CacheOptions) (string, error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// If content type not specified, try to detect it from file extension
	if opts.ContentType == "application/octet-stream" {
		ext := strings.ToLower(filepath.Ext(path))
		switch ext {
		case ".jpg", ".jpeg":
			opts.ContentType = "image/jpeg"
		case ".png":
			opts.ContentType = "image/png"
		case ".gif":
			opts.ContentType = "image/gif"
		case ".pdf":
			opts.ContentType = "application/pdf"
		case ".txt":
			opts.ContentType = "text/plain"
		case ".html", ".htm":
			opts.ContentType = "text/html"
		case ".json":
			opts.ContentType = "application/json"
		case ".xml":
			opts.ContentType = "application/xml"
		case ".zip":
			opts.ContentType = "application/zip"
		}
	}

	// Cache file
	return CacheFile(ctx, file, key, opts)
}

// GetCachedFile retrieves a cached file from MinIO
// If the file is not in cache, it calls the generator function to create it
func GetCachedFile(ctx context.Context, key string, opts CacheOptions, generator func() (io.Reader, error)) (string, error) {
	client := GetClient()

	// Generate object name from key
	hash := md5.Sum([]byte(key))
	objectName := fmt.Sprintf("%s/%s", opts.Prefix, hex.EncodeToString(hash[:]))

	// Check if file exists in cache
	exists, err := client.FileExists(ctx, objectName)
	if err == nil && exists {
		// File exists in cache, return URL
		return client.GetFileURL(objectName), nil
	}

	// File doesn't exist or error occurred, generate the file
	reader, err := generator()
	if err != nil {
		return "", fmt.Errorf("failed to generate file: %w", err)
	}

	// Cache the generated file
	err = client.UploadFile(ctx, objectName, reader, opts.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to cache file: %w", err)
	}

	// Return file URL
	return client.GetFileURL(objectName), nil
}

// GetCachedFileWithTempFile is similar to GetCachedFile but creates a temporary file
// This is useful when the generator function can't directly produce an io.Reader
func GetCachedFileWithTempFile(ctx context.Context, key string, opts CacheOptions, generator func(tempPath string) error) (string, error) {
	client := GetClient()

	// Generate object name from key
	hash := md5.Sum([]byte(key))
	objectName := fmt.Sprintf("%s/%s", opts.Prefix, hex.EncodeToString(hash[:]))

	// Check if file exists in cache
	exists, err := client.FileExists(ctx, objectName)
	if err == nil && exists {
		// File exists in cache, return URL
		return client.GetFileURL(objectName), nil
	}

	// Create temporary file
	tempFile, err := ioutil.TempFile("", "minio-cache-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	tempPath := tempFile.Name()
	tempFile.Close()
	defer os.Remove(tempPath)

	// Generate the file
	err = generator(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to generate file: %w", err)
	}

	// Open the generated file
	file, err := os.Open(tempPath)
	if err != nil {
		return "", fmt.Errorf("failed to open generated file: %w", err)
	}
	defer file.Close()

	// Cache the generated file
	err = client.UploadFile(ctx, objectName, file, opts.ContentType)
	if err != nil {
		return "", fmt.Errorf("failed to cache file: %w", err)
	}

	// Return file URL
	return client.GetFileURL(objectName), nil
}
