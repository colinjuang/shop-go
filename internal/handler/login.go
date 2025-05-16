package handler

import (
	"net/http"
	"time"

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
	// username and password
	username := c.Query("username")
	password := c.Query("password")

	// check if username and password are correct
	user, err := h.userService.GetUserByUsername(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Invalid username or password"))
		return
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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
	// json body
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// check if username already exists
	user, err := h.userService.GetUserByUsername(json.Username)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, "Failed to hash password"))
		return
	}

	// create user
	user = &model.User{
		Username: json.Username,
		Password: string(hashedPassword),
		OpenID:   "",
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
