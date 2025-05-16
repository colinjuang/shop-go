package dto

import "time"

type PromotionResponse struct {
	ID            uint64    `json:"id"`
	Title         string    `json:"title"`
	ImageUrl      string    `json:"imageUrl"`
	SubCategoryID uint64    `json:"subCategoryId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
