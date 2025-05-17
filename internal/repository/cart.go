package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// CartRepository handles database operations for cart items
type CartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new cart repository
func NewCartRepository() *CartRepository {
	return &CartRepository{
		db: database.GetDB(),
	}
}

// AddToCart adds a product to the cart
func (r *CartRepository) AddToCart(userID uint64, productID uint64, quantity int) error {
	// Check if the product exists in the cart
	var existingItem model.CartItem
	result := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&existingItem)

	if result.Error == nil {
		// Update quantity if the product already exists
		existingItem.Quantity += quantity
		return r.db.Save(&existingItem).Error
	}

	// Add new item if it doesn't exist
	cartItem := model.CartItem{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
		Selected:  true,
	}

	return r.db.Create(&cartItem).Error
}

// GetCartItems gets all cart items for a user
func (r *CartRepository) GetCartItems(userID uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("user_id = ?", userID).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// UpdateCartItemStatus updates the status of a cart item
func (r *CartRepository) UpdateCartItemStatus(id uint64, selected bool) error {
	return r.db.Model(&model.CartItem{}).Where("id = ?", id).Update("selected", selected).Error
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (r *CartRepository) UpdateAllCartItemStatus(userID uint64, selected bool) error {
	return r.db.Model(&model.CartItem{}).Where("user_id = ?", userID).Update("selected", selected).Error
}

// DeleteCartItem deletes a cart item
func (r *CartRepository) DeleteCartItem(id uint64) error {
	return r.db.Delete(&model.CartItem{}, "id = ?", id).Error
}

// GetCartItemsByIDs gets cart items by IDs
func (r *CartRepository) GetCartItemsByIDs(ids []uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("id IN ?", ids).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// GetSelectedCartItems gets selected cart items for a user
func (r *CartRepository) GetSelectedCartItems(userID uint64) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := r.db.Where("user_id = ? AND selected = ?", userID, true).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}
