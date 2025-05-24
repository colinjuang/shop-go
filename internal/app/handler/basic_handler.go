package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BasicHandler struct {
	db *gorm.DB
}

func NewBasicHandler(db *gorm.DB) *BasicHandler {
	return &BasicHandler{
		db: db,
	}
}

func (h *BasicHandler) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func (h *BasicHandler) DBHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := database.HealthCheck(h.db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database connection failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Database connection successful"})
	}
}

func (h *BasicHandler) DBStats() gin.HandlerFunc {
	return func(c *gin.Context) {
		stats := database.Stats(h.db)
		c.JSON(http.StatusOK, stats)
	}
}
