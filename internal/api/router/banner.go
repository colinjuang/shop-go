package router

import (
	"github.com/colinjuang/shop-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

// RegisterBannerApi registers all banner api
func RegisterBannerApi(api *gin.RouterGroup) {
	bannerHandler := handler.NewBannerHandler()
	// 获取轮播图
	api.GET("/banner", bannerHandler.GetBanners)
}
