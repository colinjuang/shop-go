package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPromotionRouter registers all promotion routes
func RegisterPromotionRouter(api *gin.RouterGroup) {
	promotionHandler := handler.NewPromotionHandler()
	// 获取促销广告
	api.GET("/promotions", promotionHandler.GetPromotions)
}
