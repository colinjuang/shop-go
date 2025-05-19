package v1

import (
	"github.com/colinjuang/shop-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterLoginApi(api *gin.RouterGroup) {
	loginHandler := handler.NewLoginHandler()
	// 登录
	api.POST("/login", loginHandler.Login)
	// 注册
	api.POST("/register", loginHandler.Register)
}
