package model

import (
	"time"
)

// Product represents a product
type Product struct {
	ID          uint      `json:"id" gorm:"column:id;primaryKey"`
	Name        string    `json:"name" gorm:"column:name;not null"`
	Description string    `json:"description" gorm:"column:description"`
	Price       float64   `json:"price" gorm:"column:price;type:decimal(10,2)"`
	Stock       int       `json:"stock" gorm:"column:stock;default:0"`
	CategoryID  uint      `json:"categoryId" gorm:"column:category_id;index"`
	Images      string    `json:"images" gorm:"column:images"`
	MainImage   string    `json:"mainImage" gorm:"column:main_image"`
	Status      int       `json:"status" gorm:"column:status;default:1"` // 1: on sale, 0: off sale
	Hot         bool      `json:"hot" gorm:"column:hot;default:false"`
	Recommend   bool      `json:"recommend" gorm:"column:recommend;default:false"`
	SortOrder   int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
