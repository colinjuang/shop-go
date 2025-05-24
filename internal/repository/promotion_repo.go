package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"gorm.io/gorm"
)

// PromotionRepository 促销仓库
type PromotionRepository struct {
	db *gorm.DB
}

// NewPromotionRepository
func NewPromotionRepository(db *gorm.DB) *PromotionRepository {
	return &PromotionRepository{
		db: db,
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
