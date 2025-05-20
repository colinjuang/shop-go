package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterUserApi registers all user api
func RegisterUserApi(router *gin.Engine) {
	userHandler := handler.NewUserHandler()

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 获取用户信息
		api.GET("/user/info", userHandler.GetUserInfo)
		// 更新用户信息
		api.PUT("/user/info", userHandler.UpdateUserInfo)
	}
}
