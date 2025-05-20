package request

// OrderRequest represents the order creation request
type CreateOrderRequest struct {
	CartIDs     []uint64 `json:"cartIDs"`
	AddressID   uint64   `json:"addressID" binding:"required"`
	PaymentType int      `json:"paymentType" binding:"required"`
}

type CreateOrderAndPayRequest struct {
	AddressID uint64 `json:"addressID" binding:"required"`
	ProductID uint64 `json:"productID" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
	Blessing  string `json:"blessing"`
	Remark    string `json:"remark"`
}
