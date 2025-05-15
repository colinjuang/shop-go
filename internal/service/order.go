package service

import (
	"context"
	"errors"
	"fmt"
	"time"

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
func (s *OrderService) GetOrderDetail(userId uint, cartIDs []uint, productId *uint, quantity *int) (float64, []model.CartItem, error) {
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
			if item.UserId != userId {
				return 0, nil, ErrorUnauthorized
			}
		}

		cartItems = items
	} else if productId != nil && quantity != nil {
		// Direct purchase
		product, err := s.productRepo.GetProductByID(*productId)
		if err != nil {
			return 0, nil, err
		}

		// Check stock
		if product.StockCount < *quantity {
			return 0, nil, ErrorOutOfStock
		}

		cartItems = []model.CartItem{
			{
				UserId:    userId,
				ProductId: *productId,
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
func (s *OrderService) CreateOrder(userId uint, req model.OrderRequest) (*model.Order, error) {
	ctx := context.Background()

	// Create a lock key for this order creation
	// This prevents race conditions when multiple requests try to create an order
	// for the same user with the same products
	lockKey := fmt.Sprintf("order:create:user:%d", userId)
	if len(req.CartIDs) > 0 {
		for _, id := range req.CartIDs {
			lockKey += fmt.Sprintf(":cart:%d", id)
		}
	} else {
		lockKey += fmt.Sprintf(":product:%d:qty:%d", req.ProductId, req.Quantity)
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

		if address.UserId != userId {
			return ErrorAddressNotFound
		}

		// Get order details
		totalAmount, cartItems, err := s.GetOrderDetail(userId, req.CartIDs, &req.ProductId, &req.Quantity)
		if err != nil {
			return err
		}

		// Create order
		order = &model.Order{
			UserId:        userId,
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
			product, err := s.productRepo.GetProductByID(item.ProductId)
			if err != nil {
				return err
			}

			if product.StockCount < item.Quantity {
				return ErrorOutOfStock
			}

			orderItem := model.OrderItem{
				ProductId: item.ProductId,
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
			fmt.Printf("Would decrease stock for product %d by %d\n", item.ProductId, item.Quantity)
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
	cacheKey := fmt.Sprintf("order:%d", order.ID)
	_ = s.cacheService.Set(ctx, cacheKey, *order, 30*time.Minute)

	return order, orderErr
}

// GetOrderByID gets an order by ID
func (s *OrderService) GetOrderByID(id uint, userId uint) (*model.Order, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("order:%d", id)

	// Try to get from cache
	var order model.Order
	err := s.cacheService.GetObject(ctx, cacheKey, &order)
	if err == nil && order.UserId == userId {
		return &order, nil
	}

	// If not in cache or not owned by user, get from database
	orderPtr, err := s.orderRepo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the order belongs to the user
	if orderPtr.UserId != userId {
		return nil, ErrorOrderNotFound
	}

	// Cache for future requests
	_ = s.cacheService.Set(ctx, cacheKey, *orderPtr, 30*time.Minute)

	return orderPtr, nil
}

// GetOrderByOrderNo gets an order by order number
func (s *OrderService) GetOrderByOrderNo(orderNo string, userId uint) (*model.Order, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("order:no:%s", orderNo)

	// Try to get from cache
	var order model.Order
	err := s.cacheService.GetObject(ctx, cacheKey, &order)
	if err == nil && order.UserId == userId {
		return &order, nil
	}

	// If not in cache or not owned by user, get from database
	orderPtr, err := s.orderRepo.GetOrderByOrderNo(orderNo)
	if err != nil {
		return nil, err
	}

	// Ensure the order belongs to the user
	if orderPtr.UserId != userId {
		return nil, ErrorOrderNotFound
	}

	// Cache for future requests
	_ = s.cacheService.Set(ctx, cacheKey, *orderPtr, 30*time.Minute)

	return orderPtr, nil
}

// UpdateOrderStatus updates the status of an order
func (s *OrderService) UpdateOrderStatus(id uint, userId uint, status int) error {
	ctx := context.Background()

	order, err := s.GetOrderByID(id, userId)
	if err != nil {
		return err
	}

	// Use a lock to prevent race conditions when updating order status
	lockKey := fmt.Sprintf("order:status:%d", id)

	err = redis.WithLock(ctx, lockKey, 30*time.Second, func() error {
		// Update status in database
		err := s.orderRepo.UpdateOrderStatus(order.ID, status)
		if err != nil {
			return err
		}

		// Invalidate cache
		cacheKey := fmt.Sprintf("order:%d", id)
		s.cacheService.Delete(ctx, cacheKey)

		return nil
	})

	return err
}

// GetOrdersByUserId gets orders for a user with pagination
func (s *OrderService) GetOrdersByUserId(userId uint, page, pageSize int, status *int) (*model.Pagination, error) {
	ctx := context.Background()

	// Generate cache key
	cacheKey := fmt.Sprintf("orders:user:%d:page:%d:size:%d", userId, page, pageSize)
	if status != nil {
		cacheKey += fmt.Sprintf(":status:%d", *status)
	}

	// Try to get from cache
	var pagination model.Pagination
	err := s.cacheService.GetObject(ctx, cacheKey, &pagination)
	if err == nil {
		return &pagination, nil
	}

	// If not in cache, get from database
	orders, total, err := s.orderRepo.GetOrdersByUserId(userId, page, pageSize, status)
	if err != nil {
		return nil, err
	}

	pagination = model.NewPagination(total, page, pageSize, orders)

	// Cache for 5 minutes (shorter time since order list changes more frequently)
	_ = s.cacheService.Set(ctx, cacheKey, pagination, 5*time.Minute)

	return &pagination, nil
}
