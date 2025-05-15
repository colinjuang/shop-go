package handler

import (
	"net/http"
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

// WechatLogin handles WeChat login
func (h *UserHandler) WechatLogin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Missing code"))
		return
	}

	// Here we would normally call WeChat API to get user info
	// For simplicity, we'll use the code as OpenID
	openID := code
	nickname := c.DefaultQuery("nickname", "User")
	avatar := c.DefaultQuery("avatar", "")
	gender := 0 // Default: unknown
	city := c.DefaultQuery("city", "")
	province := c.DefaultQuery("province", "")
	country := c.DefaultQuery("country", "")

	token, err := h.userService.LoginWithWechat(openID, nickname, avatar, gender, city, province, country)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"token": token,
	}))
}

// GetUserInfo gets the current user's information
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
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
