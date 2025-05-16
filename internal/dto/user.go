package dto

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
}

// WechatLoginRequest 微信登录请求
type WechatLoginRequest struct {
	Code      string `json:"code" binding:"required"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Gender    int    `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

// --- 响应结构体 ---

// UserResponse 用户信息响应（不含敏感信息）
type UserResponse struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	OpenID    string `json:"open_id"`
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatar_url"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
