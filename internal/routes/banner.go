package routes

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterBannerRoutes registers all banner routes
func RegisterBannerRoutes(api *gin.RouterGroup, bannerHandler *handler.BannerHandler) {
	// 获取轮播图
	api.GET("/banners", bannerHandler.GetBanners)
}
