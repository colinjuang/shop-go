package minio

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"shop-go/internal/config"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Client represents a MinIO client wrapper
type Client struct {
	client     *minio.Client
	bucket     string
	location   string
	publicBase string
}

var minioClient *Client

// InitClient initializes the MinIO client
func InitClient(cfg *config.MinIOConfig) (*Client, error) {
	// Initialize MinIO client
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize MinIO client: %w", err)
	}

	// Create bucket if it doesn't exist
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check if bucket exists: %w", err)
	}

	if !exists {
		err = client.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{
			Region: cfg.Location,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %w", err)
		}

		// Set bucket policy to public read
		policy := `{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Principal": {"AWS": ["*"]},
					"Action": ["s3:GetObject"],
					"Resource": ["arn:aws:s3:::%s/*"]
				}
			]
		}`
		policy = fmt.Sprintf(policy, cfg.Bucket)

		err = client.SetBucketPolicy(ctx, cfg.Bucket, policy)
		if err != nil {
			return nil, fmt.Errorf("failed to set bucket policy: %w", err)
		}
	}

	// Determine public base URL
	var publicBase string
	if cfg.UseSSL {
		publicBase = fmt.Sprintf("https://%s/%s", cfg.Endpoint, cfg.Bucket)
	} else {
		publicBase = fmt.Sprintf("http://%s/%s", cfg.Endpoint, cfg.Bucket)
	}

	minioClient = &Client{
		client:     client,
		bucket:     cfg.Bucket,
		location:   cfg.Location,
		publicBase: publicBase,
	}

	return minioClient, nil
}

// GetClient returns the MinIO client instance
func GetClient() *Client {
	if minioClient == nil {
		cfg := config.GetConfig()
		var err error
		minioClient, err = InitClient(&cfg.MinIO)
		if err != nil {
			panic(err)
		}
	}
	return minioClient
}

// UploadFile uploads a file to MinIO
func (c *Client) UploadFile(ctx context.Context, objectName string, reader io.Reader, contentType string) error {
	_, err := c.client.PutObject(ctx, c.bucket, objectName, reader, -1, minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

// GetFileURL returns the URL of a file
func (c *Client) GetFileURL(objectName string) string {
	return fmt.Sprintf("%s/%s", c.publicBase, objectName)
}

// GetPresignedURL returns a presigned URL for an object
func (c *Client) GetPresignedURL(ctx context.Context, objectName string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	presignedURL, err := c.client.PresignedGetObject(ctx, c.bucket, objectName, expiry, reqParams)
	if err != nil {
		return "", err
	}
	return presignedURL.String(), nil
}

// DeleteFile deletes a file from MinIO
func (c *Client) DeleteFile(ctx context.Context, objectName string) error {
	return c.client.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
}

// ListFiles lists files in a directory (prefix)
func (c *Client) ListFiles(ctx context.Context, prefix string) ([]string, error) {
	var files []string

	objectCh := c.client.ListObjects(ctx, c.bucket, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		files = append(files, object.Key)
	}

	return files, nil
}

// FileExists checks if a file exists
func (c *Client) FileExists(ctx context.Context, objectName string) (bool, error) {
	_, err := c.client.StatObject(ctx, c.bucket, objectName, minio.StatObjectOptions{})
	if err != nil {
		// Check if the error is because the object doesn't exist
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
