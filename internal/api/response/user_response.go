package response

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

// --- 响应结构体 ---

// UserResponse 用户信息响应（不含敏感信息）
type UserResponse struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	OpenID    string `json:"open_id"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	District  string `json:"district"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}
