package service

import (
	"time"

	"github.com/colinjuang/shop-go/internal/config"
	"github.com/colinjuang/shop-go/internal/model"
	"github.com/colinjuang/shop-go/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"gorm.io/gorm"
)

type WechatLoginService struct {
	config   *config.Config
	userRepo *repository.UserRepository
}

func NewWechatLoginService(db *gorm.DB) *WechatLoginService {
	return &WechatLoginService{
		config:   config.GetConfig(),
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *WechatLoginService) WechatMiniLogin(code string) (string, error) {
	// 初始化小程序
	miniProgram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     s.config.Wechat.AppID,
		AppSecret: s.config.Wechat.AppSecret,
		Cache:     cache.NewMemory(),
	})

	// 获取微信用户信息
	sessionInfo, err := miniProgram.GetAuth().Code2Session(code)
	if err != nil {
		return "", err
	}

	// 查询用户表是否存在
	user, err := s.userRepo.GetUserByOpenID(sessionInfo.OpenID)
	if err != nil {
		return "", err
	}

	if user == nil {
		// 创建用户
		user = &model.User{
			OpenID: sessionInfo.OpenID,
		}
		err = s.userRepo.CreateUser(user)
		if err != nil {
			return "", err
		}
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"openid": user.OpenID,
		"exp":    time.Now().Add(time.Duration(s.config.JWT.ExpiresIn) * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(s.config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
