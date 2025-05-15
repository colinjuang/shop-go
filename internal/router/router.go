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
	api.RegisterBannerApi(apiV1)
	// 分类
	api.RegisterCategoryApi(apiV1)
	// 商品
	api.RegisterProductApi(apiV1)
	// 促销广告
	api.RegisterPromotionApi(apiV1)
	// 上传
	api.RegisterUploadApi(apiV1)
	// 用户
	api.RegisterUserApi(apiV1)
	// 地址
	api.RegisterAddressApi(apiV1)
	// 购物车
	api.RegisterCartApi(apiV1)
	// 订单
	api.RegisterOrderApi(apiV1)
}
