package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		db: database.GetDB(),
	}
}

// GetProductByID gets a product by ID
func (r *ProductRepository) GetProductByID(id uint64) (*model.Product, error) {
	var product model.Product
	result := r.db.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// GetProducts gets products with pagination
func (r *ProductRepository) GetProducts(page, pageSize int, categoryID *uint64, hot, recommend *bool) ([]model.Product, int64, error) {
	var products []model.Product
	var count int64

	query := r.db.Model(&model.Product{})

	// Apply filters
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if hot != nil {
		query = query.Where("sale_count > 100")
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
