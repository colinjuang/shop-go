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
	ID            uint64      `json:"id" gorm:"column:id;primaryKey"`
	UserID        uint64      `json:"userID" gorm:"column:user_id;index;not null"`
	OrderNo       string      `json:"orderNo" gorm:"column:order_no;uniqueIndex;not null"`
	TotalAmount   float64     `json:"totalAmount" gorm:"column:total_amount;type:decimal(10,2);not null"`
	PaymentAmount float64     `json:"paymentAmount" gorm:"column:payment_amount;type:decimal(10,2);not null"`
	Status        int         `json:"status" gorm:"column:status;default:0"`
	PaymentTime   time.Time   `json:"paymentTime" gorm:"column:payment_time"`
	AddressID     uint64      `json:"addressID" gorm:"column:address_id"`
	ReceiverName  string      `json:"receiverName" gorm:"column:receiver_name"`
	ReceiverPhone string      `json:"receiverPhone" gorm:"column:receiver_phone"`
	Address       string      `json:"address" gorm:"column:address"`
	PaymentType   int         `json:"paymentType" gorm:"default:1"` // 1: wechat
	CreatedAt     time.Time   `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time   `json:"updatedAt" gorm:"column:updated_at"`
	OrderItems    []OrderItem `json:"orderItems" gorm:"foreignKey:OrderID"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint64    `json:"id" gorm:"column:id;primaryKey"`
	OrderID   uint64    `json:"orderID" gorm:"column:order_id;index;not null"`
	ProductID uint64    `json:"productID" gorm:"column:product_id;index;not null"`
	Quantity  int       `json:"quantity" gorm:"column:quantity;not null"`
	Price     float64   `json:"price" gorm:"column:price;type:decimal(10,2);not null"`
	Name      string    `json:"name" gorm:"column:name;not null"`
	ImageUrl  string    `json:"image" gorm:"column:image_url"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

// OrderRequest represents the order creation request
type OrderRequest struct {
	AddressID uint64   `json:"address_id" binding:"required"`
	CartIDs   []uint64 `json:"cart_ids"`                 // Optional for cart checkout
	ProductID uint64   `json:"product_id"`               // For direct purchase
	Quantity  int      `json:"quantity" binding:"min=1"` // For direct purchase
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
