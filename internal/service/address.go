package service

import (
	"errors"

	"github.com/colinjuang/shop-go/internal/dto"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/gin-gonic/gin"
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
func (s *AddressService) CreateAddress(c *gin.Context, req dto.AddressRequest) (*dto.AddressResponse, error) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return nil, errors.New("unauthorized")
	}

	address := &model.Address{
		UserID:       reqUser.UserID,
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
func (s *AddressService) GetAddressByID(c *gin.Context, id uint64) (*dto.AddressResponse, error) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return nil, errors.New("unauthorized")
	}

	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the address belongs to the user
	if address.UserID != reqUser.UserID {
		return nil, errors.New("address not found")
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

// UpdateAddress updates an address
func (s *AddressService) UpdateAddress(c *gin.Context, req dto.AddressRequest) error {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return errors.New("unauthorized")
	}

	address, err := s.addressRepo.GetAddressByID(req.ID)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserID != reqUser.UserID {
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
func (s *AddressService) DeleteAddress(c *gin.Context, id uint64) error {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return errors.New("unauthorized")
	}

	address, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return err
	}

	// Ensure the address belongs to the user
	if address.UserID != reqUser.UserID {
		return errors.New("address not found")
	}

	return s.addressRepo.DeleteAddress(id)
}

// GetAddressesByUserID gets all addresses for a user
func (s *AddressService) GetAddressesByUserID(c *gin.Context) ([]*dto.AddressResponse, error) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return nil, errors.New("unauthorized")
	}

	addresses, err := s.addressRepo.GetAddressesByUserID(reqUser.UserID)
	if err != nil {
		return nil, err
	}

	addressResponses := make([]*dto.AddressResponse, len(addresses))
	for i, address := range addresses {
		addressResponses[i] = &dto.AddressResponse{
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
		}
	}

	return addressResponses, nil
}

// GetDefaultAddressByUserID gets the default address for a user
func (s *AddressService) GetDefaultAddressByUserID(c *gin.Context) (*dto.AddressResponse, error) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return nil, errors.New("unauthorized")
	}

	address, err := s.addressRepo.GetDefaultAddressByUserID(reqUser.UserID)
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
