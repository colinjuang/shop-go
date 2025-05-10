package repository

import (
	"shop-go/internal/model"
)

// ProductRepository handles database operations for products
type ProductRepository struct{}

// NewProductRepository creates a new product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{}
}

// GetProductByID gets a product by ID
func (r *ProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	result := DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// GetProducts gets products with pagination
func (r *ProductRepository) GetProducts(page, pageSize int, categoryID *uint, hot, recommend *bool) ([]model.Product, int64, error) {
	var products []model.Product
	var count int64

	query := DB.Model(&model.Product{})

	// Apply filters
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if hot != nil {
		query = query.Where("hot = ?", *hot)
	}

	if recommend != nil {
		query = query.Where("recommend = ?", *recommend)
	}

	// Get total count
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

// CategoryRepository handles database operations for categories
type CategoryRepository struct{}

// NewCategoryRepository creates a new category repository
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// GetCategories gets all categories
func (r *CategoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	result := DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

// GetCategoriesByParentID gets categories by parent ID
func (r *CategoryRepository) GetCategoriesByParentID(parentID uint) ([]model.Category, error) {
	var categories []model.Category
	result := DB.Where("parent_id = ?", parentID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

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

// PromotionRepository handles database operations for promotions
type PromotionRepository struct{}

// NewPromotionRepository creates a new promotion repository
func NewPromotionRepository() *PromotionRepository {
	return &PromotionRepository{}
}

// GetPromotions gets all promotions
func (r *PromotionRepository) GetPromotions() ([]model.Promotion, error) {
	var promotions []model.Promotion
	result := DB.Order("sort_order ASC").Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}
	return promotions, nil
}
