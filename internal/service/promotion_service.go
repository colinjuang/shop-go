package service

import (
	"context"
	"time"

	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/colinjuang/shop-go/internal/response"
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
func (s *PromotionService) GetPromotions() ([]*response.PromotionResponse, error) {
	ctx := context.Background()

	// Try to get from cache
	var promotionResponses []*response.PromotionResponse
	err := s.cacheService.GetObject(ctx, constant.HomePromotions, &promotionResponses)
	if err == nil {
		return promotionResponses, nil
	}

	// If not in cache, get from database
	promotions, err := s.promotionRepo.GetPromotions()
	if err != nil {
		return nil, err
	}

	for i, promotion := range promotions {
		promotionResponses[i] = &response.PromotionResponse{
			ID:            promotion.ID,
			Title:         promotion.Title,
			ImageUrl:      minio.GetClient().GetFileURL(promotion.ImageUrl),
			SubCategoryID: promotion.SubCategoryID,
			CreatedAt:     promotion.CreatedAt,
			UpdatedAt:     promotion.UpdatedAt,
		}
	}

	// Cache for 1 minute
	err = s.cacheService.Set(ctx, constant.HomePromotions, promotions, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache promotions: %v", err)
	}

	return promotionResponses, nil
}
