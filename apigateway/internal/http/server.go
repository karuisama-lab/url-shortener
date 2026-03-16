package http

import (
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"url-shortener/internal/config"
)

func RunServer(cfg *config.Config, lg *slog.Logger) error {
	app := fiber.New()
	lg.Info("starting HTTP server", cfg.HTTPServer.Address)
	if err := app.Listen(cfg.HTTPServer.Address); err != nil {
		return err
	}

	return nil
}
