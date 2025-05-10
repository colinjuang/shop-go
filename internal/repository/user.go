package repository

import (
	"shop-go/internal/model"
)

// UserRepository handles database operations for users
type UserRepository struct{}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetUserByOpenID gets a user by OpenID
func (r *UserRepository) GetUserByOpenID(openID string) (*model.User, error) {
	var user model.User
	result := DB.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *model.User) error {
	return DB.Create(user).Error
}

// UpdateUser updates a user
func (r *UserRepository) UpdateUser(user *model.User) error {
	return DB.Save(user).Error
}

// GetUserByID gets a user by ID
func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
