package app

import (
	"github.com/ednailson/api-base-project-go/handlers"
	"github.com/ednailson/httping-go"
)

type App struct {
	server httping.IServer
}

func LoadApp(cfg Config) *App {
	var application App
	application.server = httping.NewHttpServer(cfg.API.Host, cfg.API.Port)
	return &application
}

func (a *App) loadRoutes() {
	a.server.NewRoute(nil, "/example").POST(handlers.ExampleHandler)
}

func (a *App) Run() <-chan error {
	return a.Run()
}

func (a *App) Close() {}
