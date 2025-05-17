package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// AddressRepository handles database operations for addresses
type AddressRepository struct {
	db *gorm.DB
}

// NewAddressRepository creates a new address repository
func NewAddressRepository() *AddressRepository {
	return &AddressRepository{
		db: database.GetDB(),
	}
}

// CreateAddress creates a new address
func (r *AddressRepository) CreateAddress(address *model.Address) error {
	// Set all other addresses as non-default if this one is default
	if address.IsDefault == 1 {
		r.db.Model(&model.Address{}).Where("user_id = ?", address.UserID).Update("is_default", 0)
	}

	return r.db.Create(address).Error
}

// GetAddressByID gets an address by ID
func (r *AddressRepository) GetAddressByID(id uint64) (*model.Address, error) {
	var address model.Address
	result := r.db.First(&address, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

// UpdateAddress updates an address
func (r *AddressRepository) UpdateAddress(address *model.Address) error {
	// Set all other addresses as non-default if this one is default
	if address.IsDefault == 1 {
		r.db.Model(&model.Address{}).Where("user_id = ? AND id != ?", address.UserID, address.ID).Update("is_default", 0)
	}

	return r.db.Save(address).Error
}

// DeleteAddress deletes an address
func (r *AddressRepository) DeleteAddress(id uint64) error {
	return r.db.Delete(&model.Address{}, "id = ?", id).Error
}

// GetAddressesByUserID gets all addresses for a user
func (r *AddressRepository) GetAddressesByUserID(userID uint64) ([]model.Address, error) {
	var addresses []model.Address
	result := r.db.Where("user_id = ?", userID).Find(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}

// GetDefaultAddressByUserID gets the default address for a user
func (r *AddressRepository) GetDefaultAddressByUserID(userID uint64) (*model.Address, error) {
	var address model.Address
	result := r.db.Where("user_id = ? AND is_default = ?", userID, true).First(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}
