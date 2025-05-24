package handler

import (
	"net/http"

	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/service"
	"github.com/gin-gonic/gin"
)

// BannerHandler handles home page API endpoints
type BannerHandler struct {
	bannerService *service.BannerService
}

// NewBannerHandler creates a new banner handler
func NewBannerHandler() *BannerHandler {
	return &BannerHandler{
		bannerService: service.NewBannerService(),
	}
}

// GetBanners gets all banners for the carousel
func (h *BannerHandler) GetBanners(c *gin.Context) {
	banners, err := h.bannerService.GetBanners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(banners))
}
