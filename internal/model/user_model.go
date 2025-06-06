package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	Username  string    `json:"username" gorm:"column:username;uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"column:password;not null"`
	OpenID    string    `json:"openid" gorm:"column:openid;uniqueIndex;not null"`
	Nickname  string    `json:"nickname" gorm:"column:nickname"`
	Avatar    string    `json:"avatar" gorm:"column:avatar"`
	Gender    int       `json:"gender" gorm:"column:gender"`
	City      string    `json:"city" gorm:"column:city"`
	Province  string    `json:"province" gorm:"column:province"`
	District  string    `json:"district" gorm:"column:district"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
