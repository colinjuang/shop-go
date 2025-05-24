package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterBannerApi registers all banner api
func RegisterBannerApi(router *gin.Engine) {
	bannerHandler := handler.NewBannerHandler()
	api := router.Group("/api")
	{
		// 获取轮播图
		api.GET("/banner", bannerHandler.GetBanners)
	}
}
