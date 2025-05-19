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
	ID            uint64    `json:"id" gorm:"column:id;primaryKey"`
	UserID        uint64    `json:"userID" gorm:"column:user_id;index;not null"`
	OrderNo       string    `json:"orderNo" gorm:"column:order_no;uniqueIndex;not null"`
	TotalAmount   float64   `json:"totalAmount" gorm:"column:total_amount;type:decimal(10,2);not null"`     // 总金额
	PaymentAmount float64   `json:"paymentAmount" gorm:"column:payment_amount;type:decimal(10,2);not null"` // 支付金额
	Status        int       `json:"status" gorm:"column:status;default:0"`
	PaymentTime   time.Time `json:"paymentTime" gorm:"column:payment_time"`
	AddressID     uint64    `json:"addressID" gorm:"column:address_id"`
	ReceiverName  string    `json:"receiverName" gorm:"column:receiver_name"`
	ReceiverPhone string    `json:"receiverPhone" gorm:"column:receiver_phone"`
	Address       string    `json:"address" gorm:"column:address"`
	PaymentType   int       `json:"paymentType" gorm:"default:1"` // 1: wechat
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type OrderWithOrderItem struct {
	Order
	OrderItem []OrderItem `json:"orderItem" gorm:"foreignKey:OrderID"`
}
