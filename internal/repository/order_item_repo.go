package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{
		db: db,
	}
}

func (r *OrderItemRepository) CreateOrderItem(orderItem []model.OrderItem) error {
	return r.db.Create(orderItem).Error
}

func (r *OrderItemRepository) GetOrderItemsByOrderID(orderID uint64) ([]model.OrderItem, error) {
	var orderItems []model.OrderItem
	result := r.db.Where("order_id = ?", orderID).Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}
