package service

import (
	"context"
	"fmt"
	"os"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"time"
)

// ReportService handles business logic for generating reports
type ReportService struct {
	productService *ProductService
	orderService   *OrderService
}

// NewReportService creates a new report service
func NewReportService() *ReportService {
	return &ReportService{
		productService: NewProductService(),
		orderService:   NewOrderService(),
	}
}

// GenerateProductCatalogPDF generates a PDF catalog of products
// This is just an example that simulates PDF generation
func (s *ReportService) GenerateProductCatalogPDF(ctx context.Context, categoryID *uint) (string, error) {
	// Create a cache key based on parameters
	key := fmt.Sprintf("product_catalog")
	if categoryID != nil {
		key += fmt.Sprintf("_cat_%d", *categoryID)
	}
	key += fmt.Sprintf("_%s", time.Now().Format("2006-01-02"))

	// Set cache options
	opts := minio.DefaultCacheOptions()
	opts.Prefix = "reports"
	opts.ContentType = "application/pdf"
	opts.TTL = 24 * time.Hour // Daily reports

	// Get cached file or generate a new one
	return minio.GetCachedFileWithTempFile(ctx, key, opts, func(tempPath string) error {
		// This simulates generating a PDF file
		// In a real application, you would use a PDF library like gofpdf

		// Get products from category
		pagination, err := s.productService.GetProducts(1, 100, categoryID, nil, nil)
		if err != nil {
			return err
		}

		// Create a simple text file to simulate PDF content
		file, err := os.Create(tempPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write header
		file.WriteString("PRODUCT CATALOG\n")
		file.WriteString("==============\n\n")
		file.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))

		// Write products
		products, ok := pagination.Data.([]model.Product)
		if !ok {
			return fmt.Errorf("unexpected data type in pagination")
		}

		for _, product := range products {
			file.WriteString(fmt.Sprintf("Product: %s\n", product.Name))
			file.WriteString(fmt.Sprintf("Price: %.2f\n", product.Price))
			file.WriteString(fmt.Sprintf("Description: %s\n\n", product.FloralLanguage))
		}

		return nil
	})
}

// GenerateOrderInvoicePDF generates a PDF invoice for an order
func (s *ReportService) GenerateOrderInvoicePDF(ctx context.Context, orderID uint, userID uint) (string, error) {
	// Get order
	order, err := s.orderService.GetOrderByID(orderID, userID)
	if err != nil {
		return "", err
	}

	// Create a cache key based on order ID and updated timestamp
	key := fmt.Sprintf("order_invoice_%d_%d", order.ID, order.UpdatedAt.Unix())

	// Set cache options
	opts := minio.DefaultCacheOptions()
	opts.Prefix = "invoices"
	opts.ContentType = "application/pdf"
	opts.TTL = 30 * 24 * time.Hour // 30 days

	// Get cached file or generate a new one
	return minio.GetCachedFileWithTempFile(ctx, key, opts, func(tempPath string) error {
		// This simulates generating a PDF invoice
		// In a real application, you would use a PDF library like gofpdf

		// Create a simple text file to simulate PDF content
		file, err := os.Create(tempPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write header
		file.WriteString("INVOICE\n")
		file.WriteString("=======\n\n")
		file.WriteString(fmt.Sprintf("Order Number: %s\n", order.OrderNo))
		file.WriteString(fmt.Sprintf("Date: %s\n", order.CreatedAt.Format("2006-01-02")))
		file.WriteString(fmt.Sprintf("Customer ID: %d\n\n", order.UserID))

		// Shipping address
		file.WriteString("Shipping Address:\n")
		file.WriteString(fmt.Sprintf("%s\n", order.ReceiverName))
		file.WriteString(fmt.Sprintf("%s\n", order.ReceiverPhone))
		file.WriteString(fmt.Sprintf("%s\n\n", order.Address))

		// Order items
		file.WriteString("Items:\n")
		for _, item := range order.OrderItems {
			file.WriteString(fmt.Sprintf("%s - Qty: %d - Price: %.2f - Total: %.2f\n",
				item.Name, item.Quantity, item.Price, float64(item.Quantity)*item.Price))
		}
		file.WriteString("\n")

		// Totals
		file.WriteString(fmt.Sprintf("Subtotal: %.2f\n", order.TotalAmount))
		file.WriteString(fmt.Sprintf("Total: %.2f\n", order.PaymentAmount))

		return nil
	})
}

// ExportProductsToCSV exports products to a CSV file
func (s *ReportService) ExportProductsToCSV(ctx context.Context, categoryID *uint) (string, error) {
	// Create a cache key based on parameters
	key := fmt.Sprintf("products_export")
	if categoryID != nil {
		key += fmt.Sprintf("_cat_%d", *categoryID)
	}
	key += fmt.Sprintf("_%s", time.Now().Format("2006-01-02"))

	// Set cache options
	opts := minio.DefaultCacheOptions()
	opts.Prefix = "exports"
	opts.ContentType = "text/csv"
	opts.TTL = 24 * time.Hour

	// Get cached file or generate a new one
	return minio.GetCachedFileWithTempFile(ctx, key, opts, func(tempPath string) error {
		// Get products from category
		pagination, err := s.productService.GetProducts(1, 1000, categoryID, nil, nil)
		if err != nil {
			return err
		}

		// Create CSV file
		file, err := os.Create(tempPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write header
		file.WriteString("ID,Name,Price,Stock,Description,Category\n")

		// Write products
		products, ok := pagination.Data.([]model.Product)
		if !ok {
			return fmt.Errorf("unexpected data type in pagination")
		}

		for _, product := range products {
			file.WriteString(fmt.Sprintf("%d,%s,%.2f,%d,%s,%d\n",
				product.ID, product.Name, product.Price, product.StockCount,
				product.FloralLanguage, product.CategoryID))
		}

		return nil
	})
}
