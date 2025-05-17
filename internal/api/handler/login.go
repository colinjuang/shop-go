package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/api/middleware"
	"github.com/colinjuang/shop-go/internal/api/request"
	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"
	"golang.org/x/crypto/bcrypt"

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

	// check if username and password are correct
	user, err := h.userService.GetUserByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Invalid username or password"))
		return
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Invalid username or password"))
		return
	}

	// generate token
	token, err := middleware.GenerateToken(middleware.UserClaim{
		UserID: user.ID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to generate token"))
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

	// check if username already exists
	user, err := h.userService.GetUserByUsername(userRegisterRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to get user"))
		return
	}

	// check if password is correct
	if user != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Username already exists"))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to hash password"))
		return
	}

	// create user
	user = &model.User{
		Username: userRegisterRequest.Username,
		Password: string(hashedPassword),
		Nickname: userRegisterRequest.Nickname,
		Avatar:   userRegisterRequest.Avatar,
		Gender:   userRegisterRequest.Gender,
	}

	// save user
	err = h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Failed to create user"))
		return
	}

	// return success
	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"message": "User created successfully",
	}))
}
