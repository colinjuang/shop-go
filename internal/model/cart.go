package model

import (
	"time"
)

// CartItem represents an item in the cart
type CartItem struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	UserId    uint      `json:"userId" gorm:"column:user_id;index;not null"`
	ProductId uint      `json:"productId" gorm:"column:product_id;index;not null"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	Selected  bool      `json:"selected" gorm:"default:true"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductId"`
}

// CartItemResponse is the response for cart item list
type CartItemResponse struct {
	ID         uint    `json:"id"`
	ProductId  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Selected   bool    `json:"selected"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ImageUrl   string  `json:"image_url"`
	StockCount int     `json:"stock_count"`
}
