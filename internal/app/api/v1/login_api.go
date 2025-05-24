package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterLoginApi(router *gin.Engine, db *gorm.DB) {
	loginHandler := handler.NewLoginHandler(db)
	api := router.Group("/api")
	{
		// 登录
		api.POST("/login", loginHandler.Login)
		// 注册
		api.POST("/register", loginHandler.Register)
	}
}
