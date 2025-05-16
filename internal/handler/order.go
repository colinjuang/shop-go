package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/service"

	"github.com/gin-gonic/gin"
)

// OrderHandler handles order-related API endpoints
type OrderHandler struct {
	orderService   *service.OrderService
	addressService *service.AddressService
}

// NewOrderHandler creates a new order handler
func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		orderService:   service.NewOrderService(),
		addressService: service.NewAddressService(),
	}
}

// GetOrderDetail gets order details for checkout
func (h *OrderHandler) GetOrderDetail(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	// Handle direct purchase
	var productID *uint64
	var quantity *int

	if pidStr := c.Query("product_id"); pidStr != "" {
		if pid, err := strconv.ParseUint(pidStr, 10, 64); err == nil {
			productID = &pid
		}
	}

	if qtyStr := c.Query("quantity"); qtyStr != "" {
		if qty, err := strconv.Atoi(qtyStr); err == nil && qty > 0 {
			quantity = &qty
		}
	}

	// Handle cart checkout
	var cartIDs []uint64
	if cidsStr := c.QueryArray("cart_ids[]"); len(cidsStr) > 0 {
		for _, cidStr := range cidsStr {
			if cid, err := strconv.ParseUint(cidStr, 10, 64); err == nil {
				cartIDs = append(cartIDs, cid)
			}
		}
	}

	// Get order details
	totalAmount, cartItems, err := h.orderService.GetOrderDetail(userID.(uint64), cartIDs, productID, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// Prepare response
	resp := gin.H{
		"total_amount": totalAmount,
		"items":        cartItems,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(resp))
}

// GetOrderAddress gets the user's addresses for order
func (h *OrderHandler) GetOrderAddress(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	addresses, err := h.addressService.GetAddressesByUserID(userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// Try to get default address first
	defaultAddress, _ := h.addressService.GetDefaultAddressByUserID(userID.(uint64))

	resp := gin.H{
		"addresses":       addresses,
		"default_address": defaultAddress,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(resp))
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var req model.OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	order, err := h.orderService.CreateOrder(userID.(uint64), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(order))
}

// GetWechatPayInfo gets WeChat payment information
func (h *OrderHandler) GetWechatPayInfo(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	orderNo := c.Query("order_no")
	if orderNo == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Missing order number"))
		return
	}

	// Get order by order number
	order, err := h.orderService.GetOrderByOrderNo(orderNo, userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// In a real implementation, we would call WeChat Payment API
	// For now, we'll return a mock response
	paymentResponse := model.PaymentResponse{
		PaymentID: "wx" + orderNo,
		AppID:     "your-app-id",
		TimeStamp: strconv.FormatInt(order.CreatedAt.Unix(), 10),
		NonceStr:  "random-string",
		Package:   "prepay_id=wx123456789",
		SignType:  "MD5",
		PaySign:   "signature",
	}

	c.JSON(http.StatusOK, model.SuccessResponse(paymentResponse))
}

// CheckWechatPayStatus checks WeChat payment status
func (h *OrderHandler) CheckWechatPayStatus(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	orderNo := c.Query("order_no")
	if orderNo == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResponse(http.StatusBadRequest, "Missing order number"))
		return
	}

	// Get order by order number
	order, err := h.orderService.GetOrderByOrderNo(orderNo, userID.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// In a real implementation, we would check payment status with WeChat
	// For now, we'll simulate payment success
	if order.Status == model.OrderStatusPending {
		// Update order status to paid
		err = h.orderService.UpdateOrderStatus(order.ID, userID.(uint64), model.OrderStatusPaid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		// Reload order
		order, err = h.orderService.GetOrderByID(order.ID, userID.(uint64))
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	resp := gin.H{
		"order_no": order.OrderNo,
		"status":   order.Status,
		"paid":     order.Status >= model.OrderStatusPaid,
	}

	c.JSON(http.StatusOK, model.SuccessResponse(resp))
}

// GetOrderList gets orders for a user with pagination
func (h *OrderHandler) GetOrderList(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	// Get query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// Apply minimum values
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 10
	}

	// Get status filter
	var status *int
	if statusStr := c.Query("status"); statusStr != "" {
		if s, err := strconv.Atoi(statusStr); err == nil {
			status = &s
		}
	}

	// Get orders
	pagination, err := h.orderService.GetOrdersByUserID(userID.(uint64), page, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse(pagination))
}
