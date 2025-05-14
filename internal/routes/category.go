package routes

import (
	"shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterCategoryRoutes registers all category routes
func RegisterCategoryRoutes(api *gin.RouterGroup, categoryHandler *handler.CategoryHandler) {

	// 获取所有分类
	api.GET("/categories", categoryHandler.GetAllCategories)
	// 获取子分类
	api.GET("/categories/:id/subs", categoryHandler.GetSubCategories)
	// 获取分类树
	api.GET("/categories/tree", categoryHandler.GetCategoryTree)
}
