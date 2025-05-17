package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// PromotionRepository handles database operations for promotions
type PromotionRepository struct {
	db *gorm.DB
}

// NewPromotionRepository creates a new promotion repository
func NewPromotionRepository() *PromotionRepository {
	return &PromotionRepository{
		db: database.GetDB(),
	}
}

// GetPromotions gets all promotions
func (r *PromotionRepository) GetPromotions() ([]model.Promotion, error) {
	var promotions []model.Promotion
	result := r.db.Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}
