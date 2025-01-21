package data

import (
	"context"
	"pet/internal/data/ent"
	"pet/internal/dto/request"
)

type User struct {
	*Data
}

func NewUser(data *Data) *User {
	return &User{Data: data}
}

func (u *User) Create(ctx context.Context, req *request.CreateUserRequest) (*ent.User, error) {
	user, err := u.db.User.Create().
		SetAddress(req.Address).
		SetName(req.Name).
		Save(ctx)
	return user, err
}
