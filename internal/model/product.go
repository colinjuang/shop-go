package model

import (
	"time"
)

// Product represents a product
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2)"`
	Stock       int       `json:"stock" gorm:"default:0"`
	CategoryID  uint      `json:"category_id" gorm:"index"`
	Images      string    `json:"images"`
	MainImage   string    `json:"main_image"`
	Status      int       `json:"status" gorm:"default:1"` // 1: on sale, 0: off sale
	Hot         bool      `json:"hot" gorm:"default:false"`
	Recommend   bool      `json:"recommend" gorm:"default:false"`
	SortOrder   int       `json:"sort_order" gorm:"default:0"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
