package v1

import (
	"github.com/colinjuang/shop-go/internal/handler"
	"github.com/colinjuang/shop-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterCartApi registers all cart related api
func RegisterCartApi(api *gin.RouterGroup) {
	cartHandler := handler.NewCartHandler()
	auth := api.Use(middleware.AuthMiddleware())
	// 添加到购物车
	auth.POST("/cart", cartHandler.AddToCart)
	// 获取购物车列表
	auth.GET("/cart", cartHandler.GetCartList)
	// 更新购物车商品状态
	auth.PUT("/cart/:productId/:selected", cartHandler.UpdateCartStatus)
	// 更新购物车所有商品状态
	auth.PUT("/cart/all/:selected", cartHandler.UpdateAllCartStatus)
	// 删除购物车商品
	auth.DELETE("/cart", cartHandler.DeleteCart)
}
