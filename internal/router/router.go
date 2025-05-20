package router

import (
	"time"

	apiv1 "github.com/colinjuang/shop-go/internal/api/v1"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterRouter sets up all the routes for the application
func RegisterRouter(router *gin.Engine) {
	// Set up CORS
	router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// 登录
	apiv1.RegisterLoginApi(router)
	// 轮播图
	apiv1.RegisterBannerApi(router)
	// 分类
	apiv1.RegisterCategoryApi(router)
	// 商品
	apiv1.RegisterProductApi(router)
	// 促销广告
	apiv1.RegisterPromotionApi(router)
	// 上传
	apiv1.RegisterUploadApi(router)
	// 微信登录
	apiv1.RegisterWechatLoginApi(router)
	// 用户
	apiv1.RegisterUserApi(router)
	// 地址
	apiv1.RegisterAddressApi(router)
	// 购物车
	apiv1.RegisterCartApi(router)
	// 订单
	apiv1.RegisterOrderApi(router)
}
