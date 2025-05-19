package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterOrderApi registers all order related api
func RegisterOrderApi(api *gin.RouterGroup) {
	orderHandler := handler.NewOrderHandler()
	auth := api.Use(middleware.AuthMiddleware())
	// 获取订单详情
	auth.GET("/order/detail", orderHandler.GetOrderDetail)
	// 获取订单地址
	auth.GET("/order/address", orderHandler.GetOrderAddress)
	// 提交订单
	auth.POST("/order/submit", orderHandler.CreateOrder)
	// 获取微信支付信息
	auth.GET("/order/pay", orderHandler.GetWechatPayInfo)
	// 检查微信支付状态
	auth.GET("/order/pay/status", orderHandler.CheckWechatPayStatus)
	// 获取订单列表
	auth.GET("/order/list", orderHandler.GetOrderList)
}
