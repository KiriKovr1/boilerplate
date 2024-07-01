package httpsrv

import (
	"boilerplate/internal/config"
	"boilerplate/internal/lib/sl"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	log        *slog.Logger
	HttpServer *http.Server
}

func New(log *slog.Logger, cfg *config.HttpServer) *App {
	router := chi.NewRouter()

	srv := &http.Server{
		Addr:         cfg.Adress,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &App{
		log,
		srv,
	}
}

func (a *App) Listen() {
	log := a.log.With(slog.String("op", "App.Listen"))

	log.Info(fmt.Sprintf("Starting server on %s", a.HttpServer.Addr))

	if err := a.HttpServer.ListenAndServe(); err != nil {
		if !errors.Is(http.ErrServerClosed, err) {
			log.Error("Unable to listen and serve", sl.Error(err))
			panic(err)
		}
	}
}

func (a *App) Stop() {
	log := a.log.With("op", "App.Stop")

	if err := a.HttpServer.Shutdown(context.Background()); err != nil {
		log.Error("Unable to Shutdown application", sl.Error(err))
		a.HttpServer.Close()
	}

	log.Info("Server closed successfull")
}
