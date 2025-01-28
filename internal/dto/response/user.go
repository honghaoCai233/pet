package response

import (
	"pet/internal/data/ent"
	"time"
)

// UserResponse 用户响应
type UserResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	Phone          string    `json:"phone"`
	Age            int       `json:"age"`
	Role           string    `json:"role"`
	Description    string    `json:"description,omitempty"`
	Rating         float64   `json:"rating,omitempty"`
	CompletedTasks int       `json:"completed_tasks"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Total int64          `json:"total"`
	Items []UserResponse `json:"items"`
}

// NewUserResponse 将 ent.User 转换为 UserResponse
func NewUserResponse(user *ent.User) *UserResponse {
	return &UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Address:        user.Address,
		Phone:          user.Phone,
		Age:            user.Age,
		Role:           user.Role,
		Description:    user.Description,
		Rating:         user.Rating,
		CompletedTasks: user.CompletedTasks,
	}
}

// NewUserListResponse 将 []*ent.User 转换为 UserListResponse
func NewUserListResponse(users []*ent.User, total int64) *UserListResponse {
	items := make([]UserResponse, len(users))
	for i, user := range users {
		items[i] = *NewUserResponse(user)
	}
	return &UserListResponse{
		Total: total,
		Items: items,
	}
}
