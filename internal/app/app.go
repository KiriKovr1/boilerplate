package app

import (
	httpsrv "boilerplate/internal/app/http"
	"boilerplate/internal/config"
	"boilerplate/internal/lib/sl"
	"log/slog"
)

type App struct {
	srv Srv
	log *slog.Logger
}

type Srv interface {
	Listen()
	Stop()
}

func MustLoad(log *slog.Logger, cfg *config.HttpServer) *App {
	app, err := New(log, cfg)
	if err != nil {
		log.Error("Unable to create application", sl.Error(err))
		panic(err)
	}

	return app
}

func New(log *slog.Logger, cfg *config.HttpServer) (*App, error) {
	srv := httpsrv.New(log, cfg)

	return &App{
		srv: srv,
		log: log,
	}, nil
}

func (a *App) Start() {
	a.srv.Listen()
}

func (a *App) Stop() {
	a.srv.Stop()
}
