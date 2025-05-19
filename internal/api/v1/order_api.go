package v1

import (
	"github.com/colinjuang/shop-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

// RegisterOrderApi registers all order related api
func RegisterOrderApi(api *gin.RouterGroup) {
	orderHandler := handler.NewOrderHandler()
	// 获取订单详情
	api.GET("/order/detail", orderHandler.GetOrderDetail)
	// 获取订单地址
	api.GET("/order/address", orderHandler.GetOrderAddress)
	// 提交订单
	api.POST("/order/submit", orderHandler.CreateOrder)
	// 获取微信支付信息
	api.GET("/order/pay", orderHandler.GetWechatPayInfo)
	// 检查微信支付状态
	api.GET("/order/pay/status", orderHandler.CheckWechatPayStatus)
	// 获取订单列表
	api.GET("/order/list", orderHandler.GetOrderList)
}
