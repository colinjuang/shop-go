package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// BannerRepository handles database operations for banners
type BannerRepository struct {
	db *gorm.DB
}

// NewBannerRepository creates a new banner repository
func NewBannerRepository() *BannerRepository {
	return &BannerRepository{
		db: database.GetDB(),
	}
}

// GetBanners gets all banners
func (r *BannerRepository) GetBanners() ([]model.Banner, error) {
	var banners []model.Banner
	result := r.db.Order("sort_order ASC").Find(&banners)
	if result.Error != nil {
		return nil, result.Error
	}
	return banners, nil
}
