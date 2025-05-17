package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// CategoryRepository 分类仓库
type CategoryRepository struct {
	db *gorm.DB
}

// NewCategoryRepository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		db: database.GetDB(),
	}
}

// GetCategories 获取所有分类
func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetCategoriesByParentID 获取父ID分类
func (r *CategoryRepository) GetCategoriesByParentID(parentID uint64) ([]model.Category, error) {
	var categories []model.Category
	result := r.db.Where("parent_id = ?", parentID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
