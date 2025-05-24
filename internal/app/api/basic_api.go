package api

import (
	"github.com/colinjuang/shop-go/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBasicApi(router *gin.Engine) {
	basicHandler := handler.NewBasicHandler()
	api := router.Group("/api")
	{
		// 健康状态检测
		api.HEAD("/health", basicHandler.Health())
		// 数据库健康检查
		api.GET("/db/health", basicHandler.DBHealth())
		// 数据库统计
		api.GET("/db/stats", basicHandler.DBStats())
	}
}
