package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/task"
)

type TaskRepo struct {
	data *Data
}

// NewTaskRepo .
func NewTaskRepo(data *Data) *TaskRepo {
	return &TaskRepo{
		data: data,
	}
}

func (r *TaskRepo) Create(ctx context.Context, t *ent.Task) (*ent.Task, error) {
	builder := r.data.db.Task.Create().
		SetPublisherID(t.PublisherID).
		SetPetID(t.PetID).
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetReward(t.Reward).
		SetStartTime(t.StartTime).
		SetEndTime(t.EndTime).
		SetStatus(t.Status).
		SetLocation(t.Location).
		SetVisitsCount(t.VisitsCount)

	if t.Requirements != "" {
		builder.SetRequirements(t.Requirements)
	}
	if t.CareInstructions != "" {
		builder.SetCareInstructions(t.CareInstructions)
	}
	if t.SitterID != 0 {
		builder.SetSitterID(t.SitterID)
	}

	return builder.Save(ctx)
}

func (r *TaskRepo) Update(ctx context.Context, t *ent.Task) (*ent.Task, error) {
	builder := r.data.db.Task.UpdateOneID(t.ID).
		SetPublisherID(t.PublisherID).
		SetPetID(t.PetID).
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetReward(t.Reward).
		SetStartTime(t.StartTime).
		SetEndTime(t.EndTime).
		SetStatus(t.Status).
		SetLocation(t.Location).
		SetVisitsCount(t.VisitsCount)

	if t.Requirements != "" {
		builder.SetRequirements(t.Requirements)
	}
	if t.CareInstructions != "" {
		builder.SetCareInstructions(t.CareInstructions)
	}
	if t.SitterID != 0 {
		builder.SetSitterID(t.SitterID)
	}

	return builder.Save(ctx)
}

func (r *TaskRepo) Get(ctx context.Context, id int) (*ent.Task, error) {
	return r.data.db.Task.Get(ctx, id)
}

func (r *TaskRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.Task.DeleteOneID(id).Exec(ctx)
}

func (r *TaskRepo) List(ctx context.Context, page, pageSize int) ([]*ent.Task, error) {
	offset := (page - 1) * pageSize
	return r.data.db.Task.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *TaskRepo) ListByPublisher(ctx context.Context, publisherID int) ([]*ent.Task, error) {
	return r.data.db.Task.Query().
		Where(task.PublisherIDEQ(publisherID)).
		All(ctx)
}

func (r *TaskRepo) ListBySitter(ctx context.Context, sitterID int) ([]*ent.Task, error) {
	return r.data.db.Task.Query().
		Where(task.SitterIDEQ(sitterID)).
		All(ctx)
}

func (r *TaskRepo) ListByPet(ctx context.Context, petID int) ([]*ent.Task, error) {
	return r.data.db.Task.Query().
		Where(task.PetIDEQ(petID)).
		All(ctx)
}

func (r *TaskRepo) UpdateStatus(ctx context.Context, id int, status string) error {
	return r.data.db.Task.UpdateOneID(id).
		SetStatus(status).
		Exec(ctx)
}
