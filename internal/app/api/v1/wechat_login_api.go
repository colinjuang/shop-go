package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterWechatLoginApi registers all wechat login api
func RegisterWechatLoginApi(router *gin.Engine) {
	wechatLoginHandler := handler.NewWechatLoginHandler()
	api := router.Group("/api")
	{
		// 微信登录
		api.GET("/login/wechat/:code", wechatLoginHandler.WechatMiniLogin)
	}
}
