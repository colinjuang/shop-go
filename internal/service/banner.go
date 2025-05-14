package service

import (
	"context"
	"shop-go/internal/model"
	"shop-go/internal/pkg/logger"
	"shop-go/internal/pkg/minio"
	"shop-go/internal/pkg/redis"
	"shop-go/internal/repository"
	"time"
)

// BannerService handles business logic for the banner page
type BannerService struct {
	bannerRepo    *repository.BannerRepository
	promotionRepo *repository.PromotionRepository
	productRepo   *repository.ProductRepository
	cacheService  *redis.CacheService
}

// NewBannerService creates a new banner service
func NewBannerService() *BannerService {
	return &BannerService{
		bannerRepo:   repository.NewBannerRepository(),
		cacheService: redis.NewCacheService(),
	}
}

// GetBanners gets all banners
func (s *BannerService) GetBanners() ([]model.Banner, error) {
	ctx := context.Background()
	cacheKey := "home:banners"

	// Try to get from cache
	var banners []model.Banner
	err := s.cacheService.GetObject(ctx, cacheKey, &banners)
	if err == nil {
		return banners, nil
	}

	// If not in cache, get from database
	banners, err = s.bannerRepo.GetBanners()
	if err != nil {
		return nil, err
	}

	minioClient := minio.GetClient()
	for i, banner := range banners {
		imageUrl := minioClient.GetFileURL(banner.ImageUrl)
		banners[i].ImageUrl = imageUrl
	}
	// Cache for 1 minute
	err = s.cacheService.Set(ctx, cacheKey, banners, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache banners: %v", err)
	}

	return banners, nil
}
