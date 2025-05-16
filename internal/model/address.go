package model

import (
	"time"
)

// Address represents a shipping address
type Address struct {
	ID         uint64    `json:"id" gorm:"column:id;primaryKey"`
	UserID     uint64    `json:"user_id" gorm:"column:user_id;index"`
	Name       string    `json:"name" gorm:"column:name;not null"`
	Phone      string    `json:"phone" gorm:"column:phone;not null"`
	Province   string    `json:"province" gorm:"column:province;not null"`
	City       string    `json:"city" gorm:"column:city;not null"`
	District   string    `json:"district" gorm:"column:district;not null"`
	DetailAddr string    `json:"detailAddr" gorm:"column:detail_addr;not null"`
	IsDefault  bool      `json:"isDefault" gorm:"default:false"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at"`
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
