package request

// User DTO 包含用户相关的请求和响应结构体

// --- 请求结构体 ---

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender" binding:"oneof=0 1 2"` // 0: 未知, 1: 男, 2: 女
	City     string `json:"city"`
	Province string `json:"province"`
	District string `json:"district"`
}

// WechatLoginRequest 微信登录请求
type WechatLoginRequest struct {
	Code     string `json:"code" binding:"required"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	City     string `json:"city"`
	Province string `json:"province"`
	District string `json:"district"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	City     string `json:"city"`
	Province string `json:"province"`
	District string `json:"district"`
}
