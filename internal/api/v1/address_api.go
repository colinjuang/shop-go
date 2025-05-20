package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAddressApi registers all user and address related api
func RegisterAddressApi(router *gin.Engine) {
	addressHandler := handler.NewAddressHandler()

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 添加地址
		api.POST("/address", addressHandler.CreateAddress)
		// 获取地址列表
		api.GET("/address", addressHandler.GetAddressList)
		// 获取地址详情
		api.GET("/address/:id", addressHandler.GetAddressDetail)
		// 更新地址
		api.PUT("/address/:id", addressHandler.UpdateAddress)
		// 删除地址
		api.DELETE("/address/:id", addressHandler.DeleteAddress)
	}
}
