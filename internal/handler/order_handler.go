package handler

import (
	"net/http"
	"strconv"

	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/request"
	"github.com/colinjuang/shop-go/internal/response"
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
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	orderIDStr := c.Query("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Invalid order ID"))
		return
	}

	// Get order details
	orderDetail, err := h.orderService.GetOrderDetail(reqUser.UserID, orderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(orderDetail))
}

// GetOrderAddress gets the user's addresses for order
func (h *OrderHandler) GetOrderAddress(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	addresses, err := h.addressService.GetAddressesByUserID(reqUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// Try to get default address first
	defaultAddress, _ := h.addressService.GetDefaultAddressByUserID(c)

	resp := gin.H{
		"addresses":       addresses,
		"default_address": defaultAddress,
	}

	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err := h.orderService.CreateOrder(reqUser.UserID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(nil))
}

// GetWechatPayInfo gets WeChat payment information
func (h *OrderHandler) GetWechatPayInfo(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	orderNo := c.Query("order_no")
	if orderNo == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Missing order number"))
		return
	}

	// Get order by order number
	order, err := h.orderService.GetOrderByOrderNo(orderNo, reqUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// In a real implementation, we would call WeChat Payment API
	// For now, we'll return a mock response
	paymentResponse := response.PaymentResponse{
		PaymentID: "wx" + orderNo,
		AppID:     "your-app-id",
		TimeStamp: strconv.FormatInt(order.CreatedAt.Unix(), 10),
		NonceStr:  "random-string",
		Package:   "prepay_id=wx123456789",
		SignType:  "MD5",
		PaySign:   "signature",
	}

	c.JSON(http.StatusOK, response.SuccessResponse(paymentResponse))
}

// CheckWechatPayStatus checks WeChat payment status
func (h *OrderHandler) CheckWechatPayStatus(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
		return
	}

	orderNo := c.Query("order_no")
	if orderNo == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Missing order number"))
		return
	}

	// Get order by order number
	order, err := h.orderService.GetOrderByOrderNo(orderNo, reqUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// In a real implementation, we would check payment status with WeChat
	// For now, we'll simulate payment success
	if order.Status == model.OrderStatusPending {
		// Update order status to paid
		err = h.orderService.UpdateOrderStatus(order.ID, reqUser.UserID, model.OrderStatusPaid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		// Reload order
		order, err = h.orderService.GetOrderByID(order.ID, reqUser.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
			return
		}
	}

	resp := gin.H{
		"order_no": order.OrderNo,
		"status":   order.Status,
		"paid":     order.Status >= model.OrderStatusPaid,
	}

	c.JSON(http.StatusOK, response.SuccessResponse(resp))
}

// GetOrderList gets orders for a user with pagination
func (h *OrderHandler) GetOrderList(c *gin.Context) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(http.StatusUnauthorized, "Unauthorized"))
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
	pagination, err := h.orderService.GetOrdersByUserID(reqUser.UserID, page, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(pagination))
}
