package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"
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
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(banners))
}
