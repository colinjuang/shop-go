package service

import (
	"context"
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/api/response"
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
func (s *CategoryService) GetCategories() ([]*response.CategoryResponse, error) {
	ctx := context.Background()
	cacheKey := "categories:all"

	// Try to get from cache
	var categoryResponses []*response.CategoryResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &categoryResponses)
	if err == nil {
		return categoryResponses, nil
	}

	// If not in cache, get from database
	categories, err := s.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, categories, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	categoryResponses = make([]*response.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = &response.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			ParentID:  category.ParentID,
			ImageUrl:  category.ImageUrl,
			SortOrder: category.SortOrder,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
	}
	return categoryResponses, nil
}

// GetCategoriesByParentID gets categories by parent ID
func (s *CategoryService) GetCategoriesByParentID(parentID uint64) ([]*response.CategoryResponse, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("categories:parent:%d", parentID)

	// Try to get from cache
	var categoryResponses []*response.CategoryResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &categoryResponses)
	if err == nil {
		return categoryResponses, nil
	}

	// If not in cache, get from database
	categories, err := s.categoryRepo.GetCategoriesByParentID(parentID)
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, categories, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	categoryResponses = make([]*response.CategoryResponse, len(categories))
	for i, category := range categories {
		categoryResponses[i] = &response.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			ParentID:  category.ParentID,
			ImageUrl:  category.ImageUrl,
			SortOrder: category.SortOrder,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
	}

	return categoryResponses, nil
}

// GetCategoryTree gets the category tree
func (s *CategoryService) GetCategoryTree() ([]*response.CategoryTreeResponse, error) {
	ctx := context.Background()
	cacheKey := "categories:tree"

	// Try to get from cache
	var tree []*response.CategoryTreeResponse
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
	treeMap := make(map[uint64]*response.CategoryTreeResponse)

	// First pass: create all tree nodes
	for _, category := range categories {
		if category.ParentID == 0 {
			node := &response.CategoryTreeResponse{
				ID:        category.ID,
				Name:      category.Name,
				ImageUrl:  category.ImageUrl,
				SortOrder: category.SortOrder,
				Children:  []response.CategoryResponse{},
			}
			tree = append(tree, node)
			treeMap[category.ID] = node
		}
	}

	// Second pass: add children using the map for O(1) parent lookup
	for _, category := range categories {
		if category.ParentID != 0 {
			if parent, exists := treeMap[category.ParentID]; exists {
				parent.Children = append(parent.Children, response.CategoryResponse{
					ID:        category.ID,
					Name:      category.Name,
					ImageUrl:  category.ImageUrl,
					SortOrder: category.SortOrder,
					CreatedAt: category.CreatedAt,
					UpdatedAt: category.UpdatedAt,
				})
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
