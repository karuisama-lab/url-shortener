package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/storage/sqlite"

	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

const (
	StoragePath = "./storage/storage.db"
)

func main() {
	cfg := config.MustLoad()

	log := SetupLogger(cfg.Env)

	storage, err := sqlite.New(StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)

	}

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
}

func SeyupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
