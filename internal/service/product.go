package service

import (
	"context"
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/dto"
	"github.com/colinjuang/shop-go/internal/model"
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
func (s *ProductService) GetProductByID(id uint64) (*dto.ProductResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("product:%d", id)

	// Try to get from cache
	var productResponse dto.ProductResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &productResponse)
	if err == nil {
		return &productResponse, nil
	}

	// If not in cache, get from database
	product, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	productResponse = dto.ProductResponse{
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
func (s *ProductService) GetProducts(page, pageSize int, categoryID *uint64, hot, recommend *bool) (*model.Pagination, error) {
	ctx := context.Background()

	// Generate cache key based on filters
	cacheKey := fmt.Sprintf("products:page:%d:size:%d", page, pageSize)
	if categoryID != nil {
		cacheKey += fmt.Sprintf(":cat:%d", *categoryID)
	}
	if hot != nil {
		cacheKey += fmt.Sprintf(":hot:%t", *hot)
	}
	if recommend != nil {
		cacheKey += fmt.Sprintf(":rec:%t", *recommend)
	}

	// Try to get from cache
	var pagination model.Pagination
	err := s.cacheService.GetObject(ctx, cacheKey, &pagination)
	if err == nil {
		return &pagination, nil
	}

	// If not in cache, get from database
	products, total, err := s.productRepo.GetProducts(page, pageSize, categoryID, hot, recommend)
	if err != nil {
		return nil, err
	}

	pagination = model.NewPagination(total, page, pageSize, products)

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, pagination, 1*time.Minute)
	if err != nil {
		// Just log the error, don't fail the request
		logger.Warnf("Failed to cache products: %v", err)
	}

	return &pagination, nil
}

// GetRecommendProducts gets recommended products
func (s *ProductService) GetRecommendProducts(limit int) ([]*dto.ProductResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("home:recommend:limit:%d", limit)

	// Try to get from cache
	var productResponses []*dto.ProductResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &productResponses)
	if err == nil {
		return productResponses, nil
	}

	// If not in cache, get from database
	recommend := true
	products, _, err := s.productRepo.GetProducts(1, limit, nil, nil, &recommend)
	if err != nil {
		return nil, err
	}

	productResponses = make([]*dto.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = &dto.ProductResponse{
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

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, products, 1*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache recommend products: %v\n", err)
	}

	return productResponses, nil
}

// GetHotProducts gets hot products
func (s *ProductService) GetHotProducts(limit int) ([]*dto.ProductResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("home:hot:limit:%d", limit)

	// Try to get from cache
	var productResponses []*dto.ProductResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &productResponses)
	if err == nil {
		return productResponses, nil
	}

	// If not in cache, get from database
	hot := true
	products, _, err := s.productRepo.GetProducts(1, limit, nil, &hot, nil)
	if err != nil {
		return nil, err
	}

	productResponses = make([]*dto.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = &dto.ProductResponse{
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

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, productResponses, 1*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache hot products: %v\n", err)
	}

	return productResponses, nil
}
