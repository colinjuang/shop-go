package routes

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all the routes for the application
func RegisterRoutes(router *gin.Engine) {
	// Set up CORS
	router.Use(middleware.CORSMiddleware())

	// Set up global rate limiting - 100 requests per minute
	router.Use(middleware.RateLimitMiddleware(100, 1*time.Minute))

	// API group
	api := router.Group("/api")

	// 轮播图
	RegisterBannerRoutes(api, handler.NewBannerHandler())
	// 分类
	RegisterCategoryRoutes(api, handler.NewCategoryHandler())
	// 商品
	RegisterProductRoutes(api, handler.NewProductHandler())
	// 促销广告
	RegisterPromotionRoutes(api, handler.NewPromotionHandler())

	// Protected endpoints (auth required)
	// auth := api.Group("/")
	// auth.Use(middleware.AuthMiddleware())

	// 上传
	RegisterUploadRoutes(api, handler.NewUploadHandler())
	// 用户
	RegisterUserRoutes(api, handler.NewUserHandler())
	// 地址
	RegisterAddressRoutes(api, handler.NewAddressHandler())
	// 购物车
	RegisterCartRoutes(api, handler.NewCartHandler())
	// 订单
	RegisterOrderRoutes(api, handler.NewOrderHandler())

	// Static files for uploads
	router.Static("/uploads", "./uploads")
}
