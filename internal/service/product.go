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

// CategoryService handles business logic for categories
type CategoryService struct {
	categoryRepo *repository.CategoryRepository
	cacheService *redis.CacheService
}

// NewCategoryService creates a new category service
func NewCategoryService() *CategoryService {
	return &CategoryService{
		categoryRepo: repository.NewCategoryRepository(),
		cacheService: redis.NewCacheService(),
	}
}

// GetCategories gets all categories
func (s *CategoryService) GetCategories() ([]model.Category, error) {
	ctx := context.Background()
	cacheKey := "categories:all"

	// Try to get from cache
	var categories []model.Category
	err := s.cacheService.GetObject(ctx, cacheKey, &categories)
	if err == nil {
		return categories, nil
	}

	// If not in cache, get from database
	categories, err = s.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, categories, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	return categories, nil
}

// GetCategoriesByParentID gets categories by parent ID
func (s *CategoryService) GetCategoriesByParentID(parentID uint) ([]model.Category, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("categories:parent:%d", parentID)

	// Try to get from cache
	var categories []model.Category
	err := s.cacheService.GetObject(ctx, cacheKey, &categories)
	if err == nil {
		return categories, nil
	}

	// If not in cache, get from database
	categories, err = s.categoryRepo.GetCategoriesByParentID(parentID)
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, categories, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	return categories, nil
}

// HomeService handles business logic for the home page
type HomeService struct {
	bannerRepo    *repository.BannerRepository
	promotionRepo *repository.PromotionRepository
	productRepo   *repository.ProductRepository
	cacheService  *redis.CacheService
}

// NewHomeService creates a new home service
func NewHomeService() *HomeService {
	return &HomeService{
		bannerRepo:    repository.NewBannerRepository(),
		promotionRepo: repository.NewPromotionRepository(),
		productRepo:   repository.NewProductRepository(),
		cacheService:  redis.NewCacheService(),
	}
}

// GetBanners gets all banners
func (s *HomeService) GetBanners() ([]model.Banner, error) {
	ctx := context.Background()
	cacheKey := "home:banners"

	// Try to get from cache
	var banners []model.Banner
	err := s.cacheService.GetObject(ctx, cacheKey, &banners)
	if err == nil {
		return banners, nil
	}

	// If not in cache, get from database
	banners, err = s.bannerRepo.GetBanners()
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, banners, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache banners: %v", err)
	}

	return banners, nil
}

// GetPromotions gets all promotions
func (s *HomeService) GetPromotions() ([]model.Promotion, error) {
	ctx := context.Background()
	cacheKey := "home:promotions"

	// Try to get from cache
	var promotions []model.Promotion
	err := s.cacheService.GetObject(ctx, cacheKey, &promotions)
	if err == nil {
		return promotions, nil
	}

	// If not in cache, get from database
	promotions, err = s.promotionRepo.GetPromotions()
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, promotions, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache promotions: %v", err)
	}

	return promotions, nil
}

// GetRecommendProducts gets recommended products
func (s *HomeService) GetRecommendProducts(limit int) ([]model.Product, error) {
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
func (s *HomeService) GetHotProducts(limit int) ([]model.Product, error) {
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
