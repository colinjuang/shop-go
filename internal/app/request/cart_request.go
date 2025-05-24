package request

type AddToCartRequest struct {
	ProductID uint64 `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type UpdateCartStatusRequest struct {
	ID       uint64 `json:"id" binding:"required"`
	Selected bool   `json:"selected" binding:"required"`
}
