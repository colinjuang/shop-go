package api

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterUserApi registers all user api
func RegisterUserApi(api *gin.RouterGroup) {
	userHandler := handler.NewUserHandler()
	auth := api.Use(middleware.AuthMiddleware())
	// 获取用户信息
	auth.GET("/users/info", userHandler.GetUserInfo)
	// 更新用户信息
	auth.POST("/users/update", userHandler.UpdateUserInfo)
}
