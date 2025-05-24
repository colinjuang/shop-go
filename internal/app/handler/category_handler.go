package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// CategoryHandler handles home page API endpoints
type CategoryHandler struct {
	categoryService *service.CategoryService
}

// NewCategoryHandler creates a new banner handler
func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{
		categoryService: service.NewCategoryService(db),
	}
}

// GetLevel1Categories gets level 1 categories
func (h *CategoryHandler) GetLevel1Categories(c *gin.Context) {
	categories, err := h.categoryService.GetCategoriesByParentID(0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(categories))
}

// GetAllCategories gets all categories
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(categories))
}

// GetSubCategories gets subcategories for a category
func (h *CategoryHandler) GetSubCategories(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid ID"))
		return
	}

	categories, err := h.categoryService.GetCategoriesByParentID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(categories))
}

// GetCategoryTree gets the category tree
func (h *CategoryHandler) GetCategoryTree(c *gin.Context) {
	tree, err := h.categoryService.GetCategoryTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(tree))
}
