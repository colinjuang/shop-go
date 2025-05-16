package service

import (
	"errors"

	"github.com/colinjuang/shop-go/internal/dto"
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
func (s *AddressService) CreateAddress(userID uint64, req dto.AddressRequest) (*dto.AddressResponse, error) {
	address := &model.Address{
		UserID:       userID,
		Name:         req.Name,
		Phone:        req.Phone,
		Province:     req.Province,
		ProvinceCode: req.ProvinceCode,
		City:         req.City,
		CityCode:     req.CityCode,
		District:     req.District,
		DistrictCode: req.DistrictCode,
		DetailAddr:   req.DetailAddr,
		IsDefault:    req.IsDefault,
	}

	err := s.addressRepo.CreateAddress(address)
	if err != nil {
		return nil, err
	}

	return &dto.AddressResponse{
		ID:           address.ID,
		Phone:        address.Phone,
		Name:         address.Name,
		City:         address.City,
		CityCode:     address.CityCode,
		Province:     address.Province,
		ProvinceCode: address.ProvinceCode,
		District:     address.District,
		DistrictCode: address.DistrictCode,
		DetailAddr:   address.DetailAddr,
		FullAddr:     address.Province + address.City + address.District + address.DetailAddr,
		IsDefault:    address.IsDefault,
	}, nil
}

// GetAddressByID gets an address by ID
func (s *AddressService) GetAddressByID(id uint64, userID uint64) (*model.Address, error) {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the address belongs to the user
	if address.UserID != userID {
		return nil, errors.New("address not found")
	}

	return address, nil
}

// UpdateAddress updates an address
func (s *AddressService) UpdateAddress(id uint64, userID uint64, req dto.AddressRequest) error {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserID != userID {
		return errors.New("address not found")
	}

	// Update address info
	address.Name = req.Name
	address.Phone = req.Phone
	address.Province = req.Province
	address.ProvinceCode = req.ProvinceCode
	address.City = req.City
	address.CityCode = req.CityCode
	address.District = req.District
	address.DistrictCode = req.DistrictCode
	address.DetailAddr = req.DetailAddr
	address.IsDefault = req.IsDefault

	return s.addressRepo.UpdateAddress(address)
}

// DeleteAddress deletes an address
func (s *AddressService) DeleteAddress(id uint64, userID uint64) error {
	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserID != userID {
		return errors.New("address not found")
	}

	return s.addressRepo.DeleteAddress(id)
}

// GetAddressesByUserID gets all addresses for a user
func (s *AddressService) GetAddressesByUserID(userID uint64) ([]model.Address, error) {
	return s.addressRepo.GetAddressesByUserID(userID)
}

// GetDefaultAddressByUserID gets the default address for a user
func (s *AddressService) GetDefaultAddressByUserID(userID uint64) (*model.Address, error) {
	return s.addressRepo.GetDefaultAddressByUserID(userID)
}
