package cron

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCrons,
	NewEngine,
)
