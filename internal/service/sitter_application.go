package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

type SitterApplicationService struct {
	repo *data.SitterApplicationRepo
}

func NewSitterApplicationService(opt *Option) *SitterApplicationService {
	return &SitterApplicationService{
		repo: opt.SitterApplicationRepo,
	}
}

// CreateApplication 创建申请
func (s *SitterApplicationService) CreateApplication(ctx context.Context, application *ent.SitterApplication) (*ent.SitterApplication, error) {
	return s.repo.Create(ctx, application)
}

// UpdateApplication 更新申请
func (s *SitterApplicationService) UpdateApplication(ctx context.Context, application *ent.SitterApplication) (*ent.SitterApplication, error) {
	return s.repo.Update(ctx, application)
}

// GetApplication 获取申请
func (s *SitterApplicationService) GetApplication(ctx context.Context, id int) (*ent.SitterApplication, error) {
	return s.repo.Get(ctx, id)
}

// DeleteApplication 删除申请
func (s *SitterApplicationService) DeleteApplication(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListApplications 获取申请列表
func (s *SitterApplicationService) ListApplications(ctx context.Context, page, pageSize int) ([]*ent.SitterApplication, error) {
	return s.repo.List(ctx, page, pageSize)
}

// ListApplicationsBySitter 获取照护者的申请列表
func (s *SitterApplicationService) ListApplicationsBySitter(ctx context.Context, sitterID int) ([]*ent.SitterApplication, error) {
	return s.repo.ListBySitter(ctx, sitterID)
}

// ListApplicationsByOwner 获取宠物主人收到的申请列表
func (s *SitterApplicationService) ListApplicationsByOwner(ctx context.Context, ownerID int) ([]*ent.SitterApplication, error) {
	return s.repo.ListByOwner(ctx, ownerID)
}

// ListApplicationsByPet 获取宠物的申请列表
func (s *SitterApplicationService) ListApplicationsByPet(ctx context.Context, petID int) ([]*ent.SitterApplication, error) {
	return s.repo.ListByPet(ctx, petID)
}

// ListApplicationsByStatus 获取指定状态的申请列表
func (s *SitterApplicationService) ListApplicationsByStatus(ctx context.Context, status string, page, pageSize int) ([]*ent.SitterApplication, error) {
	return s.repo.ListByStatus(ctx, status, page, pageSize)
}

// UpdateApplicationStatus 更新申请状态
func (s *SitterApplicationService) UpdateApplicationStatus(ctx context.Context, id int, status string, rejectReason *string) error {
	return s.repo.UpdateStatus(ctx, id, status, rejectReason)
}
