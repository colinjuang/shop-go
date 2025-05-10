package handler

import (
	"net/http"
	"shop-go/internal/model"
	"shop-go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HomeHandler handles home page API endpoints
type HomeHandler struct {
	homeService     *service.HomeService
	categoryService *service.CategoryService
}

// NewHomeHandler creates a new home handler
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{
		homeService:     service.NewHomeService(),
		categoryService: service.NewCategoryService(),
	}
}

// GetBanners gets all banners for the carousel
func (h *HomeHandler) GetBanners(c *gin.Context) {
	banners, err := h.homeService.GetBanners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(banners))
}

// GetCategories gets top-level categories
func (h *HomeHandler) GetCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategoriesByParentID(0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(categories))
}

// GetAllCategories gets all categories
func (h *HomeHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(categories))
}

// GetSubCategories gets subcategories for a category
func (h *HomeHandler) GetSubCategories(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	categories, err := h.categoryService.GetCategoriesByParentID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(categories))
}

// GetPromotions gets all promotions
func (h *HomeHandler) GetPromotions(c *gin.Context) {
	promotions, err := h.homeService.GetPromotions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(promotions))
}

// GetRecommendProducts gets recommended products
func (h *HomeHandler) GetRecommendProducts(c *gin.Context) {
	limit := 10
	limitStr := c.DefaultQuery("limit", "10")
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	products, err := h.homeService.GetRecommendProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(products))
}

// GetHotProducts gets hot products
func (h *HomeHandler) GetHotProducts(c *gin.Context) {
	limit := 10
	limitStr := c.DefaultQuery("limit", "10")
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	products, err := h.homeService.GetHotProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(products))
}
