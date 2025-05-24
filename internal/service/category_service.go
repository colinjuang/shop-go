package service

import (
	"context"
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
	"gorm.io/gorm"
)

// CategoryService handles business logic for categories
type CategoryService struct {
	categoryRepo *repository.CategoryRepository
	cacheService *redis.CacheService
}

// NewCategoryService creates a new category service
func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		categoryRepo: repository.NewCategoryRepository(db),
		cacheService: redis.NewCacheService(),
	}
}

// GetCategories gets all categories
func (s *CategoryService) GetCategories() ([]*response.CategoryResponse, error) {
	ctx := context.Background()

	// 从缓存中获取
	var categoryResponses []*response.CategoryResponse
	err := s.cacheService.GetObject(ctx, constant.CategoryList, &categoryResponses)
	if err == nil {
		return categoryResponses, nil
	}

	// 如果不在缓存中，从数据库获取
	categories, err := s.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	// 缓存1分钟
	err = s.cacheService.Set(ctx, constant.CategoryList, categories, 1*time.Minute)
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
	cacheKey := fmt.Sprintf(constant.CategoryParentID+":%d", parentID)

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
	cacheKey := constant.CategoryTree

	// 从缓存中获取
	var tree []*response.CategoryTreeResponse
	err := s.cacheService.GetObject(ctx, cacheKey, &tree)
	if err == nil {
		return tree, nil
	}

	// 如果不在缓存中，从数据库获取
	categories, err := s.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	// 创建一个快速查找的map
	treeMap := make(map[uint64]*response.CategoryTreeResponse)

	// 第一遍：创建所有树节点
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

	// 第二遍：使用map进行O(1)父节点查找
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

	// 缓存1分钟
	err = s.cacheService.Set(ctx, cacheKey, tree, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache categories: %v", err)
	}

	return tree, nil
}
