package v1

import (
	"github.com/colinjuang/shop-go/internal/api/handler"

	"github.com/gin-gonic/gin"
)

// RegisterCategoryApi registers all category api
func RegisterCategoryApi(api *gin.RouterGroup) {
	categoryHandler := handler.NewCategoryHandler()

	// 获取所有分类
	api.GET("/category", categoryHandler.GetAllCategories)
	// 获取子分类
	api.GET("/category/:id/subs", categoryHandler.GetSubCategories)
	// 获取分类树
	api.GET("/category/tree", categoryHandler.GetCategoryTree)
	// 获取一级分类
	api.GET("/category/level1", categoryHandler.GetLevel1Categories)
}
