package model

import (
	"time"
)

type Cart struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	UserID    uint64    `json:"userID" gorm:"column:user_id;index;not null"`
	ProductID uint64    `json:"productID" gorm:"column:product_id;index;not null"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	Selected  bool      `json:"selected" gorm:"default:true"`
	Blessing  string    `json:"blessing" gorm:"column:blessing"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	Product   Product   `json:"product"`
}
