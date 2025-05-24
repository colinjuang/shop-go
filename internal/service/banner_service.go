package service

import (
	"context"
	"time"

	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/pkg/logger"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/colinjuang/shop-go/internal/server"
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
	server := server.GetServer()
	return &BannerService{
		bannerRepo:   repository.NewBannerRepository(server.DB),
		cacheService: redis.NewCacheService(),
	}
}

// GetBanners gets all banners
func (s *BannerService) GetBanners() ([]*response.BannerResponse, error) {
	ctx := context.Background()

	// Try to get from cache
	var bannerResponses []*response.BannerResponse
	err := s.cacheService.GetObject(ctx, constant.HomeBanners, &bannerResponses)
	if err == nil {
		return bannerResponses, nil
	}

	// If not in cache, get from database
	banners, err := s.bannerRepo.GetBanners()
	if err != nil {
		return nil, err
	}

	minioClient := minio.GetClient()
	for i, banner := range banners {
		banners[i].ImageUrl = minioClient.GetFileURL(banner.ImageUrl)
	}
	// Cache for 1 minute
	err = s.cacheService.Set(ctx, constant.HomeBanners, banners, 1*time.Minute)
	if err != nil {
		logger.Warnf("Failed to cache banners: %v", err)
	}

	bannerResponses = make([]*response.BannerResponse, len(banners))
	for i, banner := range banners {
		bannerResponses[i] = &response.BannerResponse{
			ID:        banner.ID,
			Title:     banner.Title,
			ImageUrl:  banner.ImageUrl,
			ProductID: banner.ProductID,
			SortOrder: banner.SortOrder,
			CreatedAt: banner.CreatedAt,
			UpdatedAt: banner.UpdatedAt,
		}
	}
	return bannerResponses, nil
}
