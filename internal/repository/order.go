package repository

import (
	"fmt"
	"time"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// OrderRepository handles database operations for orders
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new order repository
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: database.GetDB(),
	}
}

// CreateOrder creates a new order
func (r *OrderRepository) CreateOrder(order *model.Order) error {
	// Generate order number
	now := time.Now()
	order.OrderNo = fmt.Sprintf("%s%d", now.Format("20060102150405"), order.UserID)

	return r.db.Create(order).Error
}

// GetOrderByID gets an order by ID
func (r *OrderRepository) GetOrderByID(id uint64) (*model.Order, error) {
	var order model.Order
	result := r.db.Preload("OrderItems").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// GetOrderByOrderNo gets an order by order number
func (r *OrderRepository) GetOrderByOrderNo(orderNo string) (*model.Order, error) {
	var order model.Order
	result := r.db.Where("order_no = ?", orderNo).Preload("OrderItems").First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// UpdateOrderStatus updates the status of an order
func (r *OrderRepository) UpdateOrderStatus(id uint64, status int) error {
	updates := map[string]interface{}{
		"status": status,
	}

	// Add payment time if paid
	if status == model.OrderStatusPaid {
		updates["payment_time"] = time.Now()
	}

	return r.db.Model(&model.Order{}).Where("id = ?", id).Updates(updates).Error
}

// GetOrdersByUserID gets orders for a user with pagination
func (r *OrderRepository) GetOrdersByUserID(userID uint64, page, pageSize int, status *int) ([]model.Order, int64, error) {
	var orders []model.Order
	var count int64

	query := r.db.Model(&model.Order{}).Where("user_id = ?", userID)

	// Apply status filter if provided
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// Get total count
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Preload("OrderItems").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}
