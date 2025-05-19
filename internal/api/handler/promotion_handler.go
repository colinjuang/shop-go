package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
)

// PromotionHandler handles home page API endpoints
type PromotionHandler struct {
	promotionService *service.PromotionService
}

// NewPromotionHandler creates a new promotion handler
func NewPromotionHandler() *PromotionHandler {
	return &PromotionHandler{
		promotionService: service.NewPromotionService(),
	}
}

// GetPromotions gets all promotions
func (h *PromotionHandler) GetPromotions(c *gin.Context) {
	promotions, err := h.promotionService.GetPromotions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(promotions))
}
