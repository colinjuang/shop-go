package model

import (
	"time"
)

const (
	// OrderStatusPending is the status for pending payment
	OrderStatusPending = 0
	// OrderStatusPaid is the status for paid orders
	OrderStatusPaid = 1
	// OrderStatusShipped is the status for shipped orders
	OrderStatusShipped = 2
	// OrderStatusCompleted is the status for completed orders
	OrderStatusCompleted = 3
	// OrderStatusCancelled is the status for cancelled orders
	OrderStatusCancelled = 4
)

// Order represents an order
type Order struct {
	ID            uint        `json:"id" gorm:"primaryKey"`
	UserID        uint        `json:"user_id" gorm:"index;not null"`
	OrderNo       string      `json:"order_no" gorm:"uniqueIndex;not null"`
	TotalAmount   float64     `json:"total_amount" gorm:"type:decimal(10,2);not null"`
	PaymentAmount float64     `json:"payment_amount" gorm:"type:decimal(10,2);not null"`
	Status        int         `json:"status" gorm:"default:0"`
	PaymentTime   time.Time   `json:"payment_time"`
	AddressID     uint        `json:"address_id"`
	ReceiverName  string      `json:"receiver_name"`
	ReceiverPhone string      `json:"receiver_phone"`
	Address       string      `json:"address"`
	PaymentType   int         `json:"payment_type" gorm:"default:1"` // 1: wechat
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"index;not null"`
	ProductID uint      `json:"product_id" gorm:"index;not null"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Name      string    `json:"name" gorm:"not null"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// OrderRequest represents the order creation request
type OrderRequest struct {
	AddressID uint   `json:"address_id" binding:"required"`
	CartIDs   []uint `json:"cart_ids"`                 // Optional for cart checkout
	ProductID uint   `json:"product_id"`               // For direct purchase
	Quantity  int    `json:"quantity" binding:"min=1"` // For direct purchase
}

// PaymentResponse represents the payment response
type PaymentResponse struct {
	PaymentID string `json:"payment_id"`
	AppID     string `json:"app_id"`
	TimeStamp string `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}
