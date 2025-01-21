//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"pet/configs"
	"pet/internal/clients"
	"pet/internal/cmd"
	"pet/internal/cron"
	"pet/internal/data"
	"pet/internal/route"
	"pet/internal/service"
)

func build() (*cmd.App, func(), error) {
	panic(wire.Build(
		configs.InitConfig,
		clients.ProviderSet,
		cron.ProviderSet,
		service.ProviderSet,
		route.ProviderSet,
		data.ProviderSet,
		cmd.NewApp,
	))
}
