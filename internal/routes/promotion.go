package routes

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPromotionRoutes registers all promotion routes
func RegisterPromotionRoutes(api *gin.RouterGroup, promotionHandler *handler.PromotionHandler) {
	// 获取促销广告
	api.GET("/promotions", promotionHandler.GetPromotions)
}
