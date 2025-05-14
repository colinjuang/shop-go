package routes

import (
	"shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user and address related routes
func RegisterAddressRoutes(api *gin.RouterGroup, addressHandler *handler.AddressHandler) {
	// 添加地址
	api.POST("/address/add", addressHandler.CreateAddress)
	// 获取地址列表
	api.GET("/address/list", addressHandler.GetAddressList)
	// 获取地址详情
	api.GET("/address/:id", addressHandler.GetAddressDetail)
	// 更新地址
	api.POST("/address/:id/update", addressHandler.UpdateAddress)
	// 删除地址
	api.GET("/address/:id/delete", addressHandler.DeleteAddress)
}
