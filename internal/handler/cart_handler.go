package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/middleware"
	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"
	"github.com/colinjuang/shop-go/internal/request"
	"github.com/colinjuang/shop-go/internal/response"
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
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var request request.AddToCartRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	if request.Quantity < 1 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid quantity"))
		return
	}

	err := h.cartService.AddToCart(reqUser, request.ProductID, request.Quantity)
	if err != nil {
		if err == pkgerrors.ErrOutOfStock {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// GetCartList gets all cart items for a user
func (h *CartHandler) GetCartList(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	cartItems, err := h.cartService.GetCartItems(reqUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(cartItems))
}

// UpdateCartItemStatus updates the status of a cart item
func (h *CartHandler) UpdateCartItemStatus(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Query("id")
	selectedStr := c.DefaultQuery("selected", "true")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	selected := selectedStr == "true" || selectedStr == "1"

	err = h.cartService.UpdateCartItemStatus(id, reqUser.UserID, selected)
	if err != nil {
		if err == pkgerrors.ErrCartItemNotFound {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (h *CartHandler) UpdateAllCartItemStatus(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	selectedStr := c.DefaultQuery("selected", "true")
	selected := selectedStr == "true" || selectedStr == "1"

	err := h.cartService.UpdateAllCartItemStatus(reqUser.UserID, selected)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// DeleteCartItem deletes a cart item
func (h *CartHandler) DeleteCartItem(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	err = h.cartService.DeleteCartItem(id, reqUser.UserID)
	if err != nil {
		if err == pkgerrors.ErrCartItemNotFound {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}
