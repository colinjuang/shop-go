package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUserApi registers all user and address related api
func RegisterUserApi(api *gin.RouterGroup) {
	userHandler := handler.NewUserHandler()

	// 获取用户信息
	api.GET("/users/info", userHandler.GetUserInfo)
	// 更新用户信息
	api.POST("/users/update", userHandler.UpdateUserInfo)
}
