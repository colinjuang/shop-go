package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/api/request"
	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
)

// AddressHandler handles address-related API endpoints
type AddressHandler struct {
	addressService *service.AddressService
}

// NewAddressHandler creates a new address handler
func NewAddressHandler() *AddressHandler {
	return &AddressHandler{
		addressService: service.NewAddressService(),
	}
}

// CreateAddress creates a new address
func (h *AddressHandler) CreateAddress(c *gin.Context) {
	var req request.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	address, err := h.addressService.CreateAddress(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(address))
}

// GetAddressList gets all addresses for a user
func (h *AddressHandler) GetAddressList(c *gin.Context) {
	addresses, err := h.addressService.GetAddressesByUserID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(addresses))
}

// GetAddressDetail gets an address by ID
func (h *AddressHandler) GetAddressDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	address, err := h.addressService.GetAddressByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(address))
}

// UpdateAddress updates an address
func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	var req request.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	if req.ID == 0 {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	err := h.addressService.UpdateAddress(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// DeleteAddress deletes an address
func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	err = h.addressService.DeleteAddress(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}
