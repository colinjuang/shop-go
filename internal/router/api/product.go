package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterProductApi registers all product api
func RegisterProductApi(api *gin.RouterGroup) {
	productHandler := handler.NewProductHandler()

	// 获取商品列表
	api.GET("/products", productHandler.GetProducts)
	// 获取商品详情
	api.GET("/products/:id", productHandler.GetProductDetail)
	// 获取推荐商品
	api.GET("/products/recommend", productHandler.GetRecommendProducts)
	// 获取热门商品
	api.GET("/products/hot", productHandler.GetHotProducts)
}
