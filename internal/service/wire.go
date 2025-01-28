package service

import (
	"pet/configs"
	"pet/internal/data"

	"github.com/google/wire"
)

type Option struct {
	Config   *configs.Config
	UserRepo *data.User
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewUser,
)
