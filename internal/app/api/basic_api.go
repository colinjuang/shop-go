package api

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterBasicApi(router *gin.Engine, db *gorm.DB) {
	basicHandler := handler.NewBasicHandler(db)
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
