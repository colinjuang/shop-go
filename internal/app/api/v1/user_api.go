package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserApi registers all user api
func RegisterUserApi(router *gin.Engine, db *gorm.DB) {
	userHandler := handler.NewUserHandler(db)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 获取用户信息
		api.GET("/user/info", userHandler.GetUserInfo)
		// 更新用户信息
		api.PUT("/user/info", userHandler.UpdateUserInfo)
	}
}
