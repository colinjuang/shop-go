package api

import (
	"github.com/colinjuang/shop-go/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterCartRouter registers all cart related routes
func RegisterCartRouter(api *gin.RouterGroup) {
	cartHandler := handler.NewCartHandler()

	// 添加到购物车
	api.GET("/cart/add", cartHandler.AddToCart)
	// 获取购物车列表
	api.GET("/cart/list", cartHandler.GetCartList)
	// 更新购物车商品状态
	api.GET("/cart/update", cartHandler.UpdateCartItemStatus)
	// 更新购物车所有商品状态
	api.GET("/cart/update-all", cartHandler.UpdateAllCartItemStatus)
	// 删除购物车商品
	api.GET("/cart/delete", cartHandler.DeleteCartItem)
}
