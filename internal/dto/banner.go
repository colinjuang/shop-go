package dto

import "time"

type BannerRequest struct {
	ID uint64 `json:"id"`
}

type BannerResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	ImageUrl  string    `json:"imageUrl"`
	ProductID uint64    `json:"productId"`
	SortOrder int       `json:"sortOrder"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
