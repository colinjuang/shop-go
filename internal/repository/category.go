package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// CategoryRepository handles database operations for categories
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: database.GetDB(),
	}
}

// GetCategories gets all categories
func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetCategoriesByParentID gets categories by parent ID
func (r *CategoryRepository) GetCategoriesByParentID(parentID uint64) ([]model.Category, error) {
	var categories []model.Category
	result := r.db.Where("parent_id = ?", parentID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
