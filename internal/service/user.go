package service

import (
	"errors"
	"time"

	"github.com/colinjuang/shop-go/internal/dto"
	"github.com/colinjuang/shop-go/internal/middleware"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/repository"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// LoginWithWechat handles login with WeChat
func (s *UserService) LoginWithWechat(openID, nickname, avatar string, gender int, city, province, country string) (string, error) {
	// Check if user exists
	user, err := s.userRepo.GetUserByOpenID(openID)
	if err != nil {
		// Create new user if not found
		newUser := &model.User{
			OpenID:   openID,
			Nickname: nickname,
			Avatar:   avatar,
			Gender:   gender,
			City:     city,
			Province: province,
			Country:  country,
		}

		if err := s.userRepo.CreateUser(newUser); err != nil {
			return "", err
		}

		user = newUser
	}

	// Generate JWT token
	token, err := middleware.GenerateToken(middleware.UserClaim{
		UserID: user.ID,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID gets a user by ID
func (s *UserService) GetUserByID(id uint64) (*dto.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		OpenID:    user.OpenID,
		Nickname:  user.Nickname,
		AvatarURL: user.Avatar,
		Gender:    user.Gender,
		City:      user.City,
		Province:  user.Province,
		Country:   user.Country,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(id uint64, updateInfo model.UserUpdateInfo) error {
	user, err := s.userRepo.GetUserByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	// Update user info
	user.Nickname = updateInfo.Nickname
	user.Avatar = updateInfo.Avatar
	user.Gender = updateInfo.Gender
	user.City = updateInfo.City
	user.Province = updateInfo.Province
	user.Country = updateInfo.Country

	return s.userRepo.UpdateUser(user)
}

// GetUserByUsername gets a user by username
func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}
