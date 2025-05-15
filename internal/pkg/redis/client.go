package redis

import (
	"context"
	"fmt"
	"github.com/colinjuang/shop-go/internal/config"
	"time"

	"github.com/redis/go-redis/v9"
)

// Client represents a Redis client wrapper
type Client struct {
	client *redis.Client
	prefix string
}

var redisClient *Client

// InitClient initializes the Redis client
func InitClient(cfg *config.RedisConfig) (*Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	redisClient = &Client{
		client: client,
		prefix: cfg.Prefix,
	}

	return redisClient, nil
}

// GetClient returns the Redis client instance
func GetClient() *Client {
	if redisClient == nil {
		cfg := config.GetConfig()
		var err error
		redisClient, err = InitClient(&cfg.Redis)
		if err != nil {
			panic(err)
		}
	}
	return redisClient
}

// prefixKey adds the configured prefix to a key
func (c *Client) prefixKey(key string) string {
	return c.prefix + key
}

// Set sets a key-value pair with expiration
func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, c.prefixKey(key), value, expiration).Err()
}

// Get gets a value by key
func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, c.prefixKey(key)).Result()
}

// Delete deletes a key
func (c *Client) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, c.prefixKey(key)).Err()
}

// HashSet sets a hash field
func (c *Client) HashSet(ctx context.Context, key, field string, value interface{}) error {
	return c.client.HSet(ctx, c.prefixKey(key), field, value).Err()
}

// HashGet gets a hash field
func (c *Client) HashGet(ctx context.Context, key, field string) (string, error) {
	return c.client.HGet(ctx, c.prefixKey(key), field).Result()
}

// HashGetAll gets all fields in a hash
func (c *Client) HashGetAll(ctx context.Context, key string) (map[string]string, error) {
	return c.client.HGetAll(ctx, c.prefixKey(key)).Result()
}

// Exists checks if a key exists
func (c *Client) Exists(ctx context.Context, key string) (bool, error) {
	result, err := c.client.Exists(ctx, c.prefixKey(key)).Result()
	return result > 0, err
}

// SetNX sets a key-value pair if the key does not exist
func (c *Client) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.client.SetNX(ctx, c.prefixKey(key), value, expiration).Result()
}

// Close closes the Redis client
func (c *Client) Close() error {
	return c.client.Close()
}
