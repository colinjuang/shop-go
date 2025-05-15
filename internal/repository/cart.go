package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
)

// CartRepository handles database operations for cart items
type CartRepository struct{}

// NewCartRepository creates a new cart repository
func NewCartRepository() *CartRepository {
	return &CartRepository{}
}

// AddToCart adds a product to the cart
func (r *CartRepository) AddToCart(userId, productId uint, quantity int) error {
	// Check if the product exists in the cart
	var existingItem model.CartItem
	result := DB.Where("user_id = ? AND product_id = ?", userId, productId).First(&existingItem)

	if result.Error == nil {
		// Update quantity if the product already exists
		existingItem.Quantity += quantity
		return DB.Save(&existingItem).Error
	}

	// Add new item if it doesn't exist
	cartItem := model.CartItem{
		UserId:    userId,
		ProductId: productId,
		Quantity:  quantity,
		Selected:  true,
	}

	return DB.Create(&cartItem).Error
}

// GetCartItems gets all cart items for a user
func (r *CartRepository) GetCartItems(userId uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := DB.Where("user_id = ?", userId).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// UpdateCartItemStatus updates the status of a cart item
func (r *CartRepository) UpdateCartItemStatus(id uint, selected bool) error {
	return DB.Model(&model.CartItem{}).Where("id = ?", id).Update("selected", selected).Error
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (r *CartRepository) UpdateAllCartItemStatus(userId uint, selected bool) error {
	return DB.Model(&model.CartItem{}).Where("user_id = ?", userId).Update("selected", selected).Error
}

// DeleteCartItem deletes a cart item
func (r *CartRepository) DeleteCartItem(id uint) error {
	return DB.Delete(&model.CartItem{}, id).Error
}

// GetCartItemsByIDs gets cart items by IDs
func (r *CartRepository) GetCartItemsByIDs(ids []uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := DB.Where("id IN ?", ids).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}

// GetSelectedCartItems gets selected cart items for a user
func (r *CartRepository) GetSelectedCartItems(userId uint) ([]model.CartItem, error) {
	var cartItems []model.CartItem
	result := DB.Where("user_id = ? AND selected = ?", userId, true).Preload("Product").Find(&cartItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return cartItems, nil
}
