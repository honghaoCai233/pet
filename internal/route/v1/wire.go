package v1

import (
	"pet/internal/service"

	"github.com/google/wire"
)

type Option struct {
	UserSrv *service.UserService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(Option), "*"),
	NewUser,
)
