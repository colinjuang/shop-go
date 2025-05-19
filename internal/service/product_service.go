package service

import (
	"context"
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
)

// ProductService handles business logic for products
type ProductService struct {
	productRepo  *repository.ProductRepository
	cacheService *redis.CacheService
}

// NewProductService creates a new product service
func NewProductService() *ProductService {
	return &ProductService{
		productRepo:  repository.NewProductRepository(),
		cacheService: redis.NewCacheService(),
	}
}

// GetProductByID gets a product by ID
func (s *ProductService) GetProductByID(id uint64) (*response.ProductResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf(constant.ProductPrefix+":%d", id)

	// Try to get from cache
	var productResponse response.ProductResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &productResponse)
	if err == nil {
		return &productResponse, nil
	}

	// If not in cache, get from database
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	productResponse = response.ProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		FloralLanguage: product.FloralLanguage,
		Price:          product.Price,
		MarketPrice:    product.MarketPrice,
		SaleCount:      product.SaleCount,
		StockCount:     product.StockCount,
		CategoryID:     product.CategoryID,
		SubCategoryID:  product.SubCategoryID,
		Material:       product.Material,
		Packing:        product.Packing,
		ImageUrl:       product.ImageUrl,
		Status:         product.Status,
		Recommend:      product.Recommend,
		SortOrder:      product.SortOrder,
		ApplyUser:      product.ApplyUser,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, product, 1*time.Minute)
	if err != nil {
		// Just log the error, don't fail the request
		logger.Warnf("Failed to cache product: %v", err)
	}

	return &productResponse, nil
}

// GetProducts gets products with pagination
func (s *ProductService) GetProducts(page, pageSize int, categoryID *uint64, hot, recommend *bool) (*response.Pagination, error) {
	// If not in cache, get from database
	products, total, err := s.productRepo.GetProducts(page, pageSize, categoryID, hot, recommend)
	if err != nil {
		return nil, err
	}

	pagination := response.NewPagination(total, page, pageSize, products)

	return &pagination, nil
}

// GetRecommendProducts gets recommended products
func (s *ProductService) GetRecommendProducts(limit int) ([]*response.ProductResponse, error) {
	// Try to get from cache
	var productResponses []*response.ProductResponse

	// If not in cache, get from database
	recommend := true
	products, _, err := s.productRepo.GetProducts(1, limit, nil, nil, &recommend)
	if err != nil {
		return nil, err
	}

	productResponses = make([]*response.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = &response.ProductResponse{
			ID:             product.ID,
			Name:           product.Name,
			FloralLanguage: product.FloralLanguage,
			Price:          product.Price,
			MarketPrice:    product.MarketPrice,
			SaleCount:      product.SaleCount,
			StockCount:     product.StockCount,
			CategoryID:     product.CategoryID,
			SubCategoryID:  product.SubCategoryID,
			Material:       product.Material,
			Packing:        product.Packing,
			ImageUrl:       product.ImageUrl,
			Status:         product.Status,
			Recommend:      product.Recommend,
			SortOrder:      product.SortOrder,
			ApplyUser:      product.ApplyUser,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		}
	}

	return productResponses, nil
}

// GetHotProducts gets hot products
func (s *ProductService) GetHotProducts(limit int) ([]*response.ProductResponse, error) {
	// Try to get from cache
	var productResponses []*response.ProductResponse

	// If not in cache, get from database
	hot := true
	products, _, err := s.productRepo.GetProducts(1, limit, nil, &hot, nil)
	if err != nil {
		return nil, err
	}

	productResponses = make([]*response.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = &response.ProductResponse{
			ID:             product.ID,
			Name:           product.Name,
			FloralLanguage: product.FloralLanguage,
			Price:          product.Price,
			MarketPrice:    product.MarketPrice,
			SaleCount:      product.SaleCount,
			StockCount:     product.StockCount,
			CategoryID:     product.CategoryID,
			SubCategoryID:  product.SubCategoryID,
			Material:       product.Material,
			Packing:        product.Packing,
			ImageUrl:       product.ImageUrl,
			Status:         product.Status,
			Recommend:      product.Recommend,
			SortOrder:      product.SortOrder,
			ApplyUser:      product.ApplyUser,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
		}
	}

	return productResponses, nil
}
