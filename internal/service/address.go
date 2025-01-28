package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

// AddressService 地址服务
type AddressService struct {
	repo *data.AddressRepo
}

// NewAddressService 创建地址服务实例
func NewAddressService(opt *Option) *AddressService {
	return &AddressService{
		repo: opt.AddressRepo,
	}
}

// CreateAddress 创建地址
func (s *AddressService) CreateAddress(ctx context.Context, address *ent.Address) (*ent.Address, error) {
	return s.repo.Create(ctx, address)
}

// UpdateAddress 更新地址信息
func (s *AddressService) UpdateAddress(ctx context.Context, address *ent.Address) (*ent.Address, error) {
	return s.repo.Update(ctx, address)
}

// GetAddress 获取地址信息
func (s *AddressService) GetAddress(ctx context.Context, id int) (*ent.Address, error) {
	return s.repo.Get(ctx, id)
}

// DeleteAddress 删除地址
func (s *AddressService) DeleteAddress(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListAddresses 获取地址列表
func (s *AddressService) ListAddresses(ctx context.Context, page, pageSize int) ([]*ent.Address, error) {
	return s.repo.List(ctx, page, pageSize)
}

// ListAddressesByUser 获取用户的地址列表
func (s *AddressService) ListAddressesByUser(ctx context.Context, userID int) ([]*ent.Address, error) {
	return s.repo.ListByUser(ctx, userID)
}

// GetDefaultAddress 获取用户的默认地址
func (s *AddressService) GetDefaultAddress(ctx context.Context, userID int) (*ent.Address, error) {
	return s.repo.GetDefaultAddress(ctx, userID)
}

// SetDefaultAddress 设置默认地址
func (s *AddressService) SetDefaultAddress(ctx context.Context, id int, userID int) error {
	return s.repo.SetDefault(ctx, id, userID)
}
