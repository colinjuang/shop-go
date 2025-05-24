package router

import (
	"time"

	"github.com/colinjuang/shop-go/internal/app/api"
	apiv1 "github.com/colinjuang/shop-go/internal/app/api/v1"
	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	// 设置 gin 模式
	switch cfg.Server.Environment {
	case "production":
		gin.SetMode(gin.ReleaseMode)
	case "development":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建一个带有我们日志记录器的自定义 gin.Engine
	router := gin.New()

	// 使用我们自己的日志记录器和恢复中间件
	router.Use(middleware.ZapLogger())
	router.Use(gin.Recovery())

	// 设置 CORS
	router.Use(middleware.CORSMiddleware())

	// 设置全局限流 - 每分钟100个请求
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	return router
}

// RegisterRouter sets up all the routes for the application
func RegisterRouter(router *gin.Engine) {
	// 基础 API
	api.RegisterBasicApi(router)

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
