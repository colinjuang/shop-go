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
