package constant

// 用户状态
const (
	// 正常
	UserStatusNormal = iota + 1
	// 禁用
	UserStatusDisabled
)

// 用户状态描述
var UserStatusDesc = map[int]string{
	UserStatusNormal:   "正常",
	UserStatusDisabled: "禁用",
}

// 用户角色
const (
	// 普通用户
	UserRoleNormal = iota + 1
	// 管理员
	UserRoleAdmin
)

// 用户角色描述
var UserRoleDesc = map[int]string{
	UserRoleNormal: "普通用户",
	UserRoleAdmin:  "管理员",
}
