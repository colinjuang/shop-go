package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-go/internal/model"
	"shop-go/internal/service"
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
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(promotions))
}
