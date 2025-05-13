package model

import "time"

// Promotion represents a promotion activity
type Promotion struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Title     string    `json:"title" gorm:"column:title;not null"`
	Image     string    `json:"image" gorm:"column:image;not null"`
	Link      string    `json:"link" gorm:"column:link"`
	SortOrder int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
