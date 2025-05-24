package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterCategoryApi registers all category api
func RegisterCategoryApi(router *gin.Engine, db *gorm.DB) {
	categoryHandler := handler.NewCategoryHandler(db)
	api := router.Group("/api")
	{
		// 获取所有分类
		api.GET("/category", categoryHandler.GetAllCategories)
		// 获取子分类
		api.GET("/category/:id/subs", categoryHandler.GetSubCategories)
		// 获取分类树
		api.GET("/category/tree", categoryHandler.GetCategoryTree)
		// 获取一级分类
		api.GET("/category/level1", categoryHandler.GetLevel1Categories)
	}
}
