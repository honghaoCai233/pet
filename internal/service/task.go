package service

import (
	"context"
	"pet/internal/data"
	"pet/internal/data/ent"
)

type TaskService struct {
	repo *data.TaskRepo
}

func NewTaskService(opt *Option) *TaskService {
	return &TaskService{
		repo: opt.TaskRepo,
	}
}

// CreateTask 创建任务
func (s *TaskService) CreateTask(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	return s.repo.Create(ctx, task)
}

// UpdateTask 更新任务信息
func (s *TaskService) UpdateTask(ctx context.Context, task *ent.Task) (*ent.Task, error) {
	return s.repo.Update(ctx, task)
}

// GetTask 获取任务信息
func (s *TaskService) GetTask(ctx context.Context, id int) (*ent.Task, error) {
	return s.repo.Get(ctx, id)
}

// DeleteTask 删除任务
func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// ListTasks 获取任务列表
func (s *TaskService) ListTasks(ctx context.Context, page, pageSize int) ([]*ent.Task, error) {
	return s.repo.List(ctx, page, pageSize)
}

// ListTasksByPublisher 获取用户发布的任务列表
func (s *TaskService) ListTasksByPublisher(ctx context.Context, publisherID int) ([]*ent.Task, error) {
	return s.repo.ListByPublisher(ctx, publisherID)
}

// ListTasksBySitter 获取照看者接受的任务列表
func (s *TaskService) ListTasksBySitter(ctx context.Context, sitterID int) ([]*ent.Task, error) {
	return s.repo.ListBySitter(ctx, sitterID)
}

// ListTasksByPet 获取宠物的任务列表
func (s *TaskService) ListTasksByPet(ctx context.Context, petID int) ([]*ent.Task, error) {
	return s.repo.ListByPet(ctx, petID)
}

// UpdateTaskStatus 更新任务状态
func (s *TaskService) UpdateTaskStatus(ctx context.Context, id int, status string) error {
	return s.repo.UpdateStatus(ctx, id, status)
}
