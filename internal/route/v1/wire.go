package v1

import (
	"pet/internal/service"

	"github.com/google/wire"
)

type Option struct {
	UserSrv      *service.UserService
	PetSrv       *service.PetService
	TaskSrv      *service.TaskService
	CommunitySrv *service.CommunityService
}

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewUserHandler,
	NewPetHandler,
	NewTaskHandler,
	NewCommunityHandler,
)
