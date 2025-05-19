package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/service"

	"github.com/gin-gonic/gin"
)

// ReportHandler handles report-related API endpoints
type ReportHandler struct {
	reportService *service.ReportService
}

// NewReportHandler creates a new report handler
func NewReportHandler() *ReportHandler {
	return &ReportHandler{
		reportService: service.NewReportService(),
	}
}

// GetProductCatalog generates and returns a PDF product catalog
func (h *ReportHandler) GetProductCatalog(c *gin.Context) {
	// Get category ID from query
	var categoryID *uint64
	if idStr := c.Query("category_id"); idStr != "" {
		if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
			categoryID = &id
		}
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	// Generate catalog
	pdfURL, err := h.reportService.GenerateProductCatalogPDF(ctx, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"url": pdfURL,
	}))
}

// GetOrderInvoice generates and returns a PDF invoice for an order
func (h *ReportHandler) GetOrderInvoice(c *gin.Context) {
	// Get user ID from context
	userID, _ := c.Get("user_id")

	// Get order ID from query
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid order ID"))
		return
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	// Generate invoice
	pdfURL, err := h.reportService.GenerateOrderInvoicePDF(ctx, id, userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"url": pdfURL,
	}))
}

// ExportProducts exports products to a CSV file
func (h *ReportHandler) ExportProducts(c *gin.Context) {
	// Get category ID from query
	var categoryID *uint64
	if idStr := c.Query("category_id"); idStr != "" {
		if id, err := strconv.ParseUint(idStr, 10, 64); err == nil {
			categoryID = &id
		}
	}

	// Set context with timeout
	ctx, cancel := context.WithTimeout(c.Request.Context(), 30*time.Second)
	defer cancel()

	// Export products
	csvURL, err := h.reportService.ExportProductsToCSV(ctx, categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(gin.H{
		"url": csvURL,
	}))
}
