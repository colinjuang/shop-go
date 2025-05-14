package service

import (
	"shop-go/internal/model"
	"shop-go/internal/repository"
)

// CartService handles business logic for cart items
type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

// NewCartService creates a new cart service
func NewCartService() *CartService {
	return &CartService{
		cartRepo:    repository.NewCartRepository(),
		productRepo: repository.NewProductRepository(),
	}
}

// AddToCart adds a product to the cart
func (s *CartService) AddToCart(userID, productID uint, quantity int) error {
	// Check if the product exists
	product, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		return err
	}

	// Check stock
	if product.StockCount < quantity {
		return ErrorOutOfStock
	}

	return s.cartRepo.AddToCart(userID, productID, quantity)
}

// GetCartItems gets all cart items for a user
func (s *CartService) GetCartItems(userID uint) ([]model.CartItemResponse, error) {
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return nil, err
	}

	var responses []model.CartItemResponse
	for _, item := range cartItems {
		response := model.CartItemResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			Selected:   item.Selected,
			Name:       item.Product.Name,
			Price:      item.Product.Price,
			ImageUrl:   item.Product.ImageUrl,
			StockCount: item.Product.StockCount,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateCartItemStatus updates the status of a cart item
func (s *CartService) UpdateCartItemStatus(id uint, userID uint, selected bool) error {
	// Check if the cart item belongs to the user
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return err
	}

	for _, item := range cartItems {
		if item.ID == id {
			return s.cartRepo.UpdateCartItemStatus(id, selected)
		}
	}

	return ErrorCartItemNotFound
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (s *CartService) UpdateAllCartItemStatus(userID uint, selected bool) error {
	return s.cartRepo.UpdateAllCartItemStatus(userID, selected)
}

// DeleteCartItem deletes a cart item
func (s *CartService) DeleteCartItem(id uint, userID uint) error {
	// Check if the cart item belongs to the user
	cartItems, err := s.cartRepo.GetCartItems(userID)
	if err != nil {
		return err
	}

	for _, item := range cartItems {
		if item.ID == id {
			return s.cartRepo.DeleteCartItem(id)
		}
	}

	return ErrorCartItemNotFound
}
