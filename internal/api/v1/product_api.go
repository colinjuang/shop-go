package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterProductApi registers all product api
func RegisterProductApi(api *gin.RouterGroup) {
	productHandler := handler.NewProductHandler()
	auth := api.Use(middleware.AuthMiddleware())
	// 获取商品列表
	auth.GET("/product", productHandler.GetProducts)
	// 获取商品详情
	auth.GET("/product/:id", productHandler.GetProductDetail)
	// 获取推荐商品
	auth.GET("/product/recommend", productHandler.GetRecommendProducts)
	// 获取热门商品
	auth.GET("/product/hot", productHandler.GetHotProducts)
}
