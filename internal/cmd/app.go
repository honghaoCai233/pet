package cmd

import (
	"pet/internal/cron"
	"pet/internal/route"
)

type App struct {
	cron *cron.Engine
	http *route.HttpEngine
}

func NewApp(cron *cron.Engine, http *route.HttpEngine) *App {
	return &App{cron: cron, http: http}
}

func (a *App) Run() error {
	err := a.cron.Run()
	if err != nil {
		return err
	}
	return a.http.Run()
}
