package main

import (
	"boilerplate/internal/config"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.MustLoad()
	log := SetupLogger(cfg.Env)

	_ = log

	router := chi.NewRouter()

	log.Info("Start Server", slog.String("adress", cfg.Http.Adress))

	srv := &http.Server{
		Addr:         cfg.Http.Adress,
		Handler:      router,
		ReadTimeout:  cfg.Http.Timeout,
		WriteTimeout: cfg.Http.Timeout,
		IdleTimeout:  cfg.Http.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("Failed to start Server")
	}

}
