package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterPromotionApi registers all promotion api
func RegisterPromotionApi(api *gin.RouterGroup) {
	promotionHandler := handler.NewPromotionHandler()
	// 获取促销广告
	api.GET("/promotion", promotionHandler.GetPromotions)
}
