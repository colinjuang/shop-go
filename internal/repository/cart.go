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
	var existingItem model.CartItem
	result := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingItem)

	if result.Error == nil {
		// 如果商品已存在，则更新数量
		existingItem.Quantity += quantity
		return r.db.Save(&existingItem).Error
	}

	// 如果商品不存在，则添加新商品
	cartItem := model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Selected:  true,
	}

	return r.db.Create(&cartItem).Error
}

// GetCartItems 获取用户所有购物车商品
func (r *CartRepository) GetCartItems(userID uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("user_id = ?", userID).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// UpdateCartItemStatus 更新购物车商品状态
func (r *CartRepository) UpdateCartItemStatus(id uint64, selected bool) error {
	return r.db.Model(&model.CartItem{}).Where("id = ?", id).Update("selected", selected).Error
}

// UpdateAllCartItemStatus 更新用户所有购物车商品状态
func (r *CartRepository) UpdateAllCartItemStatus(userID uint64, selected bool) error {
	return r.db.Model(&model.CartItem{}).Where("user_id = ?", userID).Update("selected", selected).Error
}

// DeleteCartItem 删除购物车商品
func (r *CartRepository) DeleteCartItem(id uint64) error {
	return r.db.Delete(&model.CartItem{}, "id = ?", id).Error
}

// GetCartItemsByIDs 获取购物车商品
func (r *CartRepository) GetCartItemsByIDs(ids []uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("id IN ?", ids).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// GetSelectedCartItems 获取用户选中的购物车商品
func (r *CartRepository) GetSelectedCartItems(userID uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("user_id = ? AND selected = ?", userID, true).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}
