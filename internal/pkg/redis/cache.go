package redis

import (
	"context"
	"encoding/json"
	"time"
)

// CacheService provides cache operations using Redis
type CacheService struct {
	client *Client
}

// NewCacheService creates a new cache service
func NewCacheService() *CacheService {
	return &CacheService{
		client: GetClient(),
	}
}

// Set caches a value with expiration
func (s *CacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Marshal complex objects to JSON
	if _, ok := value.(string); !ok && value != nil {
		jsonData, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return s.client.Set(ctx, key, jsonData, expiration)
	}
	return s.client.Set(ctx, key, value, expiration)
}

// Get retrieves a value from cache as string
func (s *CacheService) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key)
}

// GetObject retrieves a value from cache and unmarshals it to the provided object
func (s *CacheService) GetObject(ctx context.Context, key string, obj interface{}) error {
	data, err := s.client.Get(ctx, key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(data), obj)
}

// Delete removes a key from cache
func (s *CacheService) Delete(ctx context.Context, key string) error {
	return s.client.Delete(ctx, key)
}

// Exists checks if a key exists in cache
func (s *CacheService) Exists(ctx context.Context, key string) (bool, error) {
	return s.client.Exists(ctx, key)
}

// Obtain tries to get a value from cache; if not exists, it calls the function and caches the result
func (s *CacheService) Obtain(ctx context.Context, key string, expiration time.Duration, fn func() (interface{}, error)) (string, error) {
	// Try to get from cache
	exists, err := s.Exists(ctx, key)
	if err != nil {
		return "", err
	}

	if exists {
		return s.Get(ctx, key)
	}

	// Call the function to get the value
	result, err := fn()
	if err != nil {
		return "", err
	}

	// Cache the result
	err = s.Set(ctx, key, result, expiration)
	if err != nil {
		return "", err
	}

	// Convert result to string if needed
	if str, ok := result.(string); ok {
		return str, nil
	}

	// For complex objects, we need to marshal them back
	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// ObtainObject is similar to Obtain but for complex objects
func (s *CacheService) ObtainObject(ctx context.Context, key string, expiration time.Duration, obj interface{}, fn func() (interface{}, error)) error {
	exists, err := s.Exists(ctx, key)
	if err != nil {
		return err
	}

	if exists {
		return s.GetObject(ctx, key, obj)
	}

	result, err := fn()
	if err != nil {
		return err
	}

	err = s.Set(ctx, key, result, expiration)
	if err != nil {
		return err
	}

	// If result is already the target object, return it
	if resultObj, ok := result.(interface{}); ok && resultObj == obj {
		return nil
	}

	// Otherwise, get the object from cache
	return s.GetObject(ctx, key, obj)
}
