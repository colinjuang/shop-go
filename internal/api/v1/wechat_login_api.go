package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterWechatLoginApi registers all wechat login api
func RegisterWechatLoginApi(api *gin.RouterGroup) {
	wechatLoginHandler := handler.NewWechatLoginHandler()

	// 微信登录
	api.GET("/login/wechat/:code", wechatLoginHandler.WechatMiniLogin)
}
