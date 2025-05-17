package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// BannerRepository 轮播图仓库
type BannerRepository struct {
	db *gorm.DB
}

// NewBannerRepository
func NewBannerRepository() *BannerRepository {
	return &BannerRepository{
		db: database.GetDB(),
	}
}

// GetBanners 获取所有轮播图
func (r *BannerRepository) GetBanners() ([]model.Banner, error) {
	var banners []model.Banner
	result := r.db.Order("sort_order ASC").Find(&banners)
	if result.Error != nil {
		return nil, result.Error
	}
	return banners, nil
}
