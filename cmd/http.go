package main

import (
	"boilerplate/internal/app"
	"boilerplate/internal/config"
	"boilerplate/internal/lib/sl"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := sl.SetupLogger(cfg.Env)

	log.Info("Start Server", slog.String("adress", cfg.Http.Adress))

	app := app.MustLoad(log, &cfg.Http)

	go app.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	sig := <-stop

	log.Info("Try to stop server", slog.String("signal", sig.String()))
	app.Stop()

	log.Info("Server is stopped")
}
