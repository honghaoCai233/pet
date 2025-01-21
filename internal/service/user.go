package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
	"pet/internal/dto/request"
)

type UserService struct {
	opt      *Option
	userRepo *data.User
}

func NewUser(opt *Option) *UserService {
	return &UserService{opt: opt}
}

func (s *UserService) Create(ctx context.Context, req *request.CreateUserRequest) (*ent.User, error) {
	return s.userRepo.Create(ctx, req)
}
