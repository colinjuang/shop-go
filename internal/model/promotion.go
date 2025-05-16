package model

import "time"

// Promotion represents a promotion activity
type Promotion struct {
	ID            uint      `json:"id" gorm:"column:id;primaryKey"`
	Title         string    `json:"title" gorm:"column:title;not null"`
	ImageUrl      string    `json:"imageUrl" gorm:"column:image_url;not null"`
	SubCategoryID uint      `json:"subCategoryID" gorm:"column:sub_category_id;not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
