package model

import "time"

// Promotion represents a promotion activity
type Promotion struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Image     string    `json:"image" gorm:"not null"`
	Link      string    `json:"link"`
	SortOrder int       `json:"sort_order" gorm:"default:0"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
