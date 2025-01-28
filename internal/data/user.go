package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/data/ent/user"
)

type UserRepo struct {
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{
		data: data,
	}
}

func (r *UserRepo) Create(ctx context.Context, u *ent.User) (*ent.User, error) {
	builder := r.data.db.User.Create().
		SetName(u.Name).
		SetAddress(u.Address).
		SetPhone(u.Phone).
		SetAge(u.Age).
		SetRole(u.Role)

	if u.Description != "" {
		builder.SetDescription(u.Description)
	}
	if u.Rating != 0 {
		builder.SetRating(u.Rating)
	}
	if u.CompletedTasks != 0 {
		builder.SetCompletedTasks(u.CompletedTasks)
	}

	return builder.Save(ctx)
}

func (r *UserRepo) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	builder := r.data.db.User.UpdateOneID(u.ID).
		SetName(u.Name).
		SetAddress(u.Address).
		SetPhone(u.Phone).
		SetAge(u.Age).
		SetRole(u.Role)

	if u.Description != "" {
		builder.SetDescription(u.Description)
	}
	if u.Rating != 0 {
		builder.SetRating(u.Rating)
	}
	if u.CompletedTasks != 0 {
		builder.SetCompletedTasks(u.CompletedTasks)
	}

	return builder.Save(ctx)
}

func (r *UserRepo) Get(ctx context.Context, id int) (*ent.User, error) {
	return r.data.db.User.Get(ctx, id)
}

func (r *UserRepo) Delete(ctx context.Context, id int) error {
	return r.data.db.User.DeleteOneID(id).Exec(ctx)
}

func (r *UserRepo) List(ctx context.Context, page, pageSize int) ([]*ent.User, error) {
	offset := (page - 1) * pageSize
	return r.data.db.User.Query().
		Offset(offset).
		Limit(pageSize).
		All(ctx)
}

func (r *UserRepo) GetByPhone(ctx context.Context, phone string) (*ent.User, error) {
	return r.data.db.User.Query().
		Where(user.PhoneEQ(phone)).
		Only(ctx)
}
