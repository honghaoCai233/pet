package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

type CommunityService struct {
	repo *data.CommunityRepo
}

func NewCommunityService(opt *Option) *CommunityService {
	return &CommunityService{
		repo: opt.CommunityRepo,
	}
}

// CreateCommunity 创建社区帖子
func (s *CommunityService) CreateCommunity(ctx context.Context, community *ent.Community) (*ent.Community, error) {
	return s.repo.Create(ctx, community)
}

// UpdateCommunity 更新社区帖子
func (s *CommunityService) UpdateCommunity(ctx context.Context, community *ent.Community) (*ent.Community, error) {
	return s.repo.Update(ctx, community)
}

// GetCommunity 获取社区帖子
func (s *CommunityService) GetCommunity(ctx context.Context, id int) (*ent.Community, error) {
	return s.repo.Get(ctx, id)
}

// DeleteCommunity 删除社区帖子
func (s *CommunityService) DeleteCommunity(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListCommunity 获取社区帖子列表
func (s *CommunityService) ListCommunity(ctx context.Context, page, pageSize int) ([]*ent.Community, error) {
	return s.repo.List(ctx, page, pageSize)
}

// ListCommunityByAuthor 获取用户的社区帖子列表
func (s *CommunityService) ListCommunityByAuthor(ctx context.Context, authorID int) ([]*ent.Community, error) {
	return s.repo.ListByAuthor(ctx, authorID)
}

// ListCommunityByPet 获取宠物的社区帖子列表
func (s *CommunityService) ListCommunityByPet(ctx context.Context, petID int) ([]*ent.Community, error) {
	return s.repo.ListByPet(ctx, petID)
}

// ListCommunityByType 获取指定类型的社区帖子列表
func (s *CommunityService) ListCommunityByType(ctx context.Context, postType string, page, pageSize int) ([]*ent.Community, error) {
	return s.repo.ListByType(ctx, postType, page, pageSize)
}

// UpdateCommunityLikes 更新帖子点赞数
func (s *CommunityService) UpdateCommunityLikes(ctx context.Context, id int, increment bool) error {
	return s.repo.UpdateLikes(ctx, id, increment)
}

// UpdateCommunityViews 更新帖子浏览量
func (s *CommunityService) UpdateCommunityViews(ctx context.Context, id int) error {
	return s.repo.UpdateViews(ctx, id)
}
