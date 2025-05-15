package service

import (
	"errors"

	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/repository"
)

// AddressService handles business logic for addresses
type AddressService struct {
	addressRepo *repository.AddressRepository
}

// NewAddressService creates a new address service
func NewAddressService() *AddressService {
	return &AddressService{
		addressRepo: repository.NewAddressRepository(),
	}
}

// CreateAddress creates a new address
func (s *AddressService) CreateAddress(userId uint, req model.AddressRequest) (*model.Address, error) {
	address := &model.Address{
		UserId:     userId,
		Name:       req.Name,
		Phone:      req.Phone,
		Province:   req.Province,
		City:       req.City,
		District:   req.District,
		DetailAddr: req.DetailAddr,
		IsDefault:  req.IsDefault,
	}

	err := s.addressRepo.CreateAddress(address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

// GetAddressByID gets an address by ID
func (s *AddressService) GetAddressByID(id uint, userId uint) (*model.Address, error) {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the address belongs to the user
	if address.UserId != userId {
		return nil, errors.New("address not found")
	}

	return address, nil
}

// UpdateAddress updates an address
func (s *AddressService) UpdateAddress(id uint, userId uint, req model.AddressRequest) error {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserId != userId {
		return errors.New("address not found")
	}

	// Update address info
	address.Name = req.Name
	address.Phone = req.Phone
	address.Province = req.Province
	address.City = req.City
	address.District = req.District
	address.DetailAddr = req.DetailAddr
	address.IsDefault = req.IsDefault

	return s.addressRepo.UpdateAddress(address)
}

// DeleteAddress deletes an address
func (s *AddressService) DeleteAddress(id uint, userId uint) error {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserId != userId {
		return errors.New("address not found")
	}

	return s.addressRepo.DeleteAddress(id)
}

// GetAddressesByUserId gets all addresses for a user
func (s *AddressService) GetAddressesByUserId(userId uint) ([]model.Address, error) {
	return s.addressRepo.GetAddressesByUserId(userId)
}

// GetDefaultAddressByUserId gets the default address for a user
func (s *AddressService) GetDefaultAddressByUserId(userId uint) (*model.Address, error) {
	return s.addressRepo.GetDefaultAddressByUserId(userId)
}
