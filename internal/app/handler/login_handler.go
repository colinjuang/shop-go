package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/app/request"
	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/service"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	userService *service.UserService
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{
		userService: service.NewUserService(),
	}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var loginRequest request.UserLoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	token, err := h.userService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"token": token,
	}))
}

func (h *LoginHandler) Register(c *gin.Context) {
	var userRegisterRequest request.UserRegisterRequest
	if err := c.ShouldBindJSON(&userRegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.userService.Register(userRegisterRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"message": "用户创建成功",
	}))
}
