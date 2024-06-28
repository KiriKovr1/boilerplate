package httpsrv

import (
	"boilerplate/internal/config"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(log *slog.Logger, cfg *config.HttpServer) *http.Server {
	router := chi.NewRouter()

	srv := &http.Server{
		Addr:         cfg.Adress,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return srv
}
