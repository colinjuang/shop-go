package repository

import (
	"shop-go/internal/model"
)

// AddressRepository handles database operations for addresses
type AddressRepository struct{}

// NewAddressRepository creates a new address repository
func NewAddressRepository() *AddressRepository {
	return &AddressRepository{}
}

// CreateAddress creates a new address
func (r *AddressRepository) CreateAddress(address *model.Address) error {
	// Set all other addresses as non-default if this one is default
	if address.IsDefault {
		DB.Model(&model.Address{}).Where("user_id = ?", address.UserID).Update("is_default", false)
	}

	return DB.Create(address).Error
}

// GetAddressByID gets an address by ID
func (r *AddressRepository) GetAddressByID(id uint) (*model.Address, error) {
	var address model.Address
	result := DB.First(&address, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

// UpdateAddress updates an address
func (r *AddressRepository) UpdateAddress(address *model.Address) error {
	// Set all other addresses as non-default if this one is default
	if address.IsDefault {
		DB.Model(&model.Address{}).Where("user_id = ? AND id != ?", address.UserID, address.ID).Update("is_default", false)
	}

	return DB.Save(address).Error
}

// DeleteAddress deletes an address
func (r *AddressRepository) DeleteAddress(id uint) error {
	return DB.Delete(&model.Address{}, id).Error
}

// GetAddressesByUserID gets all addresses for a user
func (r *AddressRepository) GetAddressesByUserID(userID uint) ([]model.Address, error) {
	var addresses []model.Address
	result := DB.Where("user_id = ?", userID).Find(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}

// GetDefaultAddressByUserID gets the default address for a user
func (r *AddressRepository) GetDefaultAddressByUserID(userID uint) (*model.Address, error) {
	var address model.Address
	result := DB.Where("user_id = ? AND is_default = ?", userID, true).First(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}
