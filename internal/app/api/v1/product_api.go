package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterProductApi registers all product api
func RegisterProductApi(router *gin.Engine, db *gorm.DB) {
	productHandler := handler.NewProductHandler(db)
	api := router.Group("/api")
	{
		// 获取商品列表
		api.GET("/product", productHandler.GetProducts)
		// 获取商品详情
		api.GET("/product/:id", productHandler.GetProductDetail)
		// 获取推荐商品
		api.GET("/product/recommend", productHandler.GetRecommendProducts)
		// 获取热门商品
		api.GET("/product/hot", productHandler.GetHotProducts)
	}
}
