package response

// type CartResponse struct {
// 	ID        uint   `json:"id"`
// 	ProductID uint   `json:"product_id"`
// 	Quantity  int    `json:"quantity"`
// 	CreatedAt string `json:"created_at"`
// 	UpdatedAt string `json:"updated_at"`
// }

type CartResponse struct {
	ID         uint64  `json:"id"`
	ProductID  uint64  `json:"productId"`
	Quantity   int     `json:"quantity"`
	Selected   bool    `json:"selected"`
	Blessing   string  `json:"blessing"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	ImageUrl   string  `json:"imageUrl"`
	StockCount int     `json:"stockCount"`
}
