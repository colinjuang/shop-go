package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
	userIdStr := c.Query("userId")
	fmt.Println("userId", userIdStr)
	if userIdStr == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	user, err := h.userService.GetUserByID(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(user))
}

// UpdateUserInfo updates the current user's information
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var updateInfo model.UserUpdateInfo
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.userService.UpdateUser(userId.(uint), updateInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
