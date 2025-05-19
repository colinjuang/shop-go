package model

import (
	"time"
)

// Address represents a shipping address
type Address struct {
	ID           uint64    `json:"id" gorm:"column:id;primaryKey"`                     // 地址ID
	UserID       uint64    `json:"user_id" gorm:"column:user_id;index"`                // 用户ID
	Name         string    `json:"name" gorm:"column:name;not null"`                   // 收货人
	Phone        string    `json:"phone" gorm:"column:phone;not null"`                 // 手机号码
	Province     string    `json:"province" gorm:"column:province;not null"`           // 省
	ProvinceCode string    `json:"province_code" gorm:"column:province_code;not null"` // 省编码
	City         string    `json:"city" gorm:"column:city;not null"`                   // 市
	CityCode     string    `json:"city_code" gorm:"column:city_code;not null"`         // 市编码
	District     string    `json:"district" gorm:"column:district;not null"`           // 区
	DistrictCode string    `json:"district_code" gorm:"column:district_code;not null"` // 区编码
	DetailAddr   string    `json:"detail_addr" gorm:"column:detail_addr;not null"`     // 详细地址
	IsDefault    int8      `json:"is_default" gorm:"default:0"`                        // 是否默认地址
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at"`
}
