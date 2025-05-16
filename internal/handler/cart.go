package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"

	"github.com/gin-gonic/gin"
)

// CartHandler handles cart-related API endpoints
type CartHandler struct {
	cartService *service.CartService
}

// NewCartHandler creates a new cart handler
func NewCartHandler() *CartHandler {
	return &CartHandler{
		cartService: service.NewCartService(),
	}
}

// AddToCart adds a product to the cart
func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	productIDStr := c.Query("product_id")
	quantityStr := c.DefaultQuery("quantity", "1")

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid product ID"))
		return
	}

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil || quantity < 1 {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid quantity"))
		return
	}

	err = h.cartService.AddToCart(userID.(uint), uint(productID), quantity)
	if err != nil {
		if err == service.ErrorOutOfStock {
			c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}

// GetCartList gets all cart items for a user
func (h *CartHandler) GetCartList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	cartItems, err := h.cartService.GetCartItems(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(cartItems))
}

// UpdateCartItemStatus updates the status of a cart item
func (h *CartHandler) UpdateCartItemStatus(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Query("id")
	selectedStr := c.DefaultQuery("selected", "true")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	selected := selectedStr == "true" || selectedStr == "1"

	err = h.cartService.UpdateCartItemStatus(uint(id), userID.(uint), selected)
	if err != nil {
		if err == service.ErrorCartItemNotFound {
			c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (h *CartHandler) UpdateAllCartItemStatus(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	selectedStr := c.DefaultQuery("selected", "true")
	selected := selectedStr == "true" || selectedStr == "1"

	err := h.cartService.UpdateAllCartItemStatus(userID.(uint), selected)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}

// DeleteCartItem deletes a cart item
func (h *CartHandler) DeleteCartItem(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	err = h.cartService.DeleteCartItem(uint(id), userID.(uint))
	if err != nil {
		if err == service.ErrorCartItemNotFound {
			c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
