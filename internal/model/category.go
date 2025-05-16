package model

import "time"

// Category represents a product category
type Category struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	ParentID  uint      `json:"parentID" gorm:"column:parent_id;index"`
	ImageUrl  string    `json:"imageUrl" gorm:"column:image_url"`
	SortOrder int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type CategoryTree struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	ImageUrl  string     `json:"imageUrl"`
	SortOrder int        `json:"sortOrder"`
	Children  []Category `json:"children"`
}
