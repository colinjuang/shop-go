package router

import (
	"time"

	apiv1 "github.com/colinjuang/shop-go/internal/app/api/v1"
	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/colinjuang/shop-go/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	return router
}

// RegisterRouter sets up all the routes for the application
func RegisterRouter(router *gin.Engine, db *gorm.DB) {
	// Set up CORS
	router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// 登录
	apiv1.RegisterLoginApi(router, db)
	// 轮播图
	apiv1.RegisterBannerApi(router, db)
	// 分类
	apiv1.RegisterCategoryApi(router, db)
	// 商品
	apiv1.RegisterProductApi(router, db)
	// 促销广告
	apiv1.RegisterPromotionApi(router, db)
	// 上传
	apiv1.RegisterUploadApi(router, db)
	// 微信登录
	apiv1.RegisterWechatLoginApi(router, db)
	// 用户
	apiv1.RegisterUserApi(router, db)
	// 地址
	apiv1.RegisterAddressApi(router, db)
	// 购物车
	apiv1.RegisterCartApi(router, db)
	// 订单
	apiv1.RegisterOrderApi(router, db)
}
