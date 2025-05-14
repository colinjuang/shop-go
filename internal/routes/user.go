package routes

import (
	"shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all user and address related routes
func RegisterUserRoutes(api *gin.RouterGroup, userHandler *handler.UserHandler) {
	// 获取用户信息
	api.GET("/users/info", userHandler.GetUserInfo)
	// 更新用户信息
	api.POST("/users/update", userHandler.UpdateUserInfo)
}
