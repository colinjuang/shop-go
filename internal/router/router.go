package router

import (
	"time"

	"github.com/colinjuang/shop-go/internal/router/api"

	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRouter sets up all the routes for the application
func RegisterRouter(router *gin.Engine) {
	// Set up CORS
	router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// API group
	apiV1 := router.Group("/api")

	// 轮播图
	api.RegisterBannerRouter(apiV1)
	// 分类
	api.RegisterCategoryRouter(apiV1)
	// 商品
	api.RegisterProductRouter(apiV1)
	// 促销广告
	api.RegisterPromotionRouter(apiV1)
	// 上传
	api.RegisterUploadRouter(apiV1)
	// 用户
	api.RegisterUserRouter(apiV1)
	// 地址
	api.RegisterAddressRouter(apiV1)
	// 购物车
	api.RegisterCartRouter(apiV1)
	// 订单
	api.RegisterOrderRouter(apiV1)
}
