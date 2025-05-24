package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterBannerApi registers all banner api
func RegisterBannerApi(router *gin.Engine, db *gorm.DB) {
	bannerHandler := handler.NewBannerHandler(db)
	api := router.Group("/api")
	{
		// 获取轮播图
		api.GET("/banner", bannerHandler.GetBanners)
	}
}
