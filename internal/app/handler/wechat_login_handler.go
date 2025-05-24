package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
)

type WechatLoginHandler struct {
	wechatLoginService *service.WechatLoginService
}

func NewWechatLoginHandler() *WechatLoginHandler {
	return &WechatLoginHandler{
		wechatLoginService: service.NewWechatLoginService(),
	}
}

func (h *WechatLoginHandler) WechatMiniLogin(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}
	token, err := h.wechatLoginService.WechatMiniLogin(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
