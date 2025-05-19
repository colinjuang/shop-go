package model

import "time"

// Category represents a product category
type Category struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	ParentID  uint64    `json:"parentID" gorm:"column:parent_id;index"`
	ImageUrl  string    `json:"imageUrl" gorm:"column:image_url"`
	SortOrder int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
