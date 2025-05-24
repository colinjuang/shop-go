package service

import (
	"errors"
	"time"

	"github.com/colinjuang/shop-go/internal/app/middleware"
	"github.com/colinjuang/shop-go/internal/app/request"
	"github.com/colinjuang/shop-go/internal/app/response"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/colinjuang/shop-go/internal/server"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService creates a new user service
func NewUserService() *UserService {
	server := server.GetServer()
	return &UserService{
		userRepo: repository.NewUserRepository(server.DB),
	}
}

// LoginWithWechat handles login with WeChat
func (s *UserService) LoginWithWechat(openID, nickname, avatar string, gender int, city, province, district string) (string, error) {
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
			District: district,
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
func (s *UserService) GetUserByID(c *gin.Context) (*response.UserResponse, error) {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return nil, errors.New("unauthorized")
	}
	user, err := s.userRepo.GetUserByID(reqUser.UserID)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		OpenID:    user.OpenID,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		City:      user.City,
		Province:  user.Province,
		District:  user.District,
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}, nil
}

// UpdateUser updates a user
func (s *UserService) UpdateUser(c *gin.Context, updateInfo request.UserUpdateRequest) error {
	reqUser := middleware.GetRequestUser(c)
	if reqUser == nil {
		return errors.New("unauthorized")
	}

	user, err := s.userRepo.GetUserByID(reqUser.UserID)
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
	user.District = updateInfo.District

	return s.userRepo.UpdateUser(user)
}

// Login gets a user by username and password
func (s *UserService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user.ID == 0 {
		return "", errors.New("用户不存在")
	}

	// 检查密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误或用户不存在")
	}

	// 生成token
	token, err := middleware.GenerateToken(middleware.UserClaim{
		UserID: user.ID,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

// Register creates a new user
func (s *UserService) Register(userRegisterRequest request.UserRegisterRequest) error {
	user, err := s.userRepo.GetUserByUsername(userRegisterRequest.Username)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("用户已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &model.User{
		Username: userRegisterRequest.Username,
		Password: string(hashedPassword),
		Nickname: userRegisterRequest.Nickname,
		Avatar:   userRegisterRequest.Avatar,
		Gender:   userRegisterRequest.Gender,
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *model.User) error {
	return s.userRepo.CreateUser(user)
}
