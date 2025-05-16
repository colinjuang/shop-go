package repository

import "github.com/colinjuang/shop-go/internal/model"

// CategoryRepository handles database operations for categories
type CategoryRepository struct{}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// GetCategories gets all categories
func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	result := DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetCategoriesByParentID gets categories by parent ID
func (r *CategoryRepository) GetCategoriesByParentID(parentID uint64) ([]model.Category, error) {
	var categories []model.Category
	result := DB.Where("parent_id = ?", parentID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
