package service

import (
	"context"
	"fmt"
	"time"

	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"
	"github.com/colinjuang/shop-go/internal/request"

	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/colinjuang/shop-go/internal/response"
)

// OrderService handles business logic for orders
type OrderService struct {
	orderRepo     *repository.OrderRepository
	orderItemRepo *repository.OrderItemRepository
	cartRepo      *repository.CartRepository
	productRepo   *repository.ProductRepository
	addressRepo   *repository.AddressRepository
	cacheService  *redis.CacheService
}

// NewOrderService creates a new order service
func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo:     repository.NewOrderRepository(),
		orderItemRepo: repository.NewOrderItemRepository(),
		cartRepo:      repository.NewCartRepository(),
		productRepo:   repository.NewProductRepository(),
		addressRepo:   repository.NewAddressRepository(),
		cacheService:  redis.NewCacheService(),
	}
}

// GetOrderDetail gets order details for checkout
func (s *OrderService) GetOrderDetail(userID uint64, orderID uint64) (*response.OrderDetailResponse, error) {
	order, err := s.orderRepo.GetOrderByIDAndUserID(orderID, userID)
	if err != nil {
		return nil, err
	}

	orderItems, err := s.orderItemRepo.GetOrderItemsByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	address, err := s.addressRepo.GetAddressByID(order.AddressID)
	if err != nil {
		return nil, err
	}

	orderItemsResponse := make([]response.OrderItemResponse, len(orderItems))
	for i, item := range orderItems {
		orderItemsResponse[i] = response.OrderItemResponse{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Name:      item.Name,
			ImageUrl:  item.ImageUrl,
		}
	}

	orderDetail := &response.OrderDetailResponse{
		OrderID:     order.ID,
		OrderNo:     order.OrderNo,
		TotalAmount: order.TotalAmount,
		OrderItem:   orderItemsResponse,
		Address: response.AddressResponse{
			ID:           address.ID,
			Phone:        address.Phone,
			Name:         address.Name,
			City:         address.City,
			CityCode:     address.CityCode,
			Province:     address.Province,
			ProvinceCode: address.ProvinceCode,
			District:     address.District,
			DistrictCode: address.DistrictCode,
			DetailAddr:   address.DetailAddr,
			FullAddr:     address.Province + address.City + address.District + address.DetailAddr,
			IsDefault:    address.IsDefault,
		},
	}

	return orderDetail, nil
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(userID uint64, req request.CreateOrderRequest) error {
	address, err := s.addressRepo.GetAddressByID(req.AddressID)
	if err != nil {
		return err
	}

	if address.UserID != userID {
		return pkgerrors.ErrAddressNotFound
	}

	order := &model.OrderWithOrderItem{
		Order: model.Order{
			UserID:        userID,                                                                  // 用户ID
			TotalAmount:   0,                                                                       // 总金额
			PaymentAmount: 0,                                                                       // 支付金额
			Status:        model.OrderStatusPending,                                                // 订单状态
			AddressID:     req.AddressID,                                                           // 地址ID
			ReceiverName:  address.Name,                                                            // 收货人姓名
			ReceiverPhone: address.Phone,                                                           // 收货人电话
			Address:       address.Province + address.City + address.District + address.DetailAddr, // 地址
			PaymentType:   constant.PaymentMethodWechat,                                            // 默认微信支付
		},
		OrderItem: []model.OrderItem{},
	}

	var totalAmount float64
	var paymentAmount float64
	var orderItems []model.OrderItem
	for _, cartID := range req.CartIDs {
		cart, err := s.cartRepo.GetCartByID(cartID)
		if err != nil {
			return err
		}

		if cart.Product.StockCount < cart.Quantity {
			return pkgerrors.ErrOutOfStock
		}

		orderItem := model.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
			Name:      cart.Product.Name,
			ImageUrl:  cart.Product.ImageUrl,
		}
		orderItems = append(orderItems, orderItem)

		totalAmount += float64(cart.Quantity) * cart.Product.Price
		paymentAmount += float64(cart.Quantity) * cart.Product.Price
	}

	order.OrderItem = orderItems
	order.TotalAmount = totalAmount
	order.PaymentAmount = paymentAmount

	// 保存订单
	if err := s.orderRepo.CreateOrder(&order.Order); err != nil {
		return err
	}

	// 更新商品库存
	for _, item := range orderItems {
		s.productRepo.UpdateProductStock(item.ProductID, item.Quantity)
	}

	// 删除购物车
	if len(req.CartIDs) > 0 {
		for _, id := range req.CartIDs {
			s.cartRepo.DeleteCart(id)
		}
	}

	return nil
}

