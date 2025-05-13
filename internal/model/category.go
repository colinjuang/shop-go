package model

import "time"

// Category represents a product category
type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	ParentID  uint      `json:"parent_id" gorm:"index"`
	ImageUrl  string    `json:"image_url"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
