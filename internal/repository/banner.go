package repository

import "github.com/colinjuang/shop-go/internal/model"

// BannerRepository handles database operations for banners
type BannerRepository struct{}

// NewBannerRepository creates a new banner repository
func NewBannerRepository() *BannerRepository {
	return &BannerRepository{}
}

// GetBanners gets all banners
func (r *BannerRepository) GetBanners() ([]model.Banner, error) {
	var banners []model.Banner
	result := DB.Order("sort_order ASC").Find(&banners)
	if result.Error != nil {
		return nil, result.Error
	}
	return banners, nil
}
