package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/colinjuang/shop-go/internal/app/request"
	"github.com/colinjuang/shop-go/internal/app/response"
	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CartHandler handles cart-related API endpoints
type CartHandler struct {
	cartService *service.CartService
}

// NewCartHandler creates a new cart handler
func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{
		cartService: service.NewCartService(db),
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

	err := h.cartService.AddToCart(reqUser.UserID, request.ProductID, request.Quantity)
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

	cart, err := h.cartService.GetCart(reqUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(cart))
}

// UpdateCartStatus updates the status of a cart item
func (h *CartHandler) UpdateCartStatus(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	productIdStr := c.Param("productId")
	productId, err := strconv.ParseUint(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid product ID"))
		return
	}

	selectedStr := c.Param("selected")
	selected, err := strconv.ParseBool(selectedStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid selected"))
		return
	}

	err = h.cartService.UpdateCartStatus(reqUser.UserID, productId, selected)
	if err != nil {
		if err == pkgerrors.ErrCartNotFound {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// UpdateAllCartStatus updates the status of all cart items for a user
func (h *CartHandler) UpdateAllCartStatus(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	selectedStr := c.Param("selected")
	selected, err := strconv.ParseBool(selectedStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid selected"))
		return
	}

	err = h.cartService.UpdateAllCartStatus(reqUser.UserID, selected)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// DeleteCart deletes a cart item
func (h *CartHandler) DeleteCart(c *gin.Context) {
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

	err = h.cartService.DeleteCart(reqUser.UserID, id)
	if err != nil {
		if err == pkgerrors.ErrCartNotFound {
			c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}
