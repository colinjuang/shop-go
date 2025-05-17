package router

import (
	"time"

	"github.com/colinjuang/shop-go/internal/api/middleware"
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
	RegisterBannerApi(apiV1)
	// 分类
	RegisterCategoryApi(apiV1)
	// 商品
	RegisterProductApi(apiV1)
	// 促销广告
	RegisterPromotionApi(apiV1)
	// 上传
	RegisterUploadApi(apiV1)
	// 登录
	RegisterLoginApi(apiV1)
	// 微信登录
	RegisterWechatLoginApi(apiV1)
	// 用户
	RegisterUserApi(apiV1)
	// 地址
	RegisterAddressApi(apiV1)
	// 购物车
	RegisterCartApi(apiV1)
	// 订单
	RegisterOrderApi(apiV1)
}
