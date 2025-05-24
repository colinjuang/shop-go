package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"github.com/colinjuang/shop-go/internal/app/middleware"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterOrderApi registers all order related api
func RegisterOrderApi(router *gin.Engine, db *gorm.DB) {
	orderHandler := handler.NewOrderHandler(db)
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 获取订单详情
		api.GET("/order/detail", orderHandler.GetOrderDetail)
		// 立即购买
		api.POST("/order/buy", orderHandler.CreateOrderAndPay)
		// 提交订单
		api.POST("/order/submit", orderHandler.CreateOrder)
		// 获取微信支付信息
		api.GET("/order/pay", orderHandler.GetWechatPayInfo)
		// 检查微信支付状态
		api.GET("/order/pay/status", orderHandler.CheckWechatPayStatus)
		// 获取订单列表
		api.GET("/order/list", orderHandler.GetOrderList)
	}
}
