package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" gorm:"column:id;primaryKey"`
	Username  string    `json:"username" gorm:"column:username;uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"column:password;not null"`
	OpenID    string    `json:"openid" gorm:"column:openid;uniqueIndex;not null"`
	Nickname  string    `json:"nickname" gorm:"column:nickname"`
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Gender    int       `json:"gender" gorm:"column:gender"`
	City      string    `json:"city" gorm:"column:city"`
	Province  string    `json:"province" gorm:"column:province"`
	Country   string    `json:"country" gorm:"column:country"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
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
