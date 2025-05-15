package repository

import "github.com/colinjuang/shop-go/internal/model"

// PromotionRepository handles database operations for promotions
type PromotionRepository struct{}

// NewPromotionRepository creates a new promotion repository
func NewPromotionRepository() *PromotionRepository {
	return &PromotionRepository{}
}

// GetPromotions gets all promotions
func (r *PromotionRepository) GetPromotions() ([]model.Promotion, error) {
	var promotions []model.Promotion
	result := DB.Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}
