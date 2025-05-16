package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/dto"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related API endpoints
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(),
	}
}

// GetUserInfo gets the current user's information
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userInfo, err := h.userService.GetUserByID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(userInfo))
}

// UpdateUserInfo updates the current user's information
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	var updateInfo dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.userService.UpdateUser(c, updateInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
