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
	userIDStr := c.Query("userID")
	fmt.Println("userID", userIDStr)
	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid user ID"))
		return
	}

	user, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(user))
}

// UpdateUserInfo updates the current user's information
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var updateInfo model.UserUpdateInfo
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.userService.UpdateUser(userID.(uint), updateInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(nil))
}
