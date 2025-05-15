package handler

import (
	"net/http"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles product-related API endpoints
type ProductHandler struct {
	productService *service.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productService: service.NewProductService(),
	}
}

// GetProductDetail gets a product by ID
func (h *ProductHandler) GetProductDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(product))
}

// GetProducts gets products with pagination
func (h *ProductHandler) GetProducts(c *gin.Context) {
	// Get query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// Apply minimum values
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	// Get category ID filter
	var categoryID *uint
	if idStr := c.Query("category_id"); idStr != "" {
		if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
			catID := uint(id)
			categoryID = &catID
		}
	}

	// Get hot filter
	var hot *bool
	if hotStr := c.Query("hot"); hotStr != "" {
		if hotStr == "1" || hotStr == "true" {
			hotVal := true
			hot = &hotVal
		} else if hotStr == "0" || hotStr == "false" {
			hotVal := false
			hot = &hotVal
		}
	}

	// Get recommend filter
	var recommend *bool
	if recStr := c.Query("recommend"); recStr != "" {
		if recStr == "1" || recStr == "true" {
			recVal := true
			recommend = &recVal
		} else if recStr == "0" || recStr == "false" {
			recVal := false
			recommend = &recVal
		}
	}

	// Get products
	pagination, err := h.productService.GetProducts(page, pageSize, categoryID, hot, recommend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(pagination))
}

// GetRecommendProducts gets recommended products
func (h *ProductHandler) GetRecommendProducts(c *gin.Context) {
	limit := 10
	limitStr := c.DefaultQuery("limit", "10")
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	products, err := h.productService.GetRecommendProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(products))
}

// GetHotProducts gets hot products
func (h *ProductHandler) GetHotProducts(c *gin.Context) {
	limit := 10
	limitStr := c.DefaultQuery("limit", "10")
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	products, err := h.productService.GetHotProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(products))
}
