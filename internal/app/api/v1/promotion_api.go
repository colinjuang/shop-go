package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterPromotionApi registers all promotion api
func RegisterPromotionApi(router *gin.Engine, db *gorm.DB) {
	promotionHandler := handler.NewPromotionHandler(db)
	api := router.Group("/api")
	{
		// 获取促销广告
		api.GET("/promotion", promotionHandler.GetPromotions)
	}
}
