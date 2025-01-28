package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

type UserService struct {
	repo *data.UserRepo
}

func NewUserService(opt *Option) *UserService {
	return &UserService{
		repo: opt.UserRepo,
	}
}

// RegisterUser 用户注册
func (s *UserService) RegisterUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return s.repo.Register(ctx, user)
}

// LoginUser 用户登录
func (s *UserService) LoginUser(ctx context.Context, username, password string) (*ent.User, error) {
	return s.repo.Login(ctx, username, password)
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	return s.repo.ChangePassword(ctx, userID, oldPassword, newPassword)
}

// UpdateUserStatus 更新用户状态
func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, status string) error {
	return s.repo.UpdateStatus(ctx, userID, status)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return s.repo.Update(ctx, user)
}

// GetUser 获取用户信息
func (s *UserService) GetUser(ctx context.Context, id int) (*ent.User, error) {
	return s.repo.Get(ctx, id)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListUsers 获取用户列表
func (s *UserService) ListUsers(ctx context.Context, page, pageSize int) ([]*ent.User, error) {
	return s.repo.List(ctx, page, pageSize)
}

// GetUserByPhone 通过手机号获取用户
func (s *UserService) GetUserByPhone(ctx context.Context, phone string) (*ent.User, error) {
	return s.repo.GetByPhone(ctx, phone)
}

// GetUserByUsername 通过用户名获取用户
func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return s.repo.GetByUsername(ctx, username)
}

// GetUserByEmail 通过邮箱获取用户
func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

// ValidateUserCredentials 验证用户凭证（用于内部使用）
func (s *UserService) ValidateUserCredentials(ctx context.Context, identifier, password string) (*ent.User, error) {
	// identifier 可以是用户名或邮箱
	return s.repo.Login(ctx, identifier, password)
}
