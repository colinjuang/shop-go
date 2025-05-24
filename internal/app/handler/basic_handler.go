package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/colinjuang/shop-go/internal/server"
	"github.com/gin-gonic/gin"
)

type BasicHandler struct {
}

func NewBasicHandler() *BasicHandler {
	return &BasicHandler{}
}

func (h *BasicHandler) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (h *BasicHandler) DBHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := database.HealthCheck(server.GetServer().DB); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Database connection successful"})
	}
}

func (h *BasicHandler) DBStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		stats := database.Stats(server.GetServer().DB)
		c.JSON(http.StatusOK, stats)
	}
}
