package model

import (
	"time"
)

// CartItem represents an item in the cart
type CartItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"index;not null"`
	ProductID uint      `json:"product_id" gorm:"index;not null"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	Selected  bool      `json:"selected" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
}

// CartItemResponse is the response for cart item list
type CartItemResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Selected  bool    `json:"selected"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	MainImage string  `json:"main_image"`
	Stock     int     `json:"stock"`
}
