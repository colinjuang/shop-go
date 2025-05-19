package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// CartRepository 购物车仓库
type CartRepository struct {
	db *gorm.DB
}

// NewCartRepository
func NewCartRepository() *CartRepository {
	return &CartRepository{
		db: database.GetDB(),
	}
}

// AddToCart 添加商品到购物车
func (r *CartRepository) AddToCart(userID uint64, productID uint64, quantity int) error {
	// 检查商品是否已存在
	var existingItem model.Cart
	result := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingItem)

	if result.Error == nil {
		// 如果商品已存在，则更新数量
		existingItem.Quantity += quantity
		return r.db.Save(&existingItem).Error
	}

	// 如果商品不存在，则添加新商品
	cartItem := model.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Selected:  true,
	}

	return r.db.Create(&cartItem).Error
}

// GetCart 获取用户所有购物车商品
func (r *CartRepository) GetCart(userID uint64) ([]model.Cart, error) {
	var cart []model.Cart
	result := r.db.Where("user_id = ?", userID).Preload("Product").Find(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

// UpdateCartStatus 更新购物车商品状态
func (r *CartRepository) UpdateCartStatus(id uint64, selected bool) error {
	return r.db.Model(&model.Cart{}).Where("id = ?", id).Update("selected", selected).Error
}

// UpdateAllCartStatus 更新用户所有购物车商品状态
func (r *CartRepository) UpdateAllCartStatus(userID uint64, selected bool) error {
	return r.db.Model(&model.Cart{}).Where("user_id = ?", userID).Update("selected", selected).Error
}

// DeleteCart 删除购物车商品
func (r *CartRepository) DeleteCart(id uint64) error {
	return r.db.Delete(&model.Cart{}, "id = ?", id).Error
}

// GetCartByID 获取购物车商品
func (r *CartRepository) GetCartByID(id uint64) (*model.Cart, error) {
	var cart model.Cart
	result := r.db.Where("id = ?", id).Preload("Product").First(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cart, nil
}

// GetCartsByIDs 获取购物车商品
func (r *CartRepository) GetCartsByIDs(ids []uint64) ([]model.Cart, error) {
	var cart []model.Cart
	result := r.db.Where("id IN ?", ids).Preload("Product").Find(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

// GetSelectedCarts 获取用户选中的购物车商品
func (r *CartRepository) GetSelectedCarts(userID uint64) ([]model.Cart, error) {
	var cart []model.Cart
	result := r.db.Where("user_id = ? AND selected = ?", userID, true).Preload("Product").Find(&cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}
