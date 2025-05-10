package middleware

import (
	"net/http"
	"shop-go/internal/model"
	"shop-go/internal/pkg/redis"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware creates a middleware for rate limiting using Redis
// limit: maximum number of requests per window
// window: time window in seconds
func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	redisClient := redis.GetClient()

	return func(c *gin.Context) {
		// Get client IP as identifier
		clientIP := c.ClientIP()
		key := "rate_limit:" + clientIP

		ctx := c.Request.Context()

		// Check if key exists in Redis
		exists, err := redisClient.Exists(ctx, key)
		if err != nil {
			// If Redis fails, don't block the request
			c.Next()
			return
		}

		if !exists {
			// First request from this IP, initialize counter
			err = redisClient.Set(ctx, key, "1", window)
			if err != nil {
				// If Redis fails, don't block the request
				c.Next()
				return
			}
			c.Next()
			return
		}

		// Get current count
		countStr, err := redisClient.Get(ctx, key)
		if err != nil {
			// If Redis fails, don't block the request
			c.Next()
			return
		}

		count, err := strconv.Atoi(countStr)
		if err != nil {
			// If parsing fails, don't block the request
			c.Next()
			return
		}

		// Check if limit exceeded
		if count >= limit {
			c.JSON(http.StatusTooManyRequests, model.ErrorResponse(
				http.StatusTooManyRequests,
				"Rate limit exceeded. Please try again later.",
			))
			c.Abort()
			return
		}

		// Increment counter
		// We could use INCR but we're using our abstraction for consistency
		err = redisClient.Set(ctx, key, strconv.Itoa(count+1), window)
		if err != nil {
			// If Redis fails, don't block the request
			c.Next()
			return
		}

		// Set rate limit headers
		c.Header("X-RateLimit-Limit", strconv.Itoa(limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(limit-count-1))

		c.Next()
	}
}
