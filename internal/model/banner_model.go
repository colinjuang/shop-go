package model

import (
	"time"
)

// Banner represents a banner for homepage carousel
type Banner struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	Title     string    `json:"title" gorm:"column:title;not null"`
	ImageUrl  string    `json:"imageUrl" gorm:"column:image_url;not null"`
	ProductID uint64    `json:"productID" gorm:"column:product_id"`
	SortOrder int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:datetime"`
}
