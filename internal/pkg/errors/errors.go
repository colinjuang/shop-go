package errors

import (
	"errors"
	"fmt"
)

// 定义常见错误基础类型
var (
	// 资源未找到类错误
	ErrNotFound = errors.New("resource not found")

	// 权限类错误
	ErrUnauthorized = errors.New("unauthorized operation")
	ErrForbidden    = errors.New("forbidden operation")

	// 业务逻辑错误
	ErrOutOfStock    = errors.New("product out of stock")
	ErrPaymentFailed = errors.New("payment failed")
	ErrInvalidInput  = errors.New("invalid input")
)

// 特定资源错误
var (
	ErrUserNotFound     = fmt.Errorf("user not found: %w", ErrNotFound)
	ErrProductNotFound  = fmt.Errorf("product not found: %w", ErrNotFound)
	ErrOrderNotFound    = fmt.Errorf("order not found: %w", ErrNotFound)
	ErrCartItemNotFound = fmt.Errorf("cart item not found: %w", ErrNotFound)
	ErrAddressNotFound  = fmt.Errorf("address not found: %w", ErrNotFound)
)

// 错误检查辅助函数
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

func IsForbidden(err error) bool {
	return errors.Is(err, ErrForbidden)
}

// 创建特定实例错误
func NewNotFoundError(resource string, id interface{}) error {
	return fmt.Errorf("%s with ID %v not found: %w", resource, id, ErrNotFound)
}

func NewUnauthorizedError(reason string) error {
	return fmt.Errorf("%s: %w", reason, ErrUnauthorized)
}

func NewOutOfStockError(productID interface{}) error {
	return fmt.Errorf("product %v is out of stock: %w", productID, ErrOutOfStock)
}
