package route

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
	"pet/configs"
	v1user "pet/internal/route/v1"
	"pet/internal/service"
)

type WireOption struct {
	Log     *zap.SugaredLogger
	Conf    *configs.Config
	Handler *gin.Engine

	UserService *service.UserService
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(WireOption), "*"),
	v1user.ProviderSet,
	NewGinEngine,
	NewHttpEngine,
)
