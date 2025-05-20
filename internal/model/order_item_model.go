package model

import "time"

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	OrderID   uint64    `json:"orderID" gorm:"column:order_id;index;not null"`
	ProductID uint64    `json:"productID" gorm:"column:product_id;index;not null"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;not null"`
	Price     float64   `json:"price" gorm:"column:price;type:decimal(10,2);not null"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	ImageUrl  string    `json:"image" gorm:"column:image_url"`
	Blessing  string    `json:"blessing" gorm:"column:blessing"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
