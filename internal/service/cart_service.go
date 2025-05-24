package service

import (
	"github.com/colinjuang/shop-go/internal/app/response"
	pkgerrors "github.com/colinjuang/shop-go/internal/pkg/errors"
	"github.com/colinjuang/shop-go/internal/pkg/minio"
	"github.com/colinjuang/shop-go/internal/repository"
	"gorm.io/gorm"
)

// CartService handles business logic for cart items
type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

// NewCartService creates a new cart service
func NewCartService(db *gorm.DB) *CartService {
	return &CartService{
		cartRepo:    repository.NewCartRepository(db),
		productRepo: repository.NewProductRepository(db),
	}
}

// AddToCart adds a product to the cart
func (s *CartService) AddToCart(userID uint64, productID uint64, quantity int) error {
	// 检查商品是否存在
	product, err := s.productRepo.GetProductByID(productID)
	if err != nil {
		return err
	}

	// 检查库存
	if product.StockCount < quantity {
		return pkgerrors.ErrOutOfStock
	}

	return s.cartRepo.AddToCart(userID, productID, quantity)
}

// GetCart gets all cart items for a user
func (s *CartService) GetCart(userID uint64) ([]response.CartResponse, error) {
	carts, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return nil, err
	}

	minioClient := minio.GetClient()
	var responses []response.CartResponse
	for _, item := range carts {
		response := response.CartResponse{
			ID:         item.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			Selected:   item.Selected,
			Name:       item.Product.Name,
			Price:      item.Product.Price,
			ImageUrl:   minioClient.GetFileURL(item.Product.ImageUrl),
			StockCount: item.Product.StockCount,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// UpdateCartStatus updates the status of a cart item
func (s *CartService) UpdateCartStatus(userID uint64, productID uint64, selected bool) error {
	carts, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return err
	}

	for _, item := range carts {
		if item.ProductID == productID {
			return s.cartRepo.UpdateCartStatus(item.ID, selected)
		}
	}

	return pkgerrors.ErrCartNotFound
}

// UpdateAllCartStatus updates the status of all cart items for a user
func (s *CartService) UpdateAllCartStatus(userID uint64, selected bool) error {
	_, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return err
	}

	return s.cartRepo.UpdateAllCartStatus(userID, selected)
}

// DeleteCart deletes a cart item
func (s *CartService) DeleteCart(userID uint64, id uint64) error {
	carts, err := s.cartRepo.GetCart(userID)
	if err != nil {
		return err
	}

	for _, item := range carts {
		if item.ID == id {
			return s.cartRepo.DeleteCart(id)
		}
	}

	return pkgerrors.ErrCartNotFound
}
