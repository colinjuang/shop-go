package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// PromotionRepository 促销仓库
type PromotionRepository struct {
	db *gorm.DB
}

// NewPromotionRepository
func NewPromotionRepository() *PromotionRepository {
	return &PromotionRepository{
		db: database.GetDB(),
	}
}

// GetPromotions 获取所有促销
func (r *PromotionRepository) GetPromotions() ([]model.Promotion, error) {
	var promotions []model.Promotion
	result := r.db.Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}
