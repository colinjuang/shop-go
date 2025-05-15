package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterBannerRouter registers all banner routes
func RegisterBannerRouter(api *gin.RouterGroup) {
	bannerHandler := handler.NewBannerHandler()
	// 获取轮播图
	api.GET("/banners", bannerHandler.GetBanners)
}
