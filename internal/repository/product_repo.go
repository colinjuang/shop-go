package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"gorm.io/gorm"
)

// ProductRepository 商品仓库
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// GetProductByID 获取商品
func (r *ProductRepository) GetProductByID(id uint64) (*model.Product, error) {
	var product model.Product
	result := r.db.First(&product, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

// GetProducts 获取商品
func (r *ProductRepository) GetProducts(page, pageSize int, categoryID *uint64, hot, recommend *bool) ([]model.Product, int64, error) {
	var products []model.Product
	var count int64

	query := r.db.Model(&model.Product{})

	// 应用过滤
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if hot != nil {
		query = query.Where("sale_count > 100")
	}

	if recommend != nil {
		query = query.Where("recommend = ?", *recommend)
	}

	// 获取总数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页结果
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

// UpdateProductStock 更新商品库存
func (r *ProductRepository) UpdateProductStock(id uint64, stock int) error {
	// 减去库存
	return r.db.Model(&model.Product{}).Where("id = ?", id).Update("stock_count", gorm.Expr("stock_count - ?", stock)).Error
}
