package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAddressApi registers all user and address related api
func RegisterAddressApi(api *gin.RouterGroup) {
	addressHandler := handler.NewAddressHandler()
	auth := api.Use(middleware.AuthMiddleware())

	// 添加地址
	auth.POST("/address", addressHandler.CreateAddress)
	// 获取地址列表
	auth.GET("/address", addressHandler.GetAddressList)
	// 获取地址详情
	auth.GET("/address/:id", addressHandler.GetAddressDetail)
	// 更新地址
	auth.PUT("/address/:id", addressHandler.UpdateAddress)
	// 删除地址
	auth.DELETE("/address/:id", addressHandler.DeleteAddress)
}
