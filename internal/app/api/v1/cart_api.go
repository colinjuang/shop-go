package v1

import (
	"github.com/colinjuang/shop-go/internal/app/handler"
	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterCartApi registers all cart related api
func RegisterCartApi(router *gin.Engine) {
	cartHandler := handler.NewCartHandler()
	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 添加到购物车
		api.POST("/cart", cartHandler.AddToCart)
		// 获取购物车列表
		api.GET("/cart", cartHandler.GetCartList)
		// 更新购物车商品状态
		api.PUT("/cart/:productId/:selected", cartHandler.UpdateCartStatus)
		// 更新购物车所有商品状态
		api.PUT("/cart/all/:selected", cartHandler.UpdateAllCartStatus)
		// 删除购物车商品
		api.DELETE("/cart", cartHandler.DeleteCart)
	}
}
