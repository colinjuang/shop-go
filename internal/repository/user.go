package repository

import (
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/pkg/database"
	"gorm.io/gorm"
)

// UserRepository 用户仓库
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建并返回一个新的用户数据仓库实例
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

// GetUserByOpenID 获取用户
func (r *UserRepository) GetUserByOpenID(openID string) (*model.User, error) {
	var user model.User
	result := r.db.Where("open_id = ?", openID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser 创建用户
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

// UpdateUser 更新用户
func (r *UserRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

// GetUserByID 获取用户
func (r *UserRepository) GetUserByID(id uint64) (*model.User, error) {
	var user model.User
	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetUserByUsername 获取用户
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
