package service

import (
	"pet/configs"
	"pet/internal/data"

	"github.com/google/wire"
)

type Option struct {
	Config                *configs.Config
	UserRepo              *data.UserRepo
	PetRepo               *data.PetRepo
	TaskRepo              *data.TaskRepo
	CommunityRepo         *data.CommunityRepo
	SitterApplicationRepo *data.SitterApplicationRepo
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewUserService,
	NewPetService,
	NewTaskService,
	NewCommunityService,
	NewSitterApplicationService,
)
