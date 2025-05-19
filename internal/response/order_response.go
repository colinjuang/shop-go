package response

import "time"

type OrderDetailResponse struct {
	OrderID     uint64              `json:"orderID"`
	OrderNo     string              `json:"orderNo"`
	TotalAmount float64             `json:"totalAmount"`
	Items       []OrderItemResponse `json:"items"`
	Address     AddressResponse     `json:"address"`
}

type OrderItemResponse struct {
	ProductID uint64  `json:"productID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Name      string  `json:"name"`
	ImageUrl  string  `json:"imageUrl"`
}

type CreateOrderResponse struct {
	OrderID       uint64              `json:"orderID"`
	OrderNo       string              `json:"orderNo"`
	TotalAmount   float64             `json:"totalAmount"`
	PaymentAmount float64             `json:"paymentAmount"`
	Items         []OrderItemResponse `json:"items"`
	Address       AddressResponse     `json:"address"`
	PaymentType   int                 `json:"paymentType"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}
