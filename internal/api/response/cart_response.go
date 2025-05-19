package response

type CartResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}




