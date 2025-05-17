package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

// GetUserByOpenID gets a user by OpenID
func (r *UserRepository) GetUserByOpenID(openID string) (*model.User, error) {
	var user model.User
	result := r.db.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

// UpdateUser updates a user
func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

// GetUserByID gets a user by ID
func (r *UserRepository) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByUsername gets a user by username
func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var count int64
	r.db.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if count == 0 {
		return nil, nil
	}
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