// GetOrderByID gets an order by ID
func (s *OrderService) GetOrderByID(id uint64, userID uint64) (*model.Order, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf(constant.OrderPrefix+":%d", id)

	// Try to get from cache
	var order model.Order
	err := s.cacheService.GetObject(ctx, cacheKey, &order)
	if err == nil && order.UserID == userID {
		return &order, nil
	}

	// If not in cache or not owned by user, get from database
	orderPtr, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the order belongs to the user
	if orderPtr.UserID != userID {
		return nil, pkgerrors.ErrPaymentFailed
	}

	// Cache for future requests
	_ = s.cacheService.Set(ctx, cacheKey, *orderPtr, 30*time.Minute)

	return &order, nil
}

func (s *OrderService) GetOrderAndOrderItemByID(id uint64, userID uint64) (*model.OrderWithOrderItem, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf(constant.OrderPrefix+":%d", id)

	// Try to get from cache
	var order model.OrderWithOrderItem
	err := s.cacheService.GetObject(ctx, cacheKey, &order)
	if err == nil && order.UserID == userID {
		return &order, nil
	}

	// If not in cache or not owned by user, get from database
	orderPtr, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the order belongs to the user
	if orderPtr.UserID != userID {
		return nil, pkgerrors.ErrPaymentFailed
	}

	// Cache for future requests
	_ = s.cacheService.Set(ctx, cacheKey, *orderPtr, 30*time.Minute)

	return &order, nil
}

// GetOrderByOrderNo gets an order by order number
func (s *OrderService) GetOrderByOrderNo(orderNo string, userID uint64) (*model.Order, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf(constant.OrderNo+":%s", orderNo)

	// Try to get from cache
	var order model.Order
	err := s.cacheService.GetObject(ctx, cacheKey, &order)
	if err == nil && order.UserID == userID {
		return &order, nil
	}

	// If not in cache or not owned by user, get from database
	orderPtr, err := s.orderRepo.GetOrderByOrderNo(orderNo)
	if err != nil {
		return nil, err
	}

	// Ensure the order belongs to the user
	if orderPtr.UserID != userID {
		return nil, pkgerrors.ErrPaymentFailed
	}

	// Cache for future requests
	_ = s.cacheService.Set(ctx, cacheKey, *orderPtr, 30*time.Minute)

	return orderPtr, nil
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(id uint64, userID uint64, status int) error {
	ctx := context.Background()

	order, err := s.GetOrderByID(id, userID)
	if err != nil {
		return err
	}

	// Use a lock to prevent race conditions when updating order status
	lockKey := fmt.Sprintf(constant.OrderStatus+":%d", id)

	err = redis.WithLock(ctx, lockKey, 30*time.Second, func() error {
		// Update status in database
		err := s.orderRepo.UpdateOrderStatus(order.ID, status)
		if err != nil {
			return err
		}

		// Invalidate cache
		cacheKey := fmt.Sprintf(constant.OrderPrefix+":%d", id)
		s.cacheService.Delete(ctx, cacheKey)

		return nil
	})

	return err
}

// GetOrdersByUserID gets orders for a user with pagination
func (s *OrderService) GetOrdersByUserID(userID uint64, page, pageSize int, status *int) (*response.Pagination, error) {
	ctx := context.Background()

	// Generate cache key
	cacheKey := fmt.Sprintf(constant.OrderUserPage+"%d:page:%d:size:%d", userID, page, pageSize)
	if status != nil {
		cacheKey += fmt.Sprintf(":status:%d", *status)
	}

	// Try to get from cache
	var pagination response.Pagination
	err := s.cacheService.GetObject(ctx, cacheKey, &pagination)
	if err == nil {
		return &pagination, nil
	}

	// If not in cache, get from database
	orders, total, err := s.orderRepo.GetOrdersByUserID(userID, page, pageSize, status)
	if err != nil {
		return nil, err
	}

	pagination = response.NewPagination(total, page, pageSize, orders)

	// Cache for 5 minutes (shorter time since order list changes more frequently)
	_ = s.cacheService.Set(ctx, cacheKey, pagination, 5*time.Minute)

	return &pagination, nil
}
