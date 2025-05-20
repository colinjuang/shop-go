package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	RequestUserKey = "REQUEST-USER-INFO"
)

type UserClaim struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	City     string `json:"city"`
	Province string `json:"province"`
	District string `json:"district"`
}

// AuthClaims represents JWT claims
type AuthClaims struct {
	AnyJson string `json:"any_json"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token
func GenerateToken(obj any) (string, error) {
	cfg := config.GetConfig()
	anyJson, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	// Set claims
	claims := AuthClaims{
		AnyJson: string(anyJson),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(cfg.JWT.ExpiresIn))), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                                   // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                                   // 生效时间
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate signed token
	tokenString, err := token.SignedString([]byte(cfg.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string) (*AuthClaims, error) {
	cfg := config.GetConfig()

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// AuthMiddleware is a middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, response.ErrorResponse(401, "Missing token"))
			c.Abort()
			return
		}

		// Remove 'Bearer ' prefix if it exists
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token
		claims, err := ParseToken(tokenString)
		if err != nil {
			if err.Error() == "token is expired" {
				c.JSON(200, response.TokenExpiredResponse())
			} else {
				fmt.Println("1111")
				c.JSON(401, response.ErrorResponse(401, "Invalid token"))
			}
			c.Abort()
			return
		}

		var user UserClaim
		err = json.Unmarshal([]byte(claims.AnyJson), &user)
		if err != nil {
			fmt.Println("2222")
			c.JSON(401, response.ErrorResponse(401, "Invalid token"))
			c.Abort()
			return
		}
		SetRequestUser(c, &user)
		c.Next()
	}
}

// SetRequestUser 设置请求用户到上下文
func SetRequestUser(c *gin.Context, user *UserClaim) {
	c.Set(RequestUserKey, user)
}

// GetRequestUser 获取请求用户
func GetRequestUser(c *gin.Context) *UserClaim {
	user, ok := c.Get(RequestUserKey)
	if !ok {
		return nil
	}
	return user.(*UserClaim)
}
