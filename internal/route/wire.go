package route

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"

	"pet/configs"
	v1 "pet/internal/route/v1"
)

type WireOption struct {
	Log     *zap.SugaredLogger
	Conf    *configs.Config
	Handler *gin.Engine

	Routers []v1.Router
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(WireOption), "*"),
	v1.ProviderSet,
	NewGinEngine,
	NewHttpEngine,
)
