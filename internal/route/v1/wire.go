package v1

import (
	"github.com/google/wire"
)

type WireOptions struct {
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(WireOptions), "*"),
	NewRouters,
)
