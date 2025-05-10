package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OpenID    string    `json:"openid" gorm:"uniqueIndex;not null"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Gender    int       `json:"gender"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserUpdateInfo represents data used to update user information
type UserUpdateInfo struct {
	Nickname string `json:"nickname" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
	Gender   int    `json:"gender" binding:"required"`
	City     string `json:"city"`
	Province string `json:"province"`
	Country  string `json:"country"`
}
