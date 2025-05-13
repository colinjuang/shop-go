package service

import (
	"context"
	"fmt"
	"shop-go/internal/model"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/redis"
	"shop-go/internal/repository"
	"time"
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
func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("product:%d", id)

	// Try to get from cache
	var product model.Product
	err := s.cacheService.GetObject(ctx, cacheKey, &product)
	if err == nil {
		return &product, nil
	}

	// If not in cache, get from database
	productPtr, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, *productPtr, 1*time.Minute)
	if err != nil {
		// Just log the error, don't fail the request
		logger.Warnf("Failed to cache product: %v", err)
	}

	return productPtr, nil
}

// GetProducts gets products with pagination
func (s *ProductService) GetProducts(page, pageSize int, categoryID *uint, hot, recommend *bool) (*model.Pagination, error) {
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
func (s *ProductService) GetRecommendProducts(limit int) ([]model.Product, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("home:recommend:limit:%d", limit)

	// Try to get from cache
	var products []model.Product
	err := s.cacheService.GetObject(ctx, cacheKey, &products)
	if err == nil {
		return products, nil
	}

	// If not in cache, get from database
	recommend := true
	products, _, err = s.productRepo.GetProducts(1, limit, nil, nil, &recommend)
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, products, 1*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache recommend products: %v\n", err)
	}

	return products, nil
}

// GetHotProducts gets hot products
func (s *ProductService) GetHotProducts(limit int) ([]model.Product, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("home:hot:limit:%d", limit)

	// Try to get from cache
	var products []model.Product
	err := s.cacheService.GetObject(ctx, cacheKey, &products)
	if err == nil {
		return products, nil
	}

	// If not in cache, get from database
	hot := true
	products, _, err = s.productRepo.GetProducts(1, limit, nil, &hot, nil)
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, products, 1*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache hot products: %v\n", err)
	}

	return products, nil
}
