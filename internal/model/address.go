package model

import (
	"time"
)

// Address represents a shipping address
type Address struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"index"`
	Name       string    `json:"name" binding:"required"`
	Phone      string    `json:"phone" binding:"required"`
	Province   string    `json:"province" binding:"required"`
	City       string    `json:"city" binding:"required"`
	District   string    `json:"district" binding:"required"`
	DetailAddr string    `json:"detail_addr" binding:"required"`
	IsDefault  bool      `json:"is_default" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// AddressRequest represents the address creation/update request
type AddressRequest struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Province   string `json:"province" binding:"required"`
	City       string `json:"city" binding:"required"`
	District   string `json:"district" binding:"required"`
	DetailAddr string `json:"detail_addr" binding:"required"`
	IsDefault  bool   `json:"is_default"`
}
