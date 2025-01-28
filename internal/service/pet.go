package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

type PetService struct {
	repo *data.PetRepo
}

func NewPetService(opt *Option) *PetService {
	return &PetService{
		repo: opt.PetRepo,
	}
}

// CreatePet 创建宠物
func (s *PetService) CreatePet(ctx context.Context, pet *ent.Pet) (*ent.Pet, error) {
	return s.repo.Create(ctx, pet)
}

// UpdatePet 更新宠物信息
func (s *PetService) UpdatePet(ctx context.Context, pet *ent.Pet) (*ent.Pet, error) {
	return s.repo.Update(ctx, pet)
}

// GetPet 获取宠物信息
func (s *PetService) GetPet(ctx context.Context, id int) (*ent.Pet, error) {
	return s.repo.Get(ctx, id)
}

// DeletePet 删除宠物
func (s *PetService) DeletePet(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListPets 获取宠物列表
func (s *PetService) ListPets(ctx context.Context, page, pageSize int) ([]*ent.Pet, error) {
	return s.repo.List(ctx, page, pageSize)
}

// ListPetsByOwner 获取用户的宠物列表
func (s *PetService) ListPetsByOwner(ctx context.Context, ownerID int) ([]*ent.Pet, error) {
	return s.repo.ListByOwner(ctx, ownerID)
}
