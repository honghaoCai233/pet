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

// CreateUser 创建用户
func (s *UserService) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return s.repo.Create(ctx, user)
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
