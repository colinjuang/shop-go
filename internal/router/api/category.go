package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterCategoryRouter registers all category routes
func RegisterCategoryRouter(api *gin.RouterGroup) {
	categoryHandler := handler.NewCategoryHandler()

	// 获取所有分类
	api.GET("/categories", categoryHandler.GetAllCategories)
	// 获取子分类
	api.GET("/categories/:id/subs", categoryHandler.GetSubCategories)
	// 获取分类树
	api.GET("/categories/tree", categoryHandler.GetCategoryTree)
}
