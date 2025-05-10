package handler

import (
	"net/http"
	"shop-go/internal/model"
	"shop-go/internal/service"
	"strconv"

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
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var req model.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	address, err := h.addressService.CreateAddress(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(address))
}

// GetAddressList gets all addresses for a user
func (h *AddressHandler) GetAddressList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	addresses, err := h.addressService.GetAddressesByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(addresses))
}

// GetAddressDetail gets an address by ID
func (h *AddressHandler) GetAddressDetail(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	address, err := h.addressService.GetAddressByID(uint(id), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(address))
}

// UpdateAddress updates an address
func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	var req model.AddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err = h.addressService.UpdateAddress(uint(id), userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}

// DeleteAddress deletes an address
func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	err = h.addressService.DeleteAddress(uint(id), userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
