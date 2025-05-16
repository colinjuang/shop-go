package handler

import (
	"net/http"
	"time"

	"github.com/colinjuang/shop-go/internal/dto"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/golang-jwt/jwt/v4"
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
	var loginRequest dto.UserLoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// check if username and password are correct
	user, err := h.userService.GetUserByUsername(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Invalid username or password"))
		return
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Invalid username or password"))
		return
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"openID": user.OpenID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// sign token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to generate token"))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"token": tokenString,
	}))
}

func (h *LoginHandler) Register(c *gin.Context) {
	var userRegisterRequest dto.UserRegisterRequest

	if err := c.ShouldBindJSON(&userRegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// check if username already exists
	user, err := h.userService.GetUserByUsername(userRegisterRequest.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to get user"))
		return
	}

	// check if password is correct
	if user != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Username already exists"))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to hash password"))
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
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to create user"))
		return
	}

	// return success
	c.JSON(http.StatusOK, model.SuccessResponse(gin.H{
		"message": "User created successfully",
	}))
}
