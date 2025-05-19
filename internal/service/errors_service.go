package service

import "errors"

// Common service errors
var (
	ErrorOutOfStock       = errors.New("product out of stock")
	ErrorCartItemNotFound = errors.New("cart item not found")
	ErrorOrderNotFound    = errors.New("order not found")
	ErrorUserNotFound     = errors.New("user not found")
	ErrorProductNotFound  = errors.New("product not found")
	ErrorAddressNotFound  = errors.New("address not found")
	ErrorUnauthorized     = errors.New("unauthorized operation")
	ErrorPaymentFailed    = errors.New("payment failed")
)
