package data

import (
	"pet/configs"

	"github.com/google/wire"
)

type Option struct {
	Configs *configs.Config
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewData,
	NewUserRepo,
	NewPetRepo,
	NewTaskRepo,
	NewCommunityRepo,
	NewSitterApplicationRepo,
)
