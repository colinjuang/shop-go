package service

import (
	"github.com/colinjuang/shop-go/internal/api/middleware"
	"github.com/colinjuang/shop-go/internal/model"
	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"
	"github.com/colinjuang/shop-go/internal/repository"
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
func (s *CartService) AddToCart(reqUser *middleware.UserClaim, productID uint64, quantity int) error {

	// Check if the product exists
	product, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		return err
	}

	// Check stock
	if product.StockCount < quantity {
		return pkgerrors.ErrOutOfStock
	}

	return s.cartRepo.AddToCart(reqUser.UserID, productID, quantity)
}

// GetCartItems gets all cart items for a user
func (s *CartService) GetCartItems(userID uint64) ([]model.CartItemResponse, error) {
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
func (s *CartService) UpdateCartItemStatus(id uint64, userID uint64, selected bool) error {
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

	return pkgerrors.ErrCartItemNotFound
}

// UpdateAllCartItemStatus updates the status of all cart items for a user
func (s *CartService) UpdateAllCartItemStatus(userID uint64, selected bool) error {
	return s.cartRepo.UpdateAllCartItemStatus(userID, selected)
}

// DeleteCartItem deletes a cart item
func (s *CartService) DeleteCartItem(id uint64, userID uint64) error {
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

	return pkgerrors.ErrCartItemNotFound
}
