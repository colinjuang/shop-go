package service

import (
	"context"
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
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

// GetCategoryTree gets the category tree
func (s *CategoryService) GetCategoryTree() ([]*model.CategoryTree, error) {
	ctx := context.Background()
	cacheKey := "categories:tree"

	// Try to get from cache
	var tree []*model.CategoryTree
	err := s.cacheService.GetObject(ctx, cacheKey, &tree)
	if err == nil {
		return tree, nil
	}

	// If not in cache, get from database
	categories, err := s.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	// Create a map for quick lookup
	treeMap := make(map[uint64]*model.CategoryTree)

	// First pass: create all tree nodes
	for _, category := range categories {
		if category.ParentID == 0 {
			node := &model.CategoryTree{
				ID:        category.ID,
				Name:      category.Name,
				ImageUrl:  category.ImageUrl,
				SortOrder: category.SortOrder,
				Children:  []model.Category{},
			}
			tree = append(tree, node)
			treeMap[category.ID] = node
		}
	}

	// Second pass: add children using the map for O(1) parent lookup
	for _, category := range categories {
		if category.ParentID != 0 {
			if parent, exists := treeMap[category.ParentID]; exists {
				parent.Children = append(parent.Children, category)
			}
		}
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, tree, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	return tree, nil
}
