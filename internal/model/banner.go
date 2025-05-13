package model

import (
	"time"
)

// Banner represents a banner for homepage carousel
type Banner struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	ImageUrl  string    `json:"image_url" gorm:"not null"`
	LinkUrl   string    `json:"link_url"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}
