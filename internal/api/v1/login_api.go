package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterLoginApi(router *gin.Engine) {
	loginHandler := handler.NewLoginHandler()
	api := router.Group("/api")
	{
		// 登录
		api.POST("/login", loginHandler.Login)
		// 注册
		api.POST("/register", loginHandler.Register)
	}
}
