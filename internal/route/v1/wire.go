package v1

import (
	"pet/internal/clients"
	"pet/internal/service"

	"github.com/google/wire"
)

type Option struct {
	UserSrv              *service.UserService
	PetSrv               *service.PetService
	TaskSrv              *service.TaskService
	CommunitySrv         *service.CommunityService
	SitterApplicationSrv *service.SitterApplicationService
	AddressSrv           *service.AddressService
	OSSClient            *clients.OSSClient
}

// ProviderSet is router providers.
var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewRouters,
)
