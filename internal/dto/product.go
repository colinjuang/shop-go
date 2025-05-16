package dto

import "time"

type ProductRequest struct {
	ID             uint64  `json:"id"`
	Name           string  `json:"name"`
	FloralLanguage string  `json:"floralLanguage"`
	Price          float64 `json:"price"`
	MarketPrice    float64 `json:"marketPrice"`
	SaleCount      int     `json:"saleCount"`
	StockCount     int     `json:"stockCount"`
	CategoryID     uint64  `json:"categoryID"`
	SubCategoryID  uint64  `json:"subCategoryID"`
	Material       string  `json:"material"`
	Packing        string  `json:"packing"`
	ImageUrl       string  `json:"imageUrl"`
	Status         int     `json:"status"` // 1: on sale, 0: off sale
	Recommend      bool    `json:"recommend"`
	SortOrder      int     `json:"sortOrder"`
	ApplyUser      string  `json:"applyUser"`
}

type ProductResponse struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name"`
	FloralLanguage string    `json:"floralLanguage"`
	Price          float64   `json:"price"`
	MarketPrice    float64   `json:"marketPrice"`
	SaleCount      int       `json:"saleCount"`
	StockCount     int       `json:"stockCount"`
	CategoryID     uint64    `json:"categoryID"`
	SubCategoryID  uint64    `json:"subCategoryID"`
	Material       string    `json:"material"`
	Packing        string    `json:"packing"`
	ImageUrl       string    `json:"imageUrl"`
	Status         int       `json:"status"` // 1: on sale, 0: off sale
	Recommend      bool      `json:"recommend"`
	SortOrder      int       `json:"sortOrder"`
	ApplyUser      string    `json:"applyUser"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
