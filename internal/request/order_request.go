package request

// OrderRequest represents the order creation request
type CreateOrderRequest struct {
	CartIDs     []uint64 `json:"cartIDs"`
	AddressID   uint64   `json:"addressID" binding:"required"`
	PaymentType int      `json:"paymentType" binding:"required"`
}
