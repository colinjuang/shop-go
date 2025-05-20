package repository

import (
	"time"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// OrderRepository 订单仓库
type OrderRepository struct {
	db *gorm.DB
}

// NewOrderRepository
func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		db: database.GetDB(),
	}
}

// CreateOrder 创建新订单
func (r *OrderRepository) CreateOrder(order *model.Order) error {
	// 生成订单号
	// now := time.Now()
	// order.OrderNo = fmt.Sprintf("%s%d", now.Format("20060102150405"), order.UserID)

	return r.db.Create(order).Error
}

func (r *OrderRepository) GetOrderByID(id uint64) (*model.Order, error) {
	var order model.Order
	result := r.db.First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// GetOrderByID 获取订单
func (r *OrderRepository) GetOrderAndOrderItemByID(id uint64) (*model.OrderWithOrderItem, error) {
	var order model.OrderWithOrderItem
	result := r.db.Preload("OrderItem").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (r *OrderRepository) GetOrderByIDAndUserID(id uint64, userID uint64) (*model.Order, error) {
	var order model.Order
	result := r.db.Where("id = ? AND user_id = ?", id, userID).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// GetOrderByOrderNo 获取订单
func (r *OrderRepository) GetOrderByOrderNo(orderNo string) (*model.Order, error) {
	var order model.Order
	result := r.db.Where("order_no = ?", orderNo).Preload("OrderItem").First(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// UpdateOrderStatus 更新订单状态
func (r *OrderRepository) UpdateOrderStatus(id uint64, status int) error {
	updates := map[string]interface{}{
		"status": status,
	}

	// 如果已支付，则添加支付时间
	if status == model.OrderStatusPaid {
		updates["payment_time"] = time.Now()
	}

	return r.db.Model(&model.Order{}).Where("id = ?", id).Updates(updates).Error
}

// GetOrdersByUserID 获取用户订单
func (r *OrderRepository) GetOrdersByUserID(userID uint64, page, pageSize int, status *int) ([]model.Order, int64, error) {
	var orders []model.Order
	var count int64

	query := r.db.Model(&model.Order{}).Where("user_id = ?", userID)

	// 如果提供状态，则应用状态过滤
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	// 获取总数
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页结果
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Preload("OrderItem").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, count, nil
}
