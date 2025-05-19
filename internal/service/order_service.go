package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"

	"github.com/colinjuang/shop-go/internal/api/middleware"
	"github.com/colinjuang/shop-go/internal/api/response"
	"github.com/colinjuang/shop-go/internal/constant"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/redis"
	"github.com/colinjuang/shop-go/internal/repository"
)

// OrderService handles business logic for orders
type OrderService struct {
	orderRepo    *repository.OrderRepository
	cartRepo     *repository.CartRepository
	productRepo  *repository.ProductRepository
	addressRepo  *repository.AddressRepository
	cacheService *redis.CacheService
}

// NewOrderService creates a new order service
func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo:    repository.NewOrderRepository(),
		cartRepo:     repository.NewCartRepository(),
		productRepo:  repository.NewProductRepository(),
		addressRepo:  repository.NewAddressRepository(),
		cacheService: redis.NewCacheService(),
	}
}

// GetOrderDetail gets order details for checkout
func (s *OrderService) GetOrderDetail(userID uint64, cartIDs []uint64, productID *uint64, quantity *int) (float64, []model.CartItem, error) {
	var cartItems []model.CartItem
	var err error

	// Get items from cart or direct purchase
	if len(cartIDs) > 0 {
		// Cart checkout
		items, err := s.cartRepo.GetCartItemsByIDs(cartIDs)
		if err != nil {
			return 0, nil, err
		}

		// Verify ownership of cart items
		for _, item := range items {
			if item.UserID != userID {
				return 0, nil, pkgerrors.ErrPaymentFailed
			}
		}

		cartItems = items
	} else if productID != nil && quantity != nil {
		// Direct purchase
		product, err := s.productRepo.GetProductByID(*productID)
		if err != nil {
			return 0, nil, err
		}

		// Check stock
		if product.StockCount < *quantity {
			return 0, nil, pkgerrors.ErrOutOfStock
		}

		cartItems = []model.CartItem{
			{
				UserID:    userID,
				ProductID: *productID,
				Quantity:  *quantity,
				Product:   *product,
			},
		}
	} else {
		return 0, nil, errors.New("invalid request")
	}

	// Calculate total amount
	var totalAmount float64
	for _, item := range cartItems {
		totalAmount += item.Product.Price * float64(item.Quantity)
	}

	return totalAmount, cartItems, err
}

// CreateOrder creates a new order
func (s *OrderService) CreateOrder(reqUser *middleware.UserClaim, req model.OrderRequest) (*model.Order, error) {
	ctx := context.Background()

	// Create a lock key for this order creation
	// This prevents race conditions when multiple requests try to create an order
	// for the same user with the same products
	lockKey := fmt.Sprintf("order:create:user:%d", reqUser.UserID)
	if len(req.CartIDs) > 0 {
		for _, id := range req.CartIDs {
			lockKey += fmt.Sprintf(":cart:%d", id)
		}
	} else {
		lockKey += fmt.Sprintf(":product:%d:qty:%d", req.ProductID, req.Quantity)
	}

	// Use Redis distributed lock to prevent race conditions
	// Lock will expire after 30 seconds as a safety measure
	var order *model.Order
	var orderErr error

	err := redis.WithLock(ctx, lockKey, 30*time.Second, func() error {
		// Verify address
		address, err := s.addressRepo.GetAddressByID(req.AddressID)
		if err != nil {
			return err
		}

		if address.UserID != reqUser.UserID {
			return pkgerrors.ErrPaymentFailed
		}

		// Get order details
		totalAmount, cartItems, err := s.GetOrderDetail(reqUser.UserID, req.CartIDs, &req.ProductID, &req.Quantity)
		if err != nil {
			return err
		}

		// Create order
		order = &model.Order{
			UserID:        reqUser.UserID,
			TotalAmount:   totalAmount,
			PaymentAmount: totalAmount, // No discount for now
			Status:        model.OrderStatusPending,
			AddressID:     req.AddressID,
			ReceiverName:  address.Name,
			ReceiverPhone: address.Phone,
			Address:       address.Province + address.City + address.District + address.DetailAddr,
			PaymentType:   1, // Default to WeChat
		}

		// Create order items
		var orderItems []model.OrderItem
		for _, item := range cartItems {
			// Check stock again within the lock to prevent race conditions
			product, err := s.productRepo.GetProductByID(item.ProductID)
			if err != nil {
				return err
			}

			if product.StockCount < item.Quantity {
				return pkgerrors.ErrOutOfStock
			}

			orderItem := model.OrderItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Product.Price,
				Name:      item.Product.Name,
				ImageUrl:  item.Product.ImageUrl,
			}
			orderItems = append(orderItems, orderItem)
		}

		order.OrderItems = orderItems

		// Save order
		if err := s.orderRepo.CreateOrder(order); err != nil {
			return err
		}

		// Update product stock
		for _, item := range orderItems {
			fmt.Printf("Would decrease stock for product %d by %d\n", item.ProductID, item.Quantity)
		}

		// Delete cart items if from cart
		if len(req.CartIDs) > 0 {
			for _, id := range req.CartIDs {
				s.cartRepo.DeleteCartItem(id)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Cache the order for future retrievals
	cacheKey := fmt.Sprintf(constant.OrderPrefix+":%d", order.ID)
	_ = s.cacheService.Set(ctx, cacheKey, *order, 30*time.Minute)

	return order, orderErr
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

	return orderPtr, nil
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
