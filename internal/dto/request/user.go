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
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Phone       string  `json:"phone"`
	Age         int     `json:"age"`
	Role        string  `json:"role"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

// ListUsersRequest 用户列表请求
type ListUsersRequest struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=100"`
}
