package service

import (
	"context"
	"shop-go/internal/model"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/redis"
	"shop-go/internal/repository"
	"time"
)

// PromotionService handles business logic for the home page
type PromotionService struct {
	promotionRepo *repository.PromotionRepository
	cacheService  *redis.CacheService
}

// NewPromotionService creates a new home service
func NewPromotionService() *PromotionService {
	return &PromotionService{
		promotionRepo: repository.NewPromotionRepository(),
		cacheService:  redis.NewCacheService(),
	}
}

// GetPromotions gets all promotions
func (s *PromotionService) GetPromotions() ([]model.Promotion, error) {
	ctx := context.Background()
	cacheKey := "home:promotions"

	// Try to get from cache
	var promotions []model.Promotion
	err := s.cacheService.GetObject(ctx, cacheKey, &promotions)
	if err == nil {
		return promotions, nil
	}

	// If not in cache, get from database
	promotions, err = s.promotionRepo.GetPromotions()
	if err != nil {
		return nil, err
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, promotions, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache promotions: %v", err)
	}

	return promotions, nil
}
