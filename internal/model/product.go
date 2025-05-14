package model

import (
	"time"
)

// Product represents a product
type Product struct {
	ID             uint      `json:"id" gorm:"column:id;primaryKey"`
	Name           string    `json:"name" gorm:"column:name;not null"`
	FloralLanguage string    `json:"floralLanguage" gorm:"column:floral_language"`
	Price          float64   `json:"price" gorm:"column:price;type:decimal(10,2)"`
	MarketPrice    float64   `json:"marketPrice" gorm:"column:market_price;type:decimal(10,2)"`
	SaleCount      int       `json:"saleCount" gorm:"column:sale_count;default:0"`
	StockCount     int       `json:"stockCount" gorm:"column:stock_count;default:0"`
	CategoryID     uint      `json:"categoryId" gorm:"column:category_id;index"`
	SubCategoryID  uint      `json:"subCategoryId" gorm:"column:sub_category_id;index"`
	Material       string    `json:"material" gorm:"column:material"`
	Packing        string    `json:"packing" gorm:"column:packing"`
	ImageUrl       string    `json:"imageUrl" gorm:"column:image_url"`
	Status         int       `json:"status" gorm:"column:status;default:1"` // 1: on sale, 0: off sale
	Recommend      bool      `json:"recommend" gorm:"column:recommend;default:false"`
	SortOrder      int       `json:"sortOrder" gorm:"column:sort_order;default:0"`
	ApplyUser      string    `json:"applyUser" gorm:"column:apply_user"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
