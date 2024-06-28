package app

import (
	httpsrv "boilerplate/internal/app/http"
	"boilerplate/internal/config"
	"boilerplate/internal/lib/sl"
	"log/slog"
	"net/http"
)

type App struct {
	HttpServer *http.Server
	log        *slog.Logger
}

func MustLoad(log *slog.Logger, cfg *config.HttpServer) *App {
	app, err := New(log, cfg)
	if err != nil {
		log.Error("Unable to create application", sl.Error(err))
		panic("Unable to create application")
	}

	return app
}

func New(log *slog.Logger, cfg *config.HttpServer) (*App, error) {
	srv := httpsrv.New(log, cfg)

	return &App{
		HttpServer: srv,
		log:        log,
	}, nil
}

func (a *App) Listen() {
	log := a.log.With(slog.String("op", "App.Listen"))

	if err := a.HttpServer.ListenAndServe(); err != nil {
		log.Error("Unable to listen and serve", sl.Error(err))
	}

}

func (a *App) Stop() {

}
