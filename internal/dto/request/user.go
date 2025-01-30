package request

type CreateUserRequest struct {
	Name        string  `json:"name" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	Phone       string  `json:"phone" binding:"required"`
	Age         int     `json:"age" binding:"required"`
	Role        string  `json:"role"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name        string  `form:"name"`
	Address     string  `form:"address"`
	Phone       string  `form:"phone"`
	Age         int     `form:"age"`
	Role        string  `form:"role"`
	Description string  `form:"description"`
	Rating      float64 `form:"rating"`
	Gender      string  `form:"gender" binding:"omitempty,oneof=male female"`
	Birthday    string  `form:"birthday"`
}

// UpdateUserPasswordRequest 修改密码请求
type UpdateUserPasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// UpdateUserStatusRequest 更新用户状态请求
type UpdateUserStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active disabled locked"`
}

// RegisterUserRequest 用户注册请求
type RegisterUserRequest struct {
	Phone       string  `json:"phone" binding:"required"`
	Password    string  `json:"password" binding:"required,min=6"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Age         int     `json:"age"`
	Role        string  `json:"role" default:"pet_sitter"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
	Avatar      string  `json:"avatar"`
	Gender      string  `json:"gender" binding:"omitempty,oneof=male female"`
	Birthday    string  `json:"birthday"`
}

// LoginUserRequest 用户登录请求
type LoginUserRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ListUsersRequest 用户列表请求
type ListUsersRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}
