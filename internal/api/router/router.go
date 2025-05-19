package router

import (
	"time"

	"github.com/colinjuang/shop-go/internal/api/middleware"
	apiv1 "github.com/colinjuang/shop-go/internal/api/v1"
	"github.com/gin-gonic/gin"
)

// RegisterRouter sets up all the routes for the application
func RegisterRouter(router *gin.Engine) {
	// Set up CORS
	router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// API group
	api := router.Group("/api")

	// 轮播图
	apiv1.RegisterBannerApi(api)
	// 分类
	apiv1.RegisterCategoryApi(api)
	// 商品
	apiv1.RegisterProductApi(api)
	// 促销广告
	apiv1.RegisterPromotionApi(api)
	// 上传
	apiv1.RegisterUploadApi(api)
	// 登录
	apiv1.RegisterLoginApi(api)
	// 微信登录
	apiv1.RegisterWechatLoginApi(api)
	// 用户
	apiv1.RegisterUserApi(api)
	// 地址
	apiv1.RegisterAddressApi(api)
	// 购物车
	apiv1.RegisterCartApi(api)
	// 订单
	apiv1.RegisterOrderApi(api)
}
