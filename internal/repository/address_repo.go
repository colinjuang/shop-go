package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"gorm.io/gorm"
)

// AddressRepository 地址仓库
type AddressRepository struct {
	db *gorm.DB
}

// NewAddressRepository 实例
func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		db: db,
	}
}

// CreateAddress 创建新地址
func (r *AddressRepository) CreateAddress(address *model.Address) error {
	// 如果这个地址是默认地址，则将其他地址设置为非默认
	if address.IsDefault == 1 {
		r.db.Model(&model.Address{}).Where("user_id = ?", address.UserID).Update("is_default", 0)
	}

	return r.db.Create(address).Error
}

// GetAddressByID 获取地址
func (r *AddressRepository) GetAddressByID(id uint64) (*model.Address, error) {
	var address model.Address
	result := r.db.First(&address, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}

// UpdateAddress 更新地址
func (r *AddressRepository) UpdateAddress(address *model.Address) error {
	// 如果这个地址是默认地址，则将其他地址设置为非默认
	if address.IsDefault == 1 {
		r.db.Model(&model.Address{}).Where("user_id = ? AND id != ?", address.UserID, address.ID).Update("is_default", 0)
	}

	return r.db.Save(address).Error
}

// DeleteAddress 删除地址
func (r *AddressRepository) DeleteAddress(id uint64) error {
	return r.db.Delete(&model.Address{}, "id = ?", id).Error
}

// GetAddressesByUserID 获取用户所有地址
func (r *AddressRepository) GetAddressesByUserID(userID uint64) ([]model.Address, error) {
	var addresses []model.Address
	result := r.db.Where("user_id = ?", userID).Find(&addresses)
	if result.Error != nil {
		return nil, result.Error
	}
	return addresses, nil
}

// GetDefaultAddressByUserID 获取用户默认地址
func (r *AddressRepository) GetDefaultAddressByUserID(userID uint64) (*model.Address, error) {
	var address model.Address
	result := r.db.Where("user_id = ? AND is_default = ?", userID, true).First(&address)
	if result.Error != nil {
		return nil, result.Error
	}
	return &address, nil
}
