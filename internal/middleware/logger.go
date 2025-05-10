package middleware

import (
	"shop-go/internal/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ZapLogger returns a gin middleware for logging HTTP requests using zap
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log details after request is processed
		end := time.Now()
		latency := end.Sub(start)

		// Get status code and client IP
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Log with appropriate level based on status code
		if statusCode >= 500 {
			logger.Error("Server error",
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", clientIP),
				zap.Duration("latency", latency),
				zap.String("error", errorMessage),
			)
		} else if statusCode >= 400 {
			logger.Warn("Client error",
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", clientIP),
				zap.Duration("latency", latency),
				zap.String("error", errorMessage),
			)
		} else {
			logger.Info("Request completed",
				zap.Int("status", statusCode),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", clientIP),
				zap.Duration("latency", latency),
			)
		}
	}
}
